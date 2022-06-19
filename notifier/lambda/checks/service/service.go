package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/aws/aws-lambda-go/events"

	"github.com/google/go-github/v45/github"

	"github.com/yunomu/checks-codepipeline/lambda/checks/checkrun"
	"github.com/yunomu/checks-codepipeline/lambda/checks/db"
	"github.com/yunomu/checks-codepipeline/lambda/checks/pipelinestat"
)

type Service struct {
	github       *github.Client
	jobName      string
	db           db.DB
	pipelinestat *pipelinestat.CodePipeline
	region       string

	logger Logger
}

type ServiceOption func(*Service)

func SetLogger(logger Logger) ServiceOption {
	return func(svc *Service) {
		if logger != nil {
			svc.logger = logger
		} else {
			svc.logger = &defaultLogger{}
		}
	}
}

func NewService(
	github *github.Client,
	jobName string,
	db db.DB,
	pipelinestat *pipelinestat.CodePipeline,
	region string,
	ops ...ServiceOption,
) *Service {
	svc := &Service{
		github:       github,
		jobName:      jobName,
		db:           db,
		pipelinestat: pipelinestat,
		region:       region,

		logger: &defaultLogger{},
	}

	for _, f := range ops {
		f(svc)
	}

	return svc
}

type eventDetailType struct {
	Owner    string `json:"owner,omitempty"`
	Category string `json:"category,omitempty"`
	Provider string `json:"provider,omitempty"`
	Version  string `json:"version,omitempty"`
}

type eventDetail struct {
	Pipeline    string          `json:"pipeline,omitempty"`
	Version     json.Number     `json:"version,omitempty"`
	ExecutionId string          `json:"execution-id,omitempty"`
	Stage       string          `json:"stage,omitempty"`
	Action      string          `json:"action,omitempty"`
	State       string          `json:"state,omitempty"`
	Region      string          `json:"region,omitempty"`
	Type        eventDetailType `json:"type,omitempty"`
}

func (s *Service) handle(ctx context.Context, event *events.CloudWatchEvent) error {
	bs, err := event.Detail.MarshalJSON()
	if err != nil {
		s.logger.Error("event.Detail.MarshalJSON", err, "", 0, "")
		return err
	}
	detail := &eventDetail{}
	if err := json.Unmarshal(bs, detail); err != nil {
		s.logger.Error("event.Detail Unmarshal", err, "", 0, "")
		return err
	}
	execId := detail.ExecutionId

	stat, err := s.pipelinestat.GetStat(ctx, detail.Pipeline, execId)
	if err != nil {
		return err
	}

	ss := strings.Split(stat.RepoId, "/")
	if len(ss) != 2 {
		return errors.New("Invalid repoId in pipeline configuration")
	}
	owner, repo := ss[0], ss[1]

	crClient := checkrun.NewGithub(
		s.github,
		owner,
		repo,
		stat.BranchName,
		s.jobName,
	)

	p := NewDetailPrinter()
	text, err := p.Print(stat)
	if err != nil {
		return err
	}

	switch event.DetailType {
	case "CodePipeline Pipeline Execution State Change":
		state := detail.State
		switch state {
		case "STARTED":
			url := fmt.Sprintf("https://%s.console.aws.amazon.com/codesuite/codepipeline/pipelines/%s/view", s.region, detail.Pipeline)
			checkRunId, err := crClient.Create(
				ctx,
				"Deploy",
				detail.Pipeline,
				checkrun.SetDetailsURL(url),
				checkrun.SetOutputText(text),
			)
			if err != nil {
				s.logger.Error("CheckRunCreate", err, execId, 0, state)
				return err
			}

			if err := s.db.CreatePipeline(ctx, execId, checkRunId); err != nil {
				s.logger.Error("PipelineRecordPut", err, execId, checkRunId, state)
				return err
			}

		case "SUCCEEDED":
			checkRunId, err := s.db.GetCheckRunId(ctx, execId)
			if err != nil {
				s.logger.Error("CheckRunIdGet", err, execId, 0, state)
				return err
			}

			if err := crClient.Update(ctx, checkRunId, checkrun.Status_Success, text); err != nil {
				s.logger.Error("CheckRunSuccess", err, execId, checkRunId, state)
				return err
			}

		case "FAILED":
			checkRunId, err := s.db.GetCheckRunId(ctx, execId)
			if err != nil {
				s.logger.Error("CheckRunIdGet", err, execId, 0, state)
				return err
			}

			if err := crClient.Update(
				ctx,
				checkRunId,
				checkrun.Status_Failure,
				text,
			); err != nil {
				s.logger.Error("CheckRunFailure", err, execId, checkRunId, state)
				return err
			}

		case "CANCELED":
			checkRunId, err := s.db.GetCheckRunId(ctx, execId)
			if err != nil {
				s.logger.Error("CheckRunIdGet", err, execId, 0, state)
				return err
			}

			if err := crClient.Update(ctx, checkRunId, checkrun.Status_Cancelled, text); err != nil {
				s.logger.Error("CheckRunCancelled", err, execId, checkRunId, state)
				return err
			}

		default:
			s.logger.Error("UnknwonState", nil, execId, 0, state)
			return errors.New("UnknownState")
		}

	case "CodePipeline Action Execution State Change":
		checkRunId, err := s.db.GetCheckRunId(ctx, execId)
		if err != nil {
			s.logger.Error("CheckRunIdGet", err, execId, 0, detail.State)
			return err
		}

		if err := crClient.Update(ctx, checkRunId, checkrun.Status_InProgress, text); err != nil {
			s.logger.Error("CheckRun InProgress", err, execId, checkRunId, detail.State)
			return err
		}

	default:
		s.logger.Error("Unknwon Type", nil, execId, 0, detail.State)
		return errors.New("Unknown Type")
	}

	return nil
}

func (s *Service) Handle(ctx context.Context, event *events.CloudWatchEvent) error {
	s.logger.Request(event)

	if err := s.handle(ctx, event); err != nil {
		return err
	}

	return nil
}
