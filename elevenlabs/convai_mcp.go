package elevenlabs

import (
	"context"
	"errors"
)

type ConvaiMCPAPI interface {
	ListMCPServers(ctx context.Context) ([]MCPServerResponse, error)
	GetMCPServer(ctx context.Context, req *GetMCPServerReq) (MCPServerResponse, error)
	CreateMCPServer(ctx context.Context, req *CreateMCPServerReq) (MCPServerResponse, error)
	UpdateMCPServer(ctx context.Context, req *UpdateMCPServerReq) (MCPServerResponse, error)
	DeleteMCPServer(ctx context.Context, req *DeleteMCPServerReq) error
	ListMCPServerTools(ctx context.Context, req *ListMCPServerToolsReq) (ListMCPServerToolsResp, error)
	// Tool Approvals
	CreateToolApproval(ctx context.Context, req *CreateToolApprovalReq) (MCPServerResponse, error)
	UpdateApprovalPolicy(ctx context.Context, req *UpdateApprovalPolicyReq) (MCPServerResponse, error)
	DeleteToolApproval(ctx context.Context, req *DeleteToolApprovalReq) (MCPServerResponse, error)
	// Tool Configuration
	CreateToolConfig(ctx context.Context, req *CreateToolConfigReq) (MCPServerResponse, error)
	GetToolConfig(ctx context.Context, req *GetToolConfigReq) (MCPToolConfigOverride, error)
	UpdateToolConfig(ctx context.Context, req *UpdateToolConfigReq) (MCPServerResponse, error)
	DeleteToolConfig(ctx context.Context, req *DeleteToolConfigReq) (MCPServerResponse, error)
}

// MCPApprovalPolicy represents the approval policy for MCP servers
type MCPApprovalPolicy string

const (
	MCPApprovalPolicyAutoApproveAll       MCPApprovalPolicy = "auto_approve_all"
	MCPApprovalPolicyRequireApprovalAll   MCPApprovalPolicy = "require_approval_all"
	MCPApprovalPolicyRequireApprovalPerTool MCPApprovalPolicy = "require_approval_per_tool"
)

// MCPToolApprovalPolicy represents the approval policy for individual MCP tools
type MCPToolApprovalPolicy string

const (
	MCPToolApprovalPolicyAutoApproved     MCPToolApprovalPolicy = "auto_approved"
	MCPToolApprovalPolicyRequiresApproval MCPToolApprovalPolicy = "requires_approval"
)

// MCPServerTransport represents the transport type for MCP servers
type MCPServerTransport string

const (
	MCPServerTransportSSE            MCPServerTransport = "SSE"
	MCPServerTransportStreamableHTTP MCPServerTransport = "STREAMABLE_HTTP"
)

// MCPToolApprovalHash represents a tool approval hash
type MCPToolApprovalHash struct {
	ToolName       string                `json:"tool_name"`
	ToolHash       string                `json:"tool_hash"`
	ApprovalPolicy MCPToolApprovalPolicy `json:"approval_policy,omitempty"`
}

// MCPToolConfigOverride represents per-tool configuration overrides
type MCPToolConfigOverride struct {
	ToolName              string                      `json:"tool_name"`
	ForcePreToolSpeech    *bool                       `json:"force_pre_tool_speech,omitempty"`
	DisableInterruptions  *bool                       `json:"disable_interruptions,omitempty"`
	ToolCallSound         *ToolCallSoundType          `json:"tool_call_sound,omitempty"`
	ToolCallSoundBehavior *ToolCallSoundBehavior      `json:"tool_call_sound_behavior,omitempty"`
	ExecutionMode         *ToolExecutionMode          `json:"execution_mode,omitempty"`
	Assignments           []DynamicVariableAssignment `json:"assignments,omitempty"`
}

