package elevenlabs

import (
	"context"
	"errors"

	"github.com/google/go-querystring/query"
)

type ConvaiTestsAPI interface {
	ListTests(ctx context.Context, req *ListTestsReq) (ListTestsResp, error)
	GetTest(ctx context.Context, req *GetTestReq) (UnitTestResponse, error)
	CreateTest(ctx context.Context, req *CreateTestReq) (CreateTestResp, error)
	UpdateTest(ctx context.Context, req *UpdateTestReq) (UnitTestResponse, error)
	DeleteTest(ctx context.Context, req *DeleteTestReq) error
	RunTests(ctx context.Context, req *RunTestsReq) (RunTestsResp, error)
	GetTestSummaries(ctx context.Context, req *GetTestSummariesReq) (GetTestSummariesResp, error)
}

// UnitTestType represents the type of unit test
type UnitTestType string

const (
	UnitTestTypeLLM  UnitTestType = "llm"
	UnitTestTypeTool UnitTestType = "tool"
)

// UnitTestSummary represents a test summary in list responses
type UnitTestSummary struct {
	Id                   string              `json:"id"`
	Name                 string              `json:"name"`
	AccessInfo           *ResourceAccessInfo `json:"access_info,omitempty"`
	CreatedAtUnixSecs    int64               `json:"created_at_unix_secs"`
	LastUpdatedAtUnixSecs int64               `json:"last_updated_at_unix_secs"`
	Type                 UnitTestType        `json:"type"`
}

// ListTestsReq represents the request for listing tests
type ListTestsReq struct {
	Cursor   *string `url:"cursor,omitempty"`
	PageSize int     `url:"page_size,omitempty"`
	Search   *string `url:"search,omitempty"`
}

func NewListTestsReq() *ListTestsReq {
	return &ListTestsReq{
		PageSize: 30,
	}
}

func (r ListTestsReq) QueryString() string {
	v, err := query.Values(r)
	if err != nil {
		return ""
	}
	qs := v.Encode()
	if qs == "" {
		return ""
	}
	return "?" + qs
}

// ListTestsResp represents the response from listing tests
type ListTestsResp struct {
	Tests      []UnitTestSummary `json:"tests"`
	NextCursor *string           `json:"next_cursor,omitempty"`
	HasMore    bool              `json:"has_more"`
}

// ListTests lists all agent response tests with pagination support.
// https://elevenlabs.io/docs/api-reference/tests/list
func (c *Client) ListTests(ctx context.Context, req *ListTestsReq) (ListTestsResp, error) {
	if req == nil {
		req = NewListTestsReq()
	}

	body, err := c.get(ctx, "/convai/agent-testing"+req.QueryString())
	if err != nil {
		return ListTestsResp{}, err
	}

	var resp ListTestsResp
	if err := c.parseResponse(body, &resp); err != nil {
		return ListTestsResp{}, err
	}

	return resp, nil
}

// GetTestReq represents the request for getting a test
type GetTestReq struct {
	TestId string `path:"test_id"`
}

func NewGetTestReq(testId string) *GetTestReq {
	return &GetTestReq{
		TestId: testId,
	}
}

// ConversationHistoryTranscript represents a conversation transcript entry
type ConversationHistoryTranscript struct {
	Role            TranscriptRole `json:"role"`
	AgentMetadata   *AgentMetadata `json:"agent_metadata,omitempty"`
	Message         *string        `json:"message,omitempty"`
	ToolCalls       []any          `json:"tool_calls,omitempty"`
	ToolResults     []any          `json:"tool_results,omitempty"`
	TimeInCallSecs  int            `json:"time_in_call_secs"`
	Interrupted     bool           `json:"interrupted"`
	OriginalMessage *string        `json:"original_message,omitempty"`
	SourceMedium    *SourceMedium  `json:"source_medium,omitempty"`
}

// AgentSuccessfulResponseExample represents a successful response example
type AgentSuccessfulResponseExample struct {
	Response string `json:"response"`
	Type     string `json:"type"`
}

// AgentFailureResponseExample represents a failure response example
type AgentFailureResponseExample struct {
	Response string `json:"response"`
	Type     string `json:"type"`
}

// ReferencedToolType represents the type of referenced tool
type ReferencedToolType string

const (
	ReferencedToolTypeSystem                ReferencedToolType = "system"
	ReferencedToolTypeWebhook               ReferencedToolType = "webhook"
	ReferencedToolTypeClient                ReferencedToolType = "client"
	ReferencedToolTypeWorkflow              ReferencedToolType = "workflow"
	ReferencedToolTypeApiIntegrationWebhook ReferencedToolType = "api_integration_webhook"
)

// ReferencedTool represents a referenced tool
type ReferencedTool struct {
	Id   string             `json:"id"`
	Type ReferencedToolType `json:"type"`
}

