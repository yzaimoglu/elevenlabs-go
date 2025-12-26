package elevenlabs

import (
	"context"
	"errors"

	"github.com/google/go-querystring/query"
)

type ConvaiToolsAPI interface {
	ListTools(ctx context.Context) (ListToolsResp, error)
	GetTool(ctx context.Context, req *GetToolReq) (ToolResponseModel, error)
	CreateTool(ctx context.Context, req *CreateToolReq) (ToolResponseModel, error)
	UpdateTool(ctx context.Context, req *UpdateToolReq) (ToolResponseModel, error)
	DeleteTool(ctx context.Context, req *DeleteToolReq) error
	GetToolDependentAgents(ctx context.Context, req *GetToolDependentAgentsReq) (GetToolDependentAgentsResp, error)
}

// ConvaiToolType represents the type of tool (webhook, client, system)
type ConvaiToolType string

const (
	ConvaiToolTypeWebhook ConvaiToolType = "webhook"
	ConvaiToolTypeClient  ConvaiToolType = "client"
	ConvaiToolTypeSystem  ConvaiToolType = "system"
)

// Note: The following types are already defined in workflow.go and conversation_config.go:
// - ToolCallSoundType, ToolCallSoundBehavior, ToolExecutionMode (workflow.go)
// - WebhookMethod, WebhookContentType (workflow.go)
// - LiteralJsonSchemaProperty, LiteralJsonSchemaPropertyType (workflow.go)
// - QueryParamsSchema, ObjectJsonSchemaProperty, AuthConnectionLocator (workflow.go)
// - PromptBuiltInSystemToolType with SystemToolType* constants (conversation_config.go)

// ResourceAccessRole represents user role for resource access
type ResourceAccessRole string

const (
	ResourceAccessRoleAdmin     ResourceAccessRole = "admin"
	ResourceAccessRoleEditor    ResourceAccessRole = "editor"
	ResourceAccessRoleCommenter ResourceAccessRole = "commenter"
	ResourceAccessRoleViewer    ResourceAccessRole = "viewer"
)

// TransferType represents types of phone transfers
type TransferType string

const (
	TransferTypeBlind      TransferType = "blind"
	TransferTypeConference TransferType = "conference"
	TransferTypeSIPRefer   TransferType = "sip_refer"
)

// DynamicVariableAssignment for extracting values from tool responses
type DynamicVariableAssignment struct {
	Source          *string `json:"source,omitempty"`
	DynamicVariable string  `json:"dynamic_variable"`
	ValuePath       string  `json:"value_path"`
}

// DynamicVariablesConfig for dynamic variable configuration
type DynamicVariablesConfig struct {
	DynamicVariablePlaceholders map[string]any `json:"dynamic_variable_placeholders,omitempty"`
}

// ConvAISecretLocator references a secret by ID
type ConvAISecretLocator struct {
	SecretId string `json:"secret_id"`
}

// ConvAIDynamicVariable references a dynamic variable
type ConvAIDynamicVariable struct {
	VariableName string `json:"variable_name"`
}

// Note: AuthConnectionLocator, LiteralJsonSchemaProperty, LiteralJsonSchemaPropertyType,
// QueryParamsSchema, and ObjectJsonSchemaProperty are defined in workflow.go

// WebhookToolApiSchemaConfig for webhook API configuration
type WebhookToolApiSchemaConfig struct {
	URL               string                               `json:"url"`
	Method            *WebhookMethod                       `json:"method,omitempty"`
	RequestHeaders    map[string]any                       `json:"request_headers,omitempty"`
	PathParamsSchema  map[string]LiteralJsonSchemaProperty `json:"path_params_schema,omitempty"`
	QueryParamsSchema *QueryParamsSchema                   `json:"query_params_schema,omitempty"`
	RequestBodySchema *ObjectJsonSchemaProperty            `json:"request_body_schema,omitempty"`
	ContentType       *WebhookContentType                  `json:"content_type,omitempty"`
	AuthConnection    *AuthConnectionLocator               `json:"auth_connection,omitempty"`
}

// WebhookToolConfig represents a webhook tool configuration
type WebhookToolConfig struct {
	Type                  *ConvaiToolType             `json:"type,omitempty"`
	Name                  string                      `json:"name"`
	Description           string                      `json:"description"`
	ResponseTimeoutSecs   *int                        `json:"response_timeout_secs,omitempty"`
	DisableInterruptions  *bool                       `json:"disable_interruptions,omitempty"`
	ForcePreToolSpeech    *bool                       `json:"force_pre_tool_speech,omitempty"`
	Assignments           []DynamicVariableAssignment `json:"assignments,omitempty"`
	ToolCallSound         *ToolCallSoundType          `json:"tool_call_sound,omitempty"`
	ToolCallSoundBehavior *ToolCallSoundBehavior      `json:"tool_call_sound_behavior,omitempty"`
	DynamicVariables      *DynamicVariablesConfig     `json:"dynamic_variables,omitempty"`
	ExecutionMode         *ToolExecutionMode          `json:"execution_mode,omitempty"`
	APISchema             WebhookToolApiSchemaConfig  `json:"api_schema"`
}