// MCPServerConfigInput represents the input configuration for creating an MCP server
type MCPServerConfigInput struct {
	ApprovalPolicy       MCPApprovalPolicy        `json:"approval_policy,omitempty"`
	ToolApprovalHashes   []MCPToolApprovalHash    `json:"tool_approval_hashes,omitempty"`
	Transport            MCPServerTransport       `json:"transport,omitempty"`
	URL                  string                   `json:"url"`
	SecretToken          *ConvAISecretLocator     `json:"secret_token,omitempty"`
	RequestHeaders       map[string]any           `json:"request_headers,omitempty"`
	Name                 string                   `json:"name"`
	Description          string                   `json:"description,omitempty"`
	ForcePreToolSpeech   bool                     `json:"force_pre_tool_speech,omitempty"`
	DisableInterruptions bool                     `json:"disable_interruptions,omitempty"`
	ToolCallSound        *ToolCallSoundType       `json:"tool_call_sound,omitempty"`
	ToolCallSoundBehavior ToolCallSoundBehavior   `json:"tool_call_sound_behavior,omitempty"`
	ExecutionMode        ToolExecutionMode        `json:"execution_mode,omitempty"`
	ToolConfigOverrides  []MCPToolConfigOverride  `json:"tool_config_overrides,omitempty"`
	DisableCompression   bool                     `json:"disable_compression,omitempty"`
}

// MCPServerConfigOutput represents the output configuration for an MCP server
type MCPServerConfigOutput struct {
	ApprovalPolicy       MCPApprovalPolicy        `json:"approval_policy,omitempty"`
	ToolApprovalHashes   []MCPToolApprovalHash    `json:"tool_approval_hashes,omitempty"`
	Transport            MCPServerTransport       `json:"transport,omitempty"`
	URL                  any                      `json:"url"`
	SecretToken          any                      `json:"secret_token,omitempty"`
	RequestHeaders       map[string]any           `json:"request_headers,omitempty"`
	Name                 string                   `json:"name"`
	Description          string                   `json:"description"`
	ForcePreToolSpeech   bool                     `json:"force_pre_tool_speech"`
	DisableInterruptions bool                     `json:"disable_interruptions"`
	ToolCallSound        *ToolCallSoundType       `json:"tool_call_sound,omitempty"`
	ToolCallSoundBehavior ToolCallSoundBehavior   `json:"tool_call_sound_behavior,omitempty"`
	ExecutionMode        ToolExecutionMode        `json:"execution_mode,omitempty"`
	ToolConfigOverrides  []MCPToolConfigOverride  `json:"tool_config_overrides,omitempty"`
	DisableCompression   bool                     `json:"disable_compression"`
}

// DependentAgentAccessLevel represents the access level for a dependent agent
type DependentAgentAccessLevel string

const (
	DependentAgentAccessLevelAdmin     DependentAgentAccessLevel = "admin"
	DependentAgentAccessLevelEditor    DependentAgentAccessLevel = "editor"
	DependentAgentAccessLevelCommenter DependentAgentAccessLevel = "commenter"
	DependentAgentAccessLevelViewer    DependentAgentAccessLevel = "viewer"
)

// DependentAgentIdentifier represents a dependent agent
type DependentAgentIdentifier struct {
	ReferencedResourceIds []string                  `json:"referenced_resource_ids,omitempty"`
	Id                    string                    `json:"id,omitempty"`
	Name                  string                    `json:"name,omitempty"`
	Type                  string                    `json:"type,omitempty"`
	CreatedAtUnixSecs     int64                     `json:"created_at_unix_secs,omitempty"`
	AccessLevel           DependentAgentAccessLevel `json:"access_level,omitempty"`
}

// MCPServerMetadata represents metadata for an MCP server
type MCPServerMetadata struct {
	CreatedAt   int64   `json:"created_at"`
	OwnerUserId *string `json:"owner_user_id,omitempty"`
}

// MCPServerResponse represents the full MCP server response
type MCPServerResponse struct {
	Id              string                     `json:"id"`
	Config          MCPServerConfigOutput      `json:"config"`
	AccessInfo      *ResourceAccessInfo        `json:"access_info,omitempty"`
	DependentAgents []DependentAgentIdentifier `json:"dependent_agents,omitempty"`
	Metadata        MCPServerMetadata          `json:"metadata"`
}

// ListMCPServersResp represents the response from listing MCP servers
type ListMCPServersResp struct {
	MCPServers []MCPServerResponse `json:"mcp_servers"`
}

// ListMCPServers retrieves all MCP server configurations in the workspace.
// https://elevenlabs.io/docs/api-reference/mcp/list
func (c *Client) ListMCPServers(ctx context.Context) ([]MCPServerResponse, error) {
	body, err := c.get(ctx, "/convai/mcp-servers")
	if err != nil {
		return nil, err
	}

	var resp ListMCPServersResp
	if err := c.parseResponse(body, &resp); err != nil {
		return nil, err
	}

	return resp.MCPServers, nil
}

