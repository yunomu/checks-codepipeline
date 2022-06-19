package service

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/codepipeline"
)

type mockCodePipeline struct {
	GetJobDetailsFunc func(*codepipeline.GetJobDetailsInput) (*codepipeline.GetJobDetailsOutput, error)
}

func (m *mockCodePipeline) AcknowledgeJob(_ *codepipeline.AcknowledgeJobInput) (*codepipeline.AcknowledgeJobOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) AcknowledgeJobWithContext(_ aws.Context, _ *codepipeline.AcknowledgeJobInput, _ ...request.Option) (*codepipeline.AcknowledgeJobOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) AcknowledgeJobRequest(_ *codepipeline.AcknowledgeJobInput) (*request.Request, *codepipeline.AcknowledgeJobOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) AcknowledgeThirdPartyJob(_ *codepipeline.AcknowledgeThirdPartyJobInput) (*codepipeline.AcknowledgeThirdPartyJobOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) AcknowledgeThirdPartyJobWithContext(_ aws.Context, _ *codepipeline.AcknowledgeThirdPartyJobInput, _ ...request.Option) (*codepipeline.AcknowledgeThirdPartyJobOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) AcknowledgeThirdPartyJobRequest(_ *codepipeline.AcknowledgeThirdPartyJobInput) (*request.Request, *codepipeline.AcknowledgeThirdPartyJobOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) CreateCustomActionType(_ *codepipeline.CreateCustomActionTypeInput) (*codepipeline.CreateCustomActionTypeOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) CreateCustomActionTypeWithContext(_ aws.Context, _ *codepipeline.CreateCustomActionTypeInput, _ ...request.Option) (*codepipeline.CreateCustomActionTypeOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) CreateCustomActionTypeRequest(_ *codepipeline.CreateCustomActionTypeInput) (*request.Request, *codepipeline.CreateCustomActionTypeOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) CreatePipeline(_ *codepipeline.CreatePipelineInput) (*codepipeline.CreatePipelineOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) CreatePipelineWithContext(_ aws.Context, _ *codepipeline.CreatePipelineInput, _ ...request.Option) (*codepipeline.CreatePipelineOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) CreatePipelineRequest(_ *codepipeline.CreatePipelineInput) (*request.Request, *codepipeline.CreatePipelineOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) DeleteCustomActionType(_ *codepipeline.DeleteCustomActionTypeInput) (*codepipeline.DeleteCustomActionTypeOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) DeleteCustomActionTypeWithContext(_ aws.Context, _ *codepipeline.DeleteCustomActionTypeInput, _ ...request.Option) (*codepipeline.DeleteCustomActionTypeOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) DeleteCustomActionTypeRequest(_ *codepipeline.DeleteCustomActionTypeInput) (*request.Request, *codepipeline.DeleteCustomActionTypeOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) DeletePipeline(_ *codepipeline.DeletePipelineInput) (*codepipeline.DeletePipelineOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) DeletePipelineWithContext(_ aws.Context, _ *codepipeline.DeletePipelineInput, _ ...request.Option) (*codepipeline.DeletePipelineOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) DeletePipelineRequest(_ *codepipeline.DeletePipelineInput) (*request.Request, *codepipeline.DeletePipelineOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) DeleteWebhook(_ *codepipeline.DeleteWebhookInput) (*codepipeline.DeleteWebhookOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) DeleteWebhookWithContext(_ aws.Context, _ *codepipeline.DeleteWebhookInput, _ ...request.Option) (*codepipeline.DeleteWebhookOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) DeleteWebhookRequest(_ *codepipeline.DeleteWebhookInput) (*request.Request, *codepipeline.DeleteWebhookOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) DeregisterWebhookWithThirdParty(_ *codepipeline.DeregisterWebhookWithThirdPartyInput) (*codepipeline.DeregisterWebhookWithThirdPartyOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) DeregisterWebhookWithThirdPartyWithContext(_ aws.Context, _ *codepipeline.DeregisterWebhookWithThirdPartyInput, _ ...request.Option) (*codepipeline.DeregisterWebhookWithThirdPartyOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) DeregisterWebhookWithThirdPartyRequest(_ *codepipeline.DeregisterWebhookWithThirdPartyInput) (*request.Request, *codepipeline.DeregisterWebhookWithThirdPartyOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) DisableStageTransition(_ *codepipeline.DisableStageTransitionInput) (*codepipeline.DisableStageTransitionOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) DisableStageTransitionWithContext(_ aws.Context, _ *codepipeline.DisableStageTransitionInput, _ ...request.Option) (*codepipeline.DisableStageTransitionOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) DisableStageTransitionRequest(_ *codepipeline.DisableStageTransitionInput) (*request.Request, *codepipeline.DisableStageTransitionOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) EnableStageTransition(_ *codepipeline.EnableStageTransitionInput) (*codepipeline.EnableStageTransitionOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) EnableStageTransitionWithContext(_ aws.Context, _ *codepipeline.EnableStageTransitionInput, _ ...request.Option) (*codepipeline.EnableStageTransitionOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) EnableStageTransitionRequest(_ *codepipeline.EnableStageTransitionInput) (*request.Request, *codepipeline.EnableStageTransitionOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) GetActionType(_ *codepipeline.GetActionTypeInput) (*codepipeline.GetActionTypeOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) GetActionTypeWithContext(_ aws.Context, _ *codepipeline.GetActionTypeInput, _ ...request.Option) (*codepipeline.GetActionTypeOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) GetActionTypeRequest(_ *codepipeline.GetActionTypeInput) (*request.Request, *codepipeline.GetActionTypeOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) GetJobDetails(in *codepipeline.GetJobDetailsInput) (*codepipeline.GetJobDetailsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) GetJobDetailsWithContext(_ aws.Context, in *codepipeline.GetJobDetailsInput, _ ...request.Option) (*codepipeline.GetJobDetailsOutput, error) {
	if f := m.GetJobDetailsFunc; f != nil {
		return f(in)
	}
	panic("not implemented")
}

