package pipelinestat

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/codepipeline"
)

type CodePipeline struct {
	client *codepipeline.CodePipeline
}

func NewCodePipeline(client *codepipeline.CodePipeline) *CodePipeline {
	return &CodePipeline{
		client: client,
	}
}

type ActionStat struct {
	Name        string
	ExecutionId string
	Status      string
	ErrorDetail string
}

type StageStat struct {
	Name        string
	ActionStats []ActionStat
}

type Stat struct {
	RepoId      string
	BranchName  string
	ExecutionId string
	StageStats  []StageStat
}

func findAction(stat *Stat, stageName, actionName string) *ActionStat {
	for i, s := range stat.StageStats {
		if s.Name != stageName {
			continue
		}

		for j, a := range stat.StageStats[i].ActionStats {
			if a.Name != actionName {
				continue
			}

			return &stat.StageStats[i].ActionStats[j]
		}
	}

	return nil
}

func (c *CodePipeline) GetStat(ctx context.Context, name string, execId string) (*Stat, error) {
	var stat Stat
	out, err := c.client.GetPipelineWithContext(ctx, &codepipeline.GetPipelineInput{
		Name: aws.String(name),
	})
	if err != nil {
		return nil, err
	}

	if len(out.Pipeline.Stages) == 0 || aws.StringValue(out.Pipeline.Stages[0].Name) != "SourceStage" {
		return nil, errors.New("SourceStage is not found")
	}
	sourceStage := out.Pipeline.Stages[0]

	if len(sourceStage.Actions) == 0 || aws.StringValue(sourceStage.Actions[0].Name) != "Source" {
		return nil, errors.New("SourceAction is not found")
	}
	sourceAction := sourceStage.Actions[0]

	srcCfg := sourceAction.Configuration
	stat.RepoId = aws.StringValue(srcCfg["FullRepositoryId"])
	stat.BranchName = aws.StringValue(srcCfg["BranchName"])

	stat.ExecutionId = execId
	stat.StageStats = make([]StageStat, len(out.Pipeline.Stages))
	for i, s := range out.Pipeline.Stages {
		stage := &stat.StageStats[i]
		stage.Name = aws.StringValue(s.Name)
		stage.ActionStats = make([]ActionStat, len(s.Actions))

		for j, a := range s.Actions {
			action := &stage.ActionStats[j]
			action.Name = aws.StringValue(a.Name)
		}
	}

	var rerr error
	if err := c.client.ListActionExecutionsPagesWithContext(ctx, &codepipeline.ListActionExecutionsInput{
		PipelineName: aws.String(name),
		Filter: &codepipeline.ActionExecutionFilter{
			PipelineExecutionId: aws.String(execId),
		},
	}, func(out *codepipeline.ListActionExecutionsOutput, last bool) bool {
		for _, a := range out.ActionExecutionDetails {
			action := findAction(&stat, aws.StringValue(a.StageName), aws.StringValue(a.ActionName))
			if action == nil {
				continue
			}

			action.ExecutionId = aws.StringValue(a.ActionExecutionId)
			action.Status = aws.StringValue(a.Status)
			if o := a.Output; o != nil && o.ExecutionResult != nil {
				action.ErrorDetail = aws.StringValue(o.ExecutionResult.ExternalExecutionSummary)
			}
		}

		return true
	}); err != nil {
		return nil, err
	} else if rerr != nil {
		return nil, rerr
	}

	return &stat, nil
}