// ClientToolConfig represents a client tool configuration
type ClientToolConfig struct {
	Type                  *ConvaiToolType             `json:"type,omitempty"`
	Name                  string                      `json:"name"`
	Description           string                      `json:"description"`
	ResponseTimeoutSecs   *int                        `json:"response_timeout_secs,omitempty"`
	DisableInterruptions  *bool                       `json:"disable_interruptions,omitempty"`
	ForcePreToolSpeech    *bool                       `json:"force_pre_tool_speech,omitempty"`
	Assignments           []DynamicVariableAssignment `json:"assignments,omitempty"`
	ToolCallSound         *ToolCallSoundType          `json:"tool_call_sound,omitempty"`
	ToolCallSoundBehavior *ToolCallSoundBehavior      `json:"tool_call_sound_behavior,omitempty"`
	Parameters            *ObjectJsonSchemaProperty   `json:"parameters,omitempty"`
	ExpectsResponse       *bool                       `json:"expects_response,omitempty"`
	DynamicVariables      *DynamicVariablesConfig     `json:"dynamic_variables,omitempty"`
	ExecutionMode         *ToolExecutionMode          `json:"execution_mode,omitempty"`
}

// ToolAgentTransfer represents a transfer to another agent
type ToolAgentTransfer struct {
	AgentId                            string  `json:"agent_id"`
	Condition                          string  `json:"condition"`
	DelayMs                            *int    `json:"delay_ms,omitempty"`
	TransferMessage                    *string `json:"transfer_message,omitempty"`
	EnableTransferredAgentFirstMessage *bool   `json:"enable_transferred_agent_first_message,omitempty"`
}

// ToolPhoneNumberTransferDestination represents a phone transfer destination for tools
type ToolPhoneNumberTransferDestination struct {
	Type        *string `json:"type,omitempty"`
	PhoneNumber *string `json:"phone_number,omitempty"`
	SIPUri      *string `json:"sip_uri,omitempty"`
}

// ToolPhoneNumberTransfer represents a phone number transfer configuration for tools
type ToolPhoneNumberTransfer struct {
	TransferDestination *ToolPhoneNumberTransferDestination `json:"transfer_destination,omitempty"`
	PhoneNumber         *string                             `json:"phone_number,omitempty"`
	Condition           string                              `json:"condition"`
	TransferType        *TransferType                       `json:"transfer_type,omitempty"`
}

// EndCallToolParams for end call tool
type EndCallToolParams struct {
	SystemToolType *PromptBuiltInSystemToolType `json:"system_tool_type,omitempty"`
}

// LanguageDetectionToolParams for language detection tool
type LanguageDetectionToolParams struct {
	SystemToolType *PromptBuiltInSystemToolType `json:"system_tool_type,omitempty"`
}

// TransferToAgentToolParams for transfer to agent tool
type TransferToAgentToolParams struct {
	SystemToolType *PromptBuiltInSystemToolType `json:"system_tool_type,omitempty"`
	Transfers      []ToolAgentTransfer          `json:"transfers"`
}

// TransferToNumberToolParams for transfer to number tool
type TransferToNumberToolParams struct {
	SystemToolType      *PromptBuiltInSystemToolType `json:"system_tool_type,omitempty"`
	Transfers           []ToolPhoneNumberTransfer    `json:"transfers"`
	EnableClientMessage *bool                        `json:"enable_client_message,omitempty"`
}

// SkipTurnToolParams for skip turn tool
type SkipTurnToolParams struct {
	SystemToolType *PromptBuiltInSystemToolType `json:"system_tool_type,omitempty"`
}

// PlayDTMFToolParams for play DTMF tool
type PlayDTMFToolParams struct {
	SystemToolType   *PromptBuiltInSystemToolType `json:"system_tool_type,omitempty"`
	UseOutOfBandDTMF *bool                        `json:"use_out_of_band_dtmf,omitempty"`
}

// VoicemailDetectionToolParams for voicemail detection tool
type VoicemailDetectionToolParams struct {
	SystemToolType   *PromptBuiltInSystemToolType `json:"system_tool_type,omitempty"`
	VoicemailMessage *string                      `json:"voicemail_message,omitempty"`
}

