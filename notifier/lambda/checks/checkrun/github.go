package checkrun

import (
	"context"
	"errors"
	"strings"

	"github.com/google/go-github/v45/github"
)

type Github struct {
	client *github.Client
	owner  string
	repo   string
	branch string
	name   string
}

var _ CheckRun = (*Github)(nil)

func NewGithub(
	client *github.Client,
	owner string,
	repo string,
	branch string,
	name string,
) *Github {
	return &Github{
		client: client,
		owner:  owner,
		repo:   repo,
		branch: branch,
		name:   name,
	}
}

func (g *Github) Create(ctx context.Context, title, summary string, opts ...CreateOption) (int64, error) {
	branch, _, err := g.client.Repositories.GetBranch(ctx, g.owner, g.repo, g.branch, true)
	if err != nil {
		return 0, err
	}

	status := "in_progress"
	options := github.CreateCheckRunOptions{
		Name:    g.name,
		HeadSHA: branch.GetCommit().GetSHA(),
		Status:  &status,
		Output: &github.CheckRunOutput{
			Title:   &title,
			Summary: &summary,
		},
	}
	for _, f := range opts {
		f(&options)
	}

	checkRun, _, err := g.client.Checks.CreateCheckRun(ctx, g.owner, g.repo, options)
	if err != nil {
		return 0, err
	}

	return checkRun.GetID(), nil
}

func (g *Github) getCheckRun(ctx context.Context, checkRunID int64) (*github.CheckRun, error) {
	ret, _, err := g.client.Checks.GetCheckRun(ctx, g.owner, g.repo, checkRunID)

	return ret, err
}

func getTitleSummary(run *github.CheckRun) (*string, *string, *string) {
	out := run.GetOutput()
	if out == nil {
		return nil, nil, nil
	}
	return out.Title, out.Summary, out.Text
}

func (g *Github) Update(ctx context.Context, checkRunID int64, st Status, details ...string) error {
	checkRun, err := g.getCheckRun(ctx, checkRunID)
	if err != nil {
		return err
	}

	var status, conclusion *string
	switch st {
	case Status_Success, Status_Failure, Status_Cancelled:
		s := "completed"
		status = &s
		c := string(st)
		conclusion = &c
	case Status_InProgress:
		s := string(st)
		status = &s
	default:
		return errors.New("Unknown status: " + string(st))
	}

	title, summary, text := getTitleSummary(checkRun)
	if len(details) != 0 {
		t := strings.Join(details, "<br>")
		text = &t
	}

	if _, _, err := g.client.Checks.UpdateCheckRun(ctx, g.owner, g.repo, checkRunID, github.UpdateCheckRunOptions{
		Name:       g.name,
		Status:     status,
		Conclusion: conclusion,
		Output: &github.CheckRunOutput{
			Title:   title,
			Summary: summary,
			Text:    text,
		},
	}); err != nil {
		return err
	}

	return nil
}
