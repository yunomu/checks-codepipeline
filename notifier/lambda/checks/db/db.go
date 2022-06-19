package db

import (
	"context"
	"errors"
)

var (
	ErrNotExists = errors.New("Value not exists")
)

type DB interface {
	CreatePipeline(ctx context.Context, executionId string, checkRunId int64) error
	GetCheckRunId(ctx context.Context, executionId string) (int64, error)
}
