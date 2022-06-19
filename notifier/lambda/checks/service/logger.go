package service

import (
	"github.com/aws/aws-lambda-go/events"
)

type Logger interface {
	Request(event *events.CloudWatchEvent)
	Error(method string, err error, execId string, checkRunId int64, state string)
}

type defaultLogger struct{}

var _ Logger = (*defaultLogger)(nil)

func (d *defaultLogger) Request(*events.CloudWatchEvent)            {}
func (d *defaultLogger) Error(string, error, string, int64, string) {}