// UnitTestToolCallParameter represents a parameter evaluation
type UnitTestToolCallParameter struct {
	Eval any    `json:"eval"`
	Path string `json:"path"`
}

// UnitTestToolCallEvaluation represents tool call evaluation settings
type UnitTestToolCallEvaluation struct {
	Parameters     []UnitTestToolCallParameter `json:"parameters,omitempty"`
	ReferencedTool *ReferencedTool             `json:"referenced_tool,omitempty"`
	VerifyAbsence  bool                        `json:"verify_absence"`
}

// TestFromConversationMetadata represents metadata about the source conversation
type TestFromConversationMetadata struct {
	ConversationId     string                          `json:"conversation_id"`
	AgentId            string                          `json:"agent_id"`
	WorkflowNodeId     *string                         `json:"workflow_node_id,omitempty"`
	OriginalAgentReply []ConversationHistoryTranscript `json:"original_agent_reply,omitempty"`
}

// UnitTestResponse represents a full test response
type UnitTestResponse struct {
	Id                       string                          `json:"id"`
	Name                     string                          `json:"name"`
	ChatHistory              []ConversationHistoryTranscript `json:"chat_history"`
	SuccessCondition         string                          `json:"success_condition"`
	SuccessExamples          []AgentSuccessfulResponseExample `json:"success_examples"`
	FailureExamples          []AgentFailureResponseExample    `json:"failure_examples"`
	ToolCallParameters       *UnitTestToolCallEvaluation     `json:"tool_call_parameters,omitempty"`
	DynamicVariables         map[string]any                  `json:"dynamic_variables,omitempty"`
	Type                     UnitTestType                    `json:"type,omitempty"`
	FromConversationMetadata *TestFromConversationMetadata   `json:"from_conversation_metadata,omitempty"`
}

// GetTest gets an agent response test by ID.
// https://elevenlabs.io/docs/api-reference/tests/get
func (c *Client) GetTest(ctx context.Context, req *GetTestReq) (UnitTestResponse, error) {
	if req == nil {
		return UnitTestResponse{}, errors.New("request is nil")
	}

	body, err := c.get(ctx, "/convai/agent-testing/"+req.TestId)
	if err != nil {
		return UnitTestResponse{}, err
	}

	var resp UnitTestResponse
	if err := c.parseResponse(body, &resp); err != nil {
		return UnitTestResponse{}, err
	}

	return resp, nil
}

// CreateTestReq represents the request for creating a test
type CreateTestReq struct {
	Name                     string                           `json:"name"`
	ChatHistory              []ConversationHistoryTranscript  `json:"chat_history"`
	SuccessCondition         string                           `json:"success_condition"`
	SuccessExamples          []AgentSuccessfulResponseExample `json:"success_examples"`
	FailureExamples          []AgentFailureResponseExample    `json:"failure_examples"`
	ToolCallParameters       *UnitTestToolCallEvaluation      `json:"tool_call_parameters,omitempty"`
	DynamicVariables         map[string]any                   `json:"dynamic_variables,omitempty"`
	Type                     UnitTestType                     `json:"type,omitempty"`
	FromConversationMetadata *TestFromConversationMetadata    `json:"from_conversation_metadata,omitempty"`
}

func NewCreateTestReq(name, successCondition string) *CreateTestReq {
	return &CreateTestReq{
		Name:             name,
		SuccessCondition: successCondition,
		ChatHistory:      []ConversationHistoryTranscript{},
		SuccessExamples:  []AgentSuccessfulResponseExample{},
		FailureExamples:  []AgentFailureResponseExample{},
	}
}

// CreateTestResp represents the response from creating a test
type CreateTestResp struct {
	Id string `json:"id"`
}

// CreateTest creates a new agent response test.
// https://elevenlabs.io/docs/api-reference/tests/create
func (c *Client) CreateTest(ctx context.Context, req *CreateTestReq) (CreateTestResp, error) {
	if req == nil {
		return CreateTestResp{}, errors.New("request is nil")
	}

	body, err := c.post(ctx, "/convai/agent-testing/create", req)
	if err != nil {
		return CreateTestResp{}, err
	}

	var resp CreateTestResp
	if err := c.parseResponse(body, &resp); err != nil {
		return CreateTestResp{}, err
	}

	return resp, nil
}

// UpdateTestReq represents the request for updating a test
type UpdateTestReq struct {
	TestId                   string                           `path:"test_id"`
	Name                     string                           `json:"name"`
	ChatHistory              []ConversationHistoryTranscript  `json:"chat_history"`
	SuccessCondition         string                           `json:"success_condition"`
	SuccessExamples          []AgentSuccessfulResponseExample `json:"success_examples"`
	FailureExamples          []AgentFailureResponseExample    `json:"failure_examples"`
	ToolCallParameters       *UnitTestToolCallEvaluation      `json:"tool_call_parameters,omitempty"`
	DynamicVariables         map[string]any                   `json:"dynamic_variables,omitempty"`
	Type                     UnitTestType                     `json:"type,omitempty"`
	FromConversationMetadata *TestFromConversationMetadata    `json:"from_conversation_metadata,omitempty"`
}