// GetMCPServerReq represents the request for getting an MCP server
type GetMCPServerReq struct {
	MCPServerId string `path:"mcp_server_id"`
}

func NewGetMCPServerReq(mcpServerId string) *GetMCPServerReq {
	return &GetMCPServerReq{
		MCPServerId: mcpServerId,
	}
}

// GetMCPServer retrieves a specific MCP server configuration.
// https://elevenlabs.io/docs/api-reference/mcp/get
func (c *Client) GetMCPServer(ctx context.Context, req *GetMCPServerReq) (MCPServerResponse, error) {
	if req == nil {
		return MCPServerResponse{}, errors.New("request is nil")
	}

	body, err := c.get(ctx, "/convai/mcp-servers/"+req.MCPServerId)
	if err != nil {
		return MCPServerResponse{}, err
	}

	var resp MCPServerResponse
	if err := c.parseResponse(body, &resp); err != nil {
		return MCPServerResponse{}, err
	}

	return resp, nil
}

// CreateMCPServerReq represents the request for creating an MCP server
type CreateMCPServerReq struct {
	Config MCPServerConfigInput `json:"config"`
}

func NewCreateMCPServerReq(url, name string) *CreateMCPServerReq {
	return &CreateMCPServerReq{
		Config: MCPServerConfigInput{
			URL:  url,
			Name: name,
		},
	}
}

// CreateMCPServer creates a new MCP server configuration.
// https://elevenlabs.io/docs/api-reference/mcp/create
func (c *Client) CreateMCPServer(ctx context.Context, req *CreateMCPServerReq) (MCPServerResponse, error) {
	if req == nil {
		return MCPServerResponse{}, errors.New("request is nil")
	}

	body, err := c.post(ctx, "/convai/mcp-servers", req)
	if err != nil {
		return MCPServerResponse{}, err
	}

	var resp MCPServerResponse
	if err := c.parseResponse(body, &resp); err != nil {
		return MCPServerResponse{}, err
	}

	return resp, nil
}

// UpdateMCPServerReq represents the request for updating an MCP server
type UpdateMCPServerReq struct {
	MCPServerId           string                 `path:"mcp_server_id"`
	ApprovalPolicy        *MCPApprovalPolicy     `json:"approval_policy,omitempty"`
	ForcePreToolSpeech    *bool                  `json:"force_pre_tool_speech,omitempty"`
	DisableInterruptions  *bool                  `json:"disable_interruptions,omitempty"`
	ToolCallSound         *ToolCallSoundType     `json:"tool_call_sound,omitempty"`
	ToolCallSoundBehavior *ToolCallSoundBehavior `json:"tool_call_sound_behavior,omitempty"`
	ExecutionMode         *ToolExecutionMode     `json:"execution_mode,omitempty"`
	RequestHeaders        map[string]any         `json:"request_headers,omitempty"`
	DisableCompression    *bool                  `json:"disable_compression,omitempty"`
}

func NewUpdateMCPServerReq(mcpServerId string) *UpdateMCPServerReq {
	return &UpdateMCPServerReq{
		MCPServerId: mcpServerId,
	}
}

// UpdateMCPServer updates an MCP server configuration.
// https://elevenlabs.io/docs/api-reference/mcp/update
func (c *Client) UpdateMCPServer(ctx context.Context, req *UpdateMCPServerReq) (MCPServerResponse, error) {
	if req == nil {
		return MCPServerResponse{}, errors.New("request is nil")
	}

	body, err := c.patch(ctx, "/convai/mcp-servers/"+req.MCPServerId, req)
	if err != nil {
		return MCPServerResponse{}, err
	}

	var resp MCPServerResponse
	if err := c.parseResponse(body, &resp); err != nil {
		return MCPServerResponse{}, err
	}

	return resp, nil
}

// DeleteMCPServerReq represents the request for deleting an MCP server
type DeleteMCPServerReq struct {
	MCPServerId string `path:"mcp_server_id"`
}

func NewDeleteMCPServerReq(mcpServerId string) *DeleteMCPServerReq {
	return &DeleteMCPServerReq{
		MCPServerId: mcpServerId,
	}
}

