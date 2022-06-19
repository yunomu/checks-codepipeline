package checkrun

import (
	"context"

	"github.com/google/go-github/v45/github"
)

type Status string

const (
	Status_InProgress Status = "in_progress"
	Status_Success    Status = "success"
	Status_Failure    Status = "failure"
	Status_Cancelled  Status = "cancelled"
)

type CreateOption func(*github.CreateCheckRunOptions)

func SetDetailsURL(url string) CreateOption {
	return func(o *github.CreateCheckRunOptions) {
		o.DetailsURL = &url
	}
}

func SetOutputText(text string) CreateOption {
	return func(o *github.CreateCheckRunOptions) {
		o.Output.Text = &text
	}
}

type CheckRun interface {
	Create(ctx context.Context, title, summary string, opts ...CreateOption) (int64, error)
	Update(ctx context.Context, checkRunID int64, status Status, details ...string) error
}