// SystemToolConfig represents a system tool configuration
type SystemToolConfig struct {
	Type                  *ConvaiToolType             `json:"type,omitempty"`
	Name                  string                      `json:"name"`
	Description           *string                     `json:"description,omitempty"`
	ResponseTimeoutSecs   *int                        `json:"response_timeout_secs,omitempty"`
	DisableInterruptions  *bool                       `json:"disable_interruptions,omitempty"`
	ForcePreToolSpeech    *bool                       `json:"force_pre_tool_speech,omitempty"`
	Assignments           []DynamicVariableAssignment `json:"assignments,omitempty"`
	ToolCallSound         *ToolCallSoundType          `json:"tool_call_sound,omitempty"`
	ToolCallSoundBehavior *ToolCallSoundBehavior      `json:"tool_call_sound_behavior,omitempty"`
	Params                any                         `json:"params"`
}

// ToolConfig is a union type for tool configurations
type ToolConfig struct {
	// Webhook tool fields
	Type                  *ConvaiToolType             `json:"type,omitempty"`
	Name                  string                      `json:"name"`
	Description           *string                     `json:"description,omitempty"`
	ResponseTimeoutSecs   *int                        `json:"response_timeout_secs,omitempty"`
	DisableInterruptions  *bool                       `json:"disable_interruptions,omitempty"`
	ForcePreToolSpeech    *bool                       `json:"force_pre_tool_speech,omitempty"`
	Assignments           []DynamicVariableAssignment `json:"assignments,omitempty"`
	ToolCallSound         *ToolCallSoundType          `json:"tool_call_sound,omitempty"`
	ToolCallSoundBehavior *ToolCallSoundBehavior      `json:"tool_call_sound_behavior,omitempty"`
	DynamicVariables      *DynamicVariablesConfig     `json:"dynamic_variables,omitempty"`
	ExecutionMode         *ToolExecutionMode          `json:"execution_mode,omitempty"`
	// Webhook specific
	APISchema *WebhookToolApiSchemaConfig `json:"api_schema,omitempty"`
	// Client specific
	Parameters      *ObjectJsonSchemaProperty `json:"parameters,omitempty"`
	ExpectsResponse *bool                     `json:"expects_response,omitempty"`
	// System specific
	Params any `json:"params,omitempty"`
}

// ResourceAccessInfo contains access information for a resource
type ResourceAccessInfo struct {
	IsCreator    bool               `json:"is_creator"`
	CreatorName  string             `json:"creator_name"`
	CreatorEmail string             `json:"creator_email"`
	Role         ResourceAccessRole `json:"role"`
}

// ToolUsageStats contains usage statistics for a tool
type ToolUsageStats struct {
	TotalCalls     int     `json:"total_calls"`
	AvgLatencySecs float64 `json:"avg_latency_secs"`
}

// ToolResponseModel represents a tool in responses
type ToolResponseModel struct {
	Id         string             `json:"id"`
	ToolConfig ToolConfig         `json:"tool_config"`
	AccessInfo ResourceAccessInfo `json:"access_info"`
	UsageStats ToolUsageStats     `json:"usage_stats"`
}

// ListToolsResp represents the response for listing tools
type ListToolsResp struct {
	Tools []ToolResponseModel `json:"tools"`
}

// ListTools gets all available tools in the workspace.
// https://elevenlabs.io/docs/api-reference/tools/list
func (c *Client) ListTools(ctx context.Context) (ListToolsResp, error) {
	body, err := c.get(ctx, "/convai/tools")
	if err != nil {
		return ListToolsResp{}, err
	}

	var resp ListToolsResp
	if err := c.parseResponse(body, &resp); err != nil {
		return ListToolsResp{}, err
	}

	return resp, nil
}

// GetToolReq represents the request for getting a tool
type GetToolReq struct {
	ToolId string `path:"tool_id"`
}

func NewGetToolReq(toolId string) *GetToolReq {
	return &GetToolReq{
		ToolId: toolId,
	}
}

// GetTool gets a tool that is available in the workspace.
// https://elevenlabs.io/docs/api-reference/tools/get
func (c *Client) GetTool(ctx context.Context, req *GetToolReq) (ToolResponseModel, error) {
	if req == nil {
		return ToolResponseModel{}, errors.New("request is nil")
	}

	body, err := c.get(ctx, "/convai/tools/"+req.ToolId)
	if err != nil {
		return ToolResponseModel{}, err
	}

	var resp ToolResponseModel
	if err := c.parseResponse(body, &resp); err != nil {
		return ToolResponseModel{}, err
	}

	return resp, nil
}

// CreateToolReq represents the request for creating a tool
type CreateToolReq struct {
	ToolConfig ToolConfig `json:"tool_config"`
}

func NewCreateToolReq(toolConfig ToolConfig) *CreateToolReq {
	return &CreateToolReq{
		ToolConfig: toolConfig,
	}
}