// DeleteMCPServer deletes an MCP server configuration.
// https://elevenlabs.io/docs/api-reference/mcp/delete
func (c *Client) DeleteMCPServer(ctx context.Context, req *DeleteMCPServerReq) error {
	if req == nil {
		return errors.New("request is nil")
	}

	return c.delete(ctx, "/convai/mcp-servers/"+req.MCPServerId)
}

// ListMCPServerToolsReq represents the request for listing MCP server tools
type ListMCPServerToolsReq struct {
	MCPServerId string `path:"mcp_server_id"`
}

func NewListMCPServerToolsReq(mcpServerId string) *ListMCPServerToolsReq {
	return &ListMCPServerToolsReq{
		MCPServerId: mcpServerId,
	}
}

// MCPToolInputSchema represents the input schema for an MCP tool
type MCPToolInputSchema map[string]any

// MCPToolOutputSchema represents the output schema for an MCP tool
type MCPToolOutputSchema map[string]any

// MCPToolAnnotations represents annotations for an MCP tool
type MCPToolAnnotations struct {
	Title           *string `json:"title,omitempty"`
	ReadOnlyHint    *bool   `json:"readOnlyHint,omitempty"`
	DestructiveHint *bool   `json:"destructiveHint,omitempty"`
	IdempotentHint  *bool   `json:"idempotentHint,omitempty"`
	OpenWorldHint   *bool   `json:"openWorldHint,omitempty"`
}

// MCPTool represents an MCP tool
type MCPTool struct {
	Name         string              `json:"name"`
	Title        *string             `json:"title,omitempty"`
	Description  *string             `json:"description,omitempty"`
	InputSchema  MCPToolInputSchema  `json:"inputSchema"`
	OutputSchema *MCPToolOutputSchema `json:"outputSchema,omitempty"`
	Annotations  *MCPToolAnnotations `json:"annotations,omitempty"`
	Meta         map[string]any      `json:"_meta,omitempty"`
}

// ListMCPServerToolsResp represents the response from listing MCP server tools
type ListMCPServerToolsResp struct {
	Success      bool      `json:"success"`
	Tools        []MCPTool `json:"tools"`
	ErrorMessage *string   `json:"error_message,omitempty"`
}

// ListMCPServerTools retrieves all tools available for an MCP server.
// https://elevenlabs.io/docs/api-reference/mcp/list-tools
func (c *Client) ListMCPServerTools(ctx context.Context, req *ListMCPServerToolsReq) (ListMCPServerToolsResp, error) {
	if req == nil {
		return ListMCPServerToolsResp{}, errors.New("request is nil")
	}

	body, err := c.get(ctx, "/convai/mcp-servers/"+req.MCPServerId+"/tools")
	if err != nil {
		return ListMCPServerToolsResp{}, err
	}

	var resp ListMCPServerToolsResp
	if err := c.parseResponse(body, &resp); err != nil {
		return ListMCPServerToolsResp{}, err
	}

	return resp, nil
}

// CreateToolApprovalReq represents the request for creating a tool approval
type CreateToolApprovalReq struct {
	MCPServerId     string                `path:"mcp_server_id"`
	ToolName        string                `json:"tool_name"`
	ToolDescription string                `json:"tool_description"`
	InputSchema     map[string]any        `json:"input_schema,omitempty"`
	ApprovalPolicy  MCPToolApprovalPolicy `json:"approval_policy,omitempty"`
}

func NewCreateToolApprovalReq(mcpServerId, toolName, toolDescription string) *CreateToolApprovalReq {
	return &CreateToolApprovalReq{
		MCPServerId:     mcpServerId,
		ToolName:        toolName,
		ToolDescription: toolDescription,
	}
}

// CreateToolApproval adds approval for a specific MCP tool when using per-tool approval mode.
// https://elevenlabs.io/docs/api-reference/mcp/approval-policies/create
func (c *Client) CreateToolApproval(ctx context.Context, req *CreateToolApprovalReq) (MCPServerResponse, error) {
	if req == nil {
		return MCPServerResponse{}, errors.New("request is nil")
	}

	body, err := c.post(ctx, "/convai/mcp-servers/"+req.MCPServerId+"/tool-approvals", req)
	if err != nil {
		return MCPServerResponse{}, err
	}

	var resp MCPServerResponse
	if err := c.parseResponse(body, &resp); err != nil {
		return MCPServerResponse{}, err
	}

	return resp, nil
}