func (m *mockCodePipeline) GetJobDetailsRequest(_ *codepipeline.GetJobDetailsInput) (*request.Request, *codepipeline.GetJobDetailsOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) GetPipeline(_ *codepipeline.GetPipelineInput) (*codepipeline.GetPipelineOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) GetPipelineWithContext(_ aws.Context, _ *codepipeline.GetPipelineInput, _ ...request.Option) (*codepipeline.GetPipelineOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) GetPipelineRequest(_ *codepipeline.GetPipelineInput) (*request.Request, *codepipeline.GetPipelineOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) GetPipelineExecution(_ *codepipeline.GetPipelineExecutionInput) (*codepipeline.GetPipelineExecutionOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) GetPipelineExecutionWithContext(_ aws.Context, _ *codepipeline.GetPipelineExecutionInput, _ ...request.Option) (*codepipeline.GetPipelineExecutionOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) GetPipelineExecutionRequest(_ *codepipeline.GetPipelineExecutionInput) (*request.Request, *codepipeline.GetPipelineExecutionOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) GetPipelineState(_ *codepipeline.GetPipelineStateInput) (*codepipeline.GetPipelineStateOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) GetPipelineStateWithContext(_ aws.Context, _ *codepipeline.GetPipelineStateInput, _ ...request.Option) (*codepipeline.GetPipelineStateOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) GetPipelineStateRequest(_ *codepipeline.GetPipelineStateInput) (*request.Request, *codepipeline.GetPipelineStateOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) GetThirdPartyJobDetails(_ *codepipeline.GetThirdPartyJobDetailsInput) (*codepipeline.GetThirdPartyJobDetailsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) GetThirdPartyJobDetailsWithContext(_ aws.Context, _ *codepipeline.GetThirdPartyJobDetailsInput, _ ...request.Option) (*codepipeline.GetThirdPartyJobDetailsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) GetThirdPartyJobDetailsRequest(_ *codepipeline.GetThirdPartyJobDetailsInput) (*request.Request, *codepipeline.GetThirdPartyJobDetailsOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) ListActionExecutions(_ *codepipeline.ListActionExecutionsInput) (*codepipeline.ListActionExecutionsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) ListActionExecutionsWithContext(_ aws.Context, _ *codepipeline.ListActionExecutionsInput, _ ...request.Option) (*codepipeline.ListActionExecutionsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) ListActionExecutionsRequest(_ *codepipeline.ListActionExecutionsInput) (*request.Request, *codepipeline.ListActionExecutionsOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) ListActionExecutionsPages(_ *codepipeline.ListActionExecutionsInput, _ func(*codepipeline.ListActionExecutionsOutput, bool) bool) error {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) ListActionExecutionsPagesWithContext(_ aws.Context, _ *codepipeline.ListActionExecutionsInput, _ func(*codepipeline.ListActionExecutionsOutput, bool) bool, _ ...request.Option) error {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) ListActionTypes(_ *codepipeline.ListActionTypesInput) (*codepipeline.ListActionTypesOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) ListActionTypesWithContext(_ aws.Context, _ *codepipeline.ListActionTypesInput, _ ...request.Option) (*codepipeline.ListActionTypesOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) ListActionTypesRequest(_ *codepipeline.ListActionTypesInput) (*request.Request, *codepipeline.ListActionTypesOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) ListActionTypesPages(_ *codepipeline.ListActionTypesInput, _ func(*codepipeline.ListActionTypesOutput, bool) bool) error {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) ListActionTypesPagesWithContext(_ aws.Context, _ *codepipeline.ListActionTypesInput, _ func(*codepipeline.ListActionTypesOutput, bool) bool, _ ...request.Option) error {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) ListPipelineExecutions(_ *codepipeline.ListPipelineExecutionsInput) (*codepipeline.ListPipelineExecutionsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) ListPipelineExecutionsWithContext(_ aws.Context, _ *codepipeline.ListPipelineExecutionsInput, _ ...request.Option) (*codepipeline.ListPipelineExecutionsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) ListPipelineExecutionsRequest(_ *codepipeline.ListPipelineExecutionsInput) (*request.Request, *codepipeline.ListPipelineExecutionsOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) ListPipelineExecutionsPages(_ *codepipeline.ListPipelineExecutionsInput, _ func(*codepipeline.ListPipelineExecutionsOutput, bool) bool) error {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) ListPipelineExecutionsPagesWithContext(_ aws.Context, _ *codepipeline.ListPipelineExecutionsInput, _ func(*codepipeline.ListPipelineExecutionsOutput, bool) bool, _ ...request.Option) error {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) ListPipelines(_ *codepipeline.ListPipelinesInput) (*codepipeline.ListPipelinesOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) ListPipelinesWithContext(_ aws.Context, _ *codepipeline.ListPipelinesInput, _ ...request.Option) (*codepipeline.ListPipelinesOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) ListPipelinesRequest(_ *codepipeline.ListPipelinesInput) (*request.Request, *codepipeline.ListPipelinesOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) ListPipelinesPages(_ *codepipeline.ListPipelinesInput, _ func(*codepipeline.ListPipelinesOutput, bool) bool) error {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) ListPipelinesPagesWithContext(_ aws.Context, _ *codepipeline.ListPipelinesInput, _ func(*codepipeline.ListPipelinesOutput, bool) bool, _ ...request.Option) error {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) ListTagsForResource(_ *codepipeline.ListTagsForResourceInput) (*codepipeline.ListTagsForResourceOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) ListTagsForResourceWithContext(_ aws.Context, _ *codepipeline.ListTagsForResourceInput, _ ...request.Option) (*codepipeline.ListTagsForResourceOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) ListTagsForResourceRequest(_ *codepipeline.ListTagsForResourceInput) (*request.Request, *codepipeline.ListTagsForResourceOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) ListTagsForResourcePages(_ *codepipeline.ListTagsForResourceInput, _ func(*codepipeline.ListTagsForResourceOutput, bool) bool) error {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) ListTagsForResourcePagesWithContext(_ aws.Context, _ *codepipeline.ListTagsForResourceInput, _ func(*codepipeline.ListTagsForResourceOutput, bool) bool, _ ...request.Option) error {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) ListWebhooks(_ *codepipeline.ListWebhooksInput) (*codepipeline.ListWebhooksOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) ListWebhooksWithContext(_ aws.Context, _ *codepipeline.ListWebhooksInput, _ ...request.Option) (*codepipeline.ListWebhooksOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) ListWebhooksRequest(_ *codepipeline.ListWebhooksInput) (*request.Request, *codepipeline.ListWebhooksOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) ListWebhooksPages(_ *codepipeline.ListWebhooksInput, _ func(*codepipeline.ListWebhooksOutput, bool) bool) error {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) ListWebhooksPagesWithContext(_ aws.Context, _ *codepipeline.ListWebhooksInput, _ func(*codepipeline.ListWebhooksOutput, bool) bool, _ ...request.Option) error {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) PollForJobs(_ *codepipeline.PollForJobsInput) (*codepipeline.PollForJobsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) PollForJobsWithContext(_ aws.Context, _ *codepipeline.PollForJobsInput, _ ...request.Option) (*codepipeline.PollForJobsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) PollForJobsRequest(_ *codepipeline.PollForJobsInput) (*request.Request, *codepipeline.PollForJobsOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) PollForThirdPartyJobs(_ *codepipeline.PollForThirdPartyJobsInput) (*codepipeline.PollForThirdPartyJobsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) PollForThirdPartyJobsWithContext(_ aws.Context, _ *codepipeline.PollForThirdPartyJobsInput, _ ...request.Option) (*codepipeline.PollForThirdPartyJobsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) PollForThirdPartyJobsRequest(_ *codepipeline.PollForThirdPartyJobsInput) (*request.Request, *codepipeline.PollForThirdPartyJobsOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) PutActionRevision(_ *codepipeline.PutActionRevisionInput) (*codepipeline.PutActionRevisionOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) PutActionRevisionWithContext(_ aws.Context, _ *codepipeline.PutActionRevisionInput, _ ...request.Option) (*codepipeline.PutActionRevisionOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) PutActionRevisionRequest(_ *codepipeline.PutActionRevisionInput) (*request.Request, *codepipeline.PutActionRevisionOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) PutApprovalResult(_ *codepipeline.PutApprovalResultInput) (*codepipeline.PutApprovalResultOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) PutApprovalResultWithContext(_ aws.Context, _ *codepipeline.PutApprovalResultInput, _ ...request.Option) (*codepipeline.PutApprovalResultOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) PutApprovalResultRequest(_ *codepipeline.PutApprovalResultInput) (*request.Request, *codepipeline.PutApprovalResultOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) PutJobFailureResult(_ *codepipeline.PutJobFailureResultInput) (*codepipeline.PutJobFailureResultOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) PutJobFailureResultWithContext(_ aws.Context, _ *codepipeline.PutJobFailureResultInput, _ ...request.Option) (*codepipeline.PutJobFailureResultOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) PutJobFailureResultRequest(_ *codepipeline.PutJobFailureResultInput) (*request.Request, *codepipeline.PutJobFailureResultOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) PutJobSuccessResult(_ *codepipeline.PutJobSuccessResultInput) (*codepipeline.PutJobSuccessResultOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) PutJobSuccessResultWithContext(_ aws.Context, _ *codepipeline.PutJobSuccessResultInput, _ ...request.Option) (*codepipeline.PutJobSuccessResultOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) PutJobSuccessResultRequest(_ *codepipeline.PutJobSuccessResultInput) (*request.Request, *codepipeline.PutJobSuccessResultOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) PutThirdPartyJobFailureResult(_ *codepipeline.PutThirdPartyJobFailureResultInput) (*codepipeline.PutThirdPartyJobFailureResultOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) PutThirdPartyJobFailureResultWithContext(_ aws.Context, _ *codepipeline.PutThirdPartyJobFailureResultInput, _ ...request.Option) (*codepipeline.PutThirdPartyJobFailureResultOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) PutThirdPartyJobFailureResultRequest(_ *codepipeline.PutThirdPartyJobFailureResultInput) (*request.Request, *codepipeline.PutThirdPartyJobFailureResultOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) PutThirdPartyJobSuccessResult(_ *codepipeline.PutThirdPartyJobSuccessResultInput) (*codepipeline.PutThirdPartyJobSuccessResultOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) PutThirdPartyJobSuccessResultWithContext(_ aws.Context, _ *codepipeline.PutThirdPartyJobSuccessResultInput, _ ...request.Option) (*codepipeline.PutThirdPartyJobSuccessResultOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) PutThirdPartyJobSuccessResultRequest(_ *codepipeline.PutThirdPartyJobSuccessResultInput) (*request.Request, *codepipeline.PutThirdPartyJobSuccessResultOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) PutWebhook(_ *codepipeline.PutWebhookInput) (*codepipeline.PutWebhookOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) PutWebhookWithContext(_ aws.Context, _ *codepipeline.PutWebhookInput, _ ...request.Option) (*codepipeline.PutWebhookOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) PutWebhookRequest(_ *codepipeline.PutWebhookInput) (*request.Request, *codepipeline.PutWebhookOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) RegisterWebhookWithThirdParty(_ *codepipeline.RegisterWebhookWithThirdPartyInput) (*codepipeline.RegisterWebhookWithThirdPartyOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) RegisterWebhookWithThirdPartyWithContext(_ aws.Context, _ *codepipeline.RegisterWebhookWithThirdPartyInput, _ ...request.Option) (*codepipeline.RegisterWebhookWithThirdPartyOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) RegisterWebhookWithThirdPartyRequest(_ *codepipeline.RegisterWebhookWithThirdPartyInput) (*request.Request, *codepipeline.RegisterWebhookWithThirdPartyOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) RetryStageExecution(_ *codepipeline.RetryStageExecutionInput) (*codepipeline.RetryStageExecutionOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) RetryStageExecutionWithContext(_ aws.Context, _ *codepipeline.RetryStageExecutionInput, _ ...request.Option) (*codepipeline.RetryStageExecutionOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) RetryStageExecutionRequest(_ *codepipeline.RetryStageExecutionInput) (*request.Request, *codepipeline.RetryStageExecutionOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) StartPipelineExecution(_ *codepipeline.StartPipelineExecutionInput) (*codepipeline.StartPipelineExecutionOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) StartPipelineExecutionWithContext(_ aws.Context, _ *codepipeline.StartPipelineExecutionInput, _ ...request.Option) (*codepipeline.StartPipelineExecutionOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) StartPipelineExecutionRequest(_ *codepipeline.StartPipelineExecutionInput) (*request.Request, *codepipeline.StartPipelineExecutionOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) StopPipelineExecution(_ *codepipeline.StopPipelineExecutionInput) (*codepipeline.StopPipelineExecutionOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) StopPipelineExecutionWithContext(_ aws.Context, _ *codepipeline.StopPipelineExecutionInput, _ ...request.Option) (*codepipeline.StopPipelineExecutionOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) StopPipelineExecutionRequest(_ *codepipeline.StopPipelineExecutionInput) (*request.Request, *codepipeline.StopPipelineExecutionOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) TagResource(_ *codepipeline.TagResourceInput) (*codepipeline.TagResourceOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) TagResourceWithContext(_ aws.Context, _ *codepipeline.TagResourceInput, _ ...request.Option) (*codepipeline.TagResourceOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) TagResourceRequest(_ *codepipeline.TagResourceInput) (*request.Request, *codepipeline.TagResourceOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) UntagResource(_ *codepipeline.UntagResourceInput) (*codepipeline.UntagResourceOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) UntagResourceWithContext(_ aws.Context, _ *codepipeline.UntagResourceInput, _ ...request.Option) (*codepipeline.UntagResourceOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) UntagResourceRequest(_ *codepipeline.UntagResourceInput) (*request.Request, *codepipeline.UntagResourceOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) UpdateActionType(_ *codepipeline.UpdateActionTypeInput) (*codepipeline.UpdateActionTypeOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) UpdateActionTypeWithContext(_ aws.Context, _ *codepipeline.UpdateActionTypeInput, _ ...request.Option) (*codepipeline.UpdateActionTypeOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) UpdateActionTypeRequest(_ *codepipeline.UpdateActionTypeInput) (*request.Request, *codepipeline.UpdateActionTypeOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) UpdatePipeline(_ *codepipeline.UpdatePipelineInput) (*codepipeline.UpdatePipelineOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) UpdatePipelineWithContext(_ aws.Context, _ *codepipeline.UpdatePipelineInput, _ ...request.Option) (*codepipeline.UpdatePipelineOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *mockCodePipeline) UpdatePipelineRequest(_ *codepipeline.UpdatePipelineInput) (*request.Request, *codepipeline.UpdatePipelineOutput) {
	panic("not implemented") // TODO: Implement
}