// CreateTool adds a new tool to the available tools in the workspace.
// https://elevenlabs.io/docs/api-reference/tools/create
func (c *Client) CreateTool(ctx context.Context, req *CreateToolReq) (ToolResponseModel, error) {
	if req == nil {
		return ToolResponseModel{}, errors.New("request is nil")
	}

	body, err := c.post(ctx, "/convai/tools", req)
	if err != nil {
		return ToolResponseModel{}, err
	}

	var resp ToolResponseModel
	if err := c.parseResponse(body, &resp); err != nil {
		return ToolResponseModel{}, err
	}

	return resp, nil
}

// UpdateToolReq represents the request for updating a tool
type UpdateToolReq struct {
	ToolId     string     `path:"tool_id"`
	ToolConfig ToolConfig `json:"tool_config"`
}

func NewUpdateToolReq(toolId string, toolConfig ToolConfig) *UpdateToolReq {
	return &UpdateToolReq{
		ToolId:     toolId,
		ToolConfig: toolConfig,
	}
}

// UpdateTool updates a tool that is available in the workspace.
// https://elevenlabs.io/docs/api-reference/tools/update
func (c *Client) UpdateTool(ctx context.Context, req *UpdateToolReq) (ToolResponseModel, error) {
	if req == nil {
		return ToolResponseModel{}, errors.New("request is nil")
	}

	body, err := c.patch(ctx, "/convai/tools/"+req.ToolId, req)
	if err != nil {
		return ToolResponseModel{}, err
	}

	var resp ToolResponseModel
	if err := c.parseResponse(body, &resp); err != nil {
		return ToolResponseModel{}, err
	}

	return resp, nil
}

// DeleteToolReq represents the request for deleting a tool
type DeleteToolReq struct {
	ToolId string `path:"tool_id"`
}

func NewDeleteToolReq(toolId string) *DeleteToolReq {
	return &DeleteToolReq{
		ToolId: toolId,
	}
}

// DeleteTool deletes a tool from the workspace.
// https://elevenlabs.io/docs/api-reference/tools/delete
func (c *Client) DeleteTool(ctx context.Context, req *DeleteToolReq) error {
	if req == nil {
		return errors.New("request is nil")
	}

	if err := c.delete(ctx, "/convai/tools/"+req.ToolId); err != nil {
		return err
	}

	return nil
}

// GetToolDependentAgentsReq represents the request for getting dependent agents
type GetToolDependentAgentsReq struct {
	ToolId   string  `path:"tool_id"`
	Cursor   *string `url:"cursor,omitempty"`
	PageSize int     `url:"page_size,omitempty"`
}

func NewGetToolDependentAgentsReq(toolId string) *GetToolDependentAgentsReq {
	return &GetToolDependentAgentsReq{
		ToolId:   toolId,
		PageSize: defaultPageSize,
	}
}

func (r GetToolDependentAgentsReq) QueryString() string {
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

// DependentAgentType represents the type of dependent agent
type DependentAgentType string

const (
	DependentAgentTypeAvailable DependentAgentType = "available"
	DependentAgentTypeUnknown   DependentAgentType = "unknown"
)

// DependentAgent represents an agent that depends on a tool
type DependentAgent struct {
	Type                  DependentAgentType  `json:"type"`
	Id                    *string             `json:"id,omitempty"`
	Name                  *string             `json:"name,omitempty"`
	CreatedAtUnixSecs     *int64              `json:"created_at_unix_secs,omitempty"`
	AccessLevel           *ResourceAccessRole `json:"access_level,omitempty"`
	ReferencedResourceIds []string            `json:"referenced_resource_ids,omitempty"`
}

// GetToolDependentAgentsResp represents the response for getting dependent agents
type GetToolDependentAgentsResp struct {
	Agents     []DependentAgent `json:"agents"`
	NextCursor *string          `json:"next_cursor,omitempty"`
	HasMore    bool             `json:"has_more"`
}

// GetToolDependentAgents gets a list of agents depending on this tool.
// https://elevenlabs.io/docs/api-reference/tools/get-dependent-agents
func (c *Client) GetToolDependentAgents(ctx context.Context, req *GetToolDependentAgentsReq) (GetToolDependentAgentsResp, error) {
	if req == nil {
		return GetToolDependentAgentsResp{}, errors.New("request is nil")
	}

	body, err := c.get(ctx, "/convai/tools/"+req.ToolId+"/dependent-agents"+req.QueryString())
	if err != nil {
		return GetToolDependentAgentsResp{}, err
	}

	var resp GetToolDependentAgentsResp
	if err := c.parseResponse(body, &resp); err != nil {
		return GetToolDependentAgentsResp{}, err
	}

	return resp, nil
}