// UpdateApprovalPolicyReq represents the request for updating approval policy
type UpdateApprovalPolicyReq struct {
	MCPServerId    string            `path:"mcp_server_id"`
	ApprovalPolicy MCPApprovalPolicy `json:"approval_policy"`
}

func NewUpdateApprovalPolicyReq(mcpServerId string, approvalPolicy MCPApprovalPolicy) *UpdateApprovalPolicyReq {
	return &UpdateApprovalPolicyReq{
		MCPServerId:    mcpServerId,
		ApprovalPolicy: approvalPolicy,
	}
}

// UpdateApprovalPolicy updates the approval policy configuration for an MCP server.
// DEPRECATED: Use UpdateMCPServer endpoint instead.
// https://elevenlabs.io/docs/api-reference/mcp/approval-policies/update
func (c *Client) UpdateApprovalPolicy(ctx context.Context, req *UpdateApprovalPolicyReq) (MCPServerResponse, error) {
	if req == nil {
		return MCPServerResponse{}, errors.New("request is nil")
	}

	body, err := c.patch(ctx, "/convai/mcp-servers/"+req.MCPServerId+"/approval-policy", req)
	if err != nil {
		return MCPServerResponse{}, err
	}

	var resp MCPServerResponse
	if err := c.parseResponse(body, &resp); err != nil {
		return MCPServerResponse{}, err
	}

	return resp, nil
}

// DeleteToolApprovalReq represents the request for deleting a tool approval
type DeleteToolApprovalReq struct {
	MCPServerId string `path:"mcp_server_id"`
	ToolName    string `path:"tool_name"`
}

func NewDeleteToolApprovalReq(mcpServerId, toolName string) *DeleteToolApprovalReq {
	return &DeleteToolApprovalReq{
		MCPServerId: mcpServerId,
		ToolName:    toolName,
	}
}

// DeleteToolApproval removes approval for a specific MCP tool when using per-tool approval mode.
// https://elevenlabs.io/docs/api-reference/mcp/approval-policies/delete
func (c *Client) DeleteToolApproval(ctx context.Context, req *DeleteToolApprovalReq) (MCPServerResponse, error) {
	if req == nil {
		return MCPServerResponse{}, errors.New("request is nil")
	}

	body, err := c.deleteWithResponse(ctx, "/convai/mcp-servers/"+req.MCPServerId+"/tool-approvals/"+req.ToolName)
	if err != nil {
		return MCPServerResponse{}, err
	}

	var resp MCPServerResponse
	if err := c.parseResponse(body, &resp); err != nil {
		return MCPServerResponse{}, err
	}

	return resp, nil
}

// CreateToolConfigReq represents the request for creating a tool configuration
type CreateToolConfigReq struct {
	MCPServerId           string                      `path:"mcp_server_id"`
	ToolName              string                      `json:"tool_name"`
	ForcePreToolSpeech    *bool                       `json:"force_pre_tool_speech,omitempty"`
	DisableInterruptions  *bool                       `json:"disable_interruptions,omitempty"`
	ToolCallSound         *ToolCallSoundType          `json:"tool_call_sound,omitempty"`
	ToolCallSoundBehavior *ToolCallSoundBehavior      `json:"tool_call_sound_behavior,omitempty"`
	ExecutionMode         *ToolExecutionMode          `json:"execution_mode,omitempty"`
	Assignments           []DynamicVariableAssignment `json:"assignments,omitempty"`
}

func NewCreateToolConfigReq(mcpServerId, toolName string) *CreateToolConfigReq {
	return &CreateToolConfigReq{
		MCPServerId: mcpServerId,
		ToolName:    toolName,
	}
}

// CreateToolConfig creates configuration overrides for a specific MCP tool.
// https://elevenlabs.io/docs/api-reference/mcp/tool-configuration/create
func (c *Client) CreateToolConfig(ctx context.Context, req *CreateToolConfigReq) (MCPServerResponse, error) {
	if req == nil {
		return MCPServerResponse{}, errors.New("request is nil")
	}

	body, err := c.post(ctx, "/convai/mcp-servers/"+req.MCPServerId+"/tool-configs", req)
	if err != nil {
		return MCPServerResponse{}, err
	}

	var resp MCPServerResponse
	if err := c.parseResponse(body, &resp); err != nil {
		return MCPServerResponse{}, err
	}

	return resp, nil
}