func NewUpdateTestReq(testId, name, successCondition string) *UpdateTestReq {
	return &UpdateTestReq{
		TestId:           testId,
		Name:             name,
		SuccessCondition: successCondition,
		ChatHistory:      []ConversationHistoryTranscript{},
		SuccessExamples:  []AgentSuccessfulResponseExample{},
		FailureExamples:  []AgentFailureResponseExample{},
	}
}

// UpdateTest updates an agent response test by ID.
// https://elevenlabs.io/docs/api-reference/tests/update
func (c *Client) UpdateTest(ctx context.Context, req *UpdateTestReq) (UnitTestResponse, error) {
	if req == nil {
		return UnitTestResponse{}, errors.New("request is nil")
	}

	body, err := c.put(ctx, "/convai/agent-testing/"+req.TestId, req)
	if err != nil {
		return UnitTestResponse{}, err
	}

	var resp UnitTestResponse
	if err := c.parseResponse(body, &resp); err != nil {
		return UnitTestResponse{}, err
	}

	return resp, nil
}

// DeleteTestReq represents the request for deleting a test
type DeleteTestReq struct {
	TestId string `path:"test_id"`
}

func NewDeleteTestReq(testId string) *DeleteTestReq {
	return &DeleteTestReq{
		TestId: testId,
	}
}

// DeleteTest deletes an agent response test by ID.
// https://elevenlabs.io/docs/api-reference/tests/delete
func (c *Client) DeleteTest(ctx context.Context, req *DeleteTestReq) error {
	if req == nil {
		return errors.New("request is nil")
	}

	return c.delete(ctx, "/convai/agent-testing/"+req.TestId)
}

// RunTestsReq represents the request for running tests
type RunTestsReq struct {
	AgentId  string   `path:"agent_id"`
	TestIds  []string `json:"test_ids,omitempty"`
	BranchId *string  `json:"branch_id,omitempty"`
}

func NewRunTestsReq(agentId string) *RunTestsReq {
	return &RunTestsReq{
		AgentId: agentId,
	}
}

// TestResultStatus represents the status of a test result
type TestResultStatus string

const (
	TestResultStatusPassed TestResultStatus = "passed"
	TestResultStatusFailed TestResultStatus = "failed"
	TestResultStatusError  TestResultStatus = "error"
)

// TestResult represents a single test result
type TestResult struct {
	TestId         string           `json:"test_id"`
	TestName       string           `json:"test_name"`
	Status         TestResultStatus `json:"status"`
	AgentResponse  *string          `json:"agent_response,omitempty"`
	ErrorMessage   *string          `json:"error_message,omitempty"`
	EvaluationText *string          `json:"evaluation_text,omitempty"`
}

// RunTestsResp represents the response from running tests
type RunTestsResp struct {
	Results []TestResult `json:"results"`
}

// RunTests runs agent response tests against an agent.
// https://elevenlabs.io/docs/api-reference/agents/run-tests
func (c *Client) RunTests(ctx context.Context, req *RunTestsReq) (RunTestsResp, error) {
	if req == nil {
		return RunTestsResp{}, errors.New("request is nil")
	}

	body, err := c.post(ctx, "/convai/agents/"+req.AgentId+"/run-tests", req)
	if err != nil {
		return RunTestsResp{}, err
	}

	var resp RunTestsResp
	if err := c.parseResponse(body, &resp); err != nil {
		return RunTestsResp{}, err
	}

	return resp, nil
}

// GetTestSummariesReq represents the request for getting test summaries by IDs
type GetTestSummariesReq struct {
	TestIds []string `json:"test_ids"`
}

func NewGetTestSummariesReq(testIds []string) *GetTestSummariesReq {
	return &GetTestSummariesReq{
		TestIds: testIds,
	}
}

// GetTestSummariesResp represents the response for getting test summaries
type GetTestSummariesResp struct {
	Tests map[string]UnitTestSummary `json:"tests"`
}

// GetTestSummaries gets multiple agent response tests by their IDs.
// https://elevenlabs.io/docs/api-reference/tests/summaries
func (c *Client) GetTestSummaries(ctx context.Context, req *GetTestSummariesReq) (GetTestSummariesResp, error) {
	if req == nil {
		return GetTestSummariesResp{}, errors.New("request is nil")
	}

	body, err := c.post(ctx, "/convai/agent-testing/summaries", req)
	if err != nil {
		return GetTestSummariesResp{}, err
	}

	var resp GetTestSummariesResp
	if err := c.parseResponse(body, &resp); err != nil {
		return GetTestSummariesResp{}, err
	}

	return resp, nil
}