// GetToolConfigReq represents the request for getting a tool configuration
type GetToolConfigReq struct {
	MCPServerId string `path:"mcp_server_id"`
	ToolName    string `path:"tool_name"`
}

func NewGetToolConfigReq(mcpServerId, toolName string) *GetToolConfigReq {
	return &GetToolConfigReq{
		MCPServerId: mcpServerId,
		ToolName:    toolName,
	}
}

// GetToolConfig retrieves configuration overrides for a specific MCP tool.
// https://elevenlabs.io/docs/api-reference/mcp/tool-configuration/get
func (c *Client) GetToolConfig(ctx context.Context, req *GetToolConfigReq) (MCPToolConfigOverride, error) {
	if req == nil {
		return MCPToolConfigOverride{}, errors.New("request is nil")
	}

	body, err := c.get(ctx, "/convai/mcp-servers/"+req.MCPServerId+"/tool-configs/"+req.ToolName)
	if err != nil {
		return MCPToolConfigOverride{}, err
	}

	var resp MCPToolConfigOverride
	if err := c.parseResponse(body, &resp); err != nil {
		return MCPToolConfigOverride{}, err
	}

	return resp, nil
}

// UpdateToolConfigReq represents the request for updating a tool configuration
type UpdateToolConfigReq struct {
	MCPServerId           string                      `path:"mcp_server_id"`
	ToolName              string                      `path:"tool_name"`
	ForcePreToolSpeech    *bool                       `json:"force_pre_tool_speech,omitempty"`
	DisableInterruptions  *bool                       `json:"disable_interruptions,omitempty"`
	ToolCallSound         *ToolCallSoundType          `json:"tool_call_sound,omitempty"`
	ToolCallSoundBehavior *ToolCallSoundBehavior      `json:"tool_call_sound_behavior,omitempty"`
	ExecutionMode         *ToolExecutionMode          `json:"execution_mode,omitempty"`
	Assignments           []DynamicVariableAssignment `json:"assignments,omitempty"`
}

func NewUpdateToolConfigReq(mcpServerId, toolName string) *UpdateToolConfigReq {
	return &UpdateToolConfigReq{
		MCPServerId: mcpServerId,
		ToolName:    toolName,
	}
}

// UpdateToolConfig updates configuration overrides for a specific MCP tool.
// https://elevenlabs.io/docs/api-reference/mcp/tool-configuration/update
func (c *Client) UpdateToolConfig(ctx context.Context, req *UpdateToolConfigReq) (MCPServerResponse, error) {
	if req == nil {
		return MCPServerResponse{}, errors.New("request is nil")
	}

	body, err := c.patch(ctx, "/convai/mcp-servers/"+req.MCPServerId+"/tool-configs/"+req.ToolName, req)
	if err != nil {
		return MCPServerResponse{}, err
	}

	var resp MCPServerResponse
	if err := c.parseResponse(body, &resp); err != nil {
		return MCPServerResponse{}, err
	}

	return resp, nil
}

// DeleteToolConfigReq represents the request for deleting a tool configuration
type DeleteToolConfigReq struct {
	MCPServerId string `path:"mcp_server_id"`
	ToolName    string `path:"tool_name"`
}

func NewDeleteToolConfigReq(mcpServerId, toolName string) *DeleteToolConfigReq {
	return &DeleteToolConfigReq{
		MCPServerId: mcpServerId,
		ToolName:    toolName,
	}
}

// DeleteToolConfig deletes configuration overrides for a specific MCP tool.
// https://elevenlabs.io/docs/api-reference/mcp/tool-configuration/delete
func (c *Client) DeleteToolConfig(ctx context.Context, req *DeleteToolConfigReq) (MCPServerResponse, error) {
	if req == nil {
		return MCPServerResponse{}, errors.New("request is nil")
	}

	body, err := c.deleteWithResponse(ctx, "/convai/mcp-servers/"+req.MCPServerId+"/tool-configs/"+req.ToolName)
	if err != nil {
		return MCPServerResponse{}, err
	}

	var resp MCPServerResponse
	if err := c.parseResponse(body, &resp); err != nil {
		return MCPServerResponse{}, err
	}

	return resp, nil
}
