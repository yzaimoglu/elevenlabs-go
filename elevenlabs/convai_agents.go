package elevenlabs

import (
	"context"
	"errors"

	"github.com/google/go-querystring/query"
)

type ConvaiAgentsAPI interface {
	CreateAgent(ctx context.Context, req *CreateAgentReq) (CreateAgentResp, error)
	GetAgent(ctx context.Context, req *GetAgentReq) (GetAgentResp, error)
	ListAgents(ctx context.Context, req *ListAgentsReq) (ListAgentsResp, error)
	UpdateAgent(ctx context.Context, req *UpdateAgentReq) (UpdateAgentResp, error)
	DeleteAgent(ctx context.Context, req *DeleteAgentReq) error
	DuplicateAgent(ctx context.Context, req *DuplicateAgentReq) (DuplicateAgentResp, error)
	GetAgentLink(ctx context.Context, req *GetAgentLinkReq) (GetAgentLinkResp, error)
	SimulateConversation(ctx context.Context, req *SimulateConversationReq) (SimulateConversationResp, error)
	StreamSimulateConversation(ctx context.Context, req *SimulateConversationReq) (SimulateConversationResp, error)
	CalculateExpectedLLMUsage(ctx context.Context, req *CalculateExpectedLLMUsageReq) (CalculateExpectedLLMUsageResp, error)
}

type RespAgentMetadata struct {
	CreatedAtUnixSecs int64 `json:"created_at_unix_secs"`
}

type CreateAgentReq struct {
	ConversationConfig ConversationConfig `json:"conversation_config"`
	PlatformSettings   *PlatformSettings  `json:"platform_settings,omitempty"`
	Workflow           *AgentWorkflow     `json:"workflow,omitempty"`
	Name               *string            `json:"name,omitempty"`
	Tags               []string           `json:"tags,omitempty"`
}

func NewCreateAgentReq(conversationConfig ConversationConfig, platformSettings *PlatformSettings, workflow *AgentWorkflow, name *string, tags []string) *CreateAgentReq {
	return &CreateAgentReq{
		ConversationConfig: conversationConfig,
		PlatformSettings:   platformSettings,
		Workflow:           workflow,
		Name:               name,
		Tags:               tags,
	}
}

type CreateAgentResp struct {
	AgentId string `json:"agent_id"`
}

func (c *Client) CreateAgent(ctx context.Context, req *CreateAgentReq) (CreateAgentResp, error) {
	if req == nil {
		return CreateAgentResp{}, errors.New("request is nil")
	}

	body, err := c.post(ctx, "/convai/agents/create", req)
	if err != nil {
		return CreateAgentResp{}, err
	}

	var resp CreateAgentResp
	if err := c.parseResponse(body, &resp); err != nil {
		return CreateAgentResp{}, err
	}

	return resp, nil
}

type GetAgentReq struct {
	AgentId string `path:"agent_id"`
}

func NewGetAgentReq(agentId string) *GetAgentReq {
	return &GetAgentReq{
		AgentId: agentId,
	}
}

type GetAgentResp struct {
	AgentId            string             `json:"agent_id"`
	Name               string             `json:"name"`
	ConversationConfig ConversationConfig `json:"conversation_config,omitempty"`
	Metadata           RespAgentMetadata  `json:"metadata"`
	PlatformSettings   PlatformSettings   `json:"platform_settings,omitempty"`
	PhoneNumbers       []PhoneNumber      `json:"phone_numbers,omitempty"`
	Workflow           any                `json:"workflow,omitempty"`
	AccessInfo         *AccessInfo        `json:"access_info,omitempty"`
	Tags               []string           `json:"tags,omitempty"`
}

// GetAgent retrieves an agent by ID.
// https://elevenlabs.io/docs/api-reference/agents/get
func (c *Client) GetAgent(ctx context.Context, req *GetAgentReq) (GetAgentResp, error) {
	if req == nil {
		return GetAgentResp{}, errors.New("request is nil")
	}

	body, err := c.get(ctx, "/convai/agents/"+req.AgentId)
	if err != nil {
		return GetAgentResp{}, err
	}

	var resp GetAgentResp
	if err := c.parseResponse(body, &resp); err != nil {
		return GetAgentResp{}, err
	}

	return resp, nil
}

type ListAgentsReq struct {
	Cursor             *string              `url:"cursor,omitempty"`
	PageSize           int                  `url:"page_size,omitempty"`
	Search             *string              `url:"search,omitempty"`
	Archived           *bool                `url:"archived,omitempty"`
	ShowOnlyOwnedAgents *bool               `url:"show_only_owned_agents,omitempty"`
	SortDirection      *ListAgentsSortDirection `url:"sort_direction,omitempty"`
	SortBy             *ListAgentsSortBy    `url:"sort_by,omitempty"`
}

type ListAgentsSortDirection string

const (
	ListAgentsSortDirectionAsc  ListAgentsSortDirection = "asc"
	ListAgentsSortDirectionDesc ListAgentsSortDirection = "desc"
)

type ListAgentsSortBy string

const (
	ListAgentsSortByName      ListAgentsSortBy = "name"
	ListAgentsSortByCreatedAt ListAgentsSortBy = "created_at"
)

func NewListAgentsReq(cursor *string, pageSize *int, search *string) *ListAgentsReq {
	r := &ListAgentsReq{
		PageSize: defaultPageSize,
	}
	if cursor != nil {
		r.Cursor = cursor
	}
	if pageSize != nil {
		p := *pageSize

		if p < 1 || p > 100 {
			p = defaultPageSize
		}

		r.PageSize = p
	}

	if search != nil {
		r.Search = search
	}

	return r
}

func (r ListAgentsReq) QueryString() string {
	v, err := query.Values(r)
	if err != nil {
		return ""
	}

	return "?" + v.Encode()
}

type ListAgentsRespAgent struct {
	AgentId              string     `json:"agent_id"`
	Name                 string     `json:"name"`
	Tags                 []string   `json:"tags"`
	CreatedAtUnixSecs    int64      `json:"created_at_unix_secs"`
	AccessInfo           AccessInfo `json:"access_info"`
	LastCallTimeUnixSecs int64      `json:"last_call_time_unix_secs"`
	Archived             bool       `json:"archived"`
}

type ListAgentsResp struct {
	Agents     []ListAgentsRespAgent `json:"agents"`
	HasMore    bool                  `json:"has_more"`
	NextCursor *string               `json:"next_cursor,omitempty"`
}

// ListAgents lists all agents.
// https://elevenlabs.io/docs/api-reference/agents/list
func (c *Client) ListAgents(ctx context.Context, req *ListAgentsReq) (ListAgentsResp, error) {
	if req == nil {
		return ListAgentsResp{}, errors.New("request is nil")
	}

	body, err := c.get(ctx, "/convai/agents"+req.QueryString())
	if err != nil {
		return ListAgentsResp{}, err
	}

	var resp ListAgentsResp
	if err := c.parseResponse(body, &resp); err != nil {
		return ListAgentsResp{}, err
	}

	return resp, nil
}

type UpdateAgentReq struct {
	AgentId            string              `path:"agent_id"`
	ConversationConfig *ConversationConfig `json:"conversation_config,omitempty"`
	PlatformSettings   *PlatformSettings   `json:"platform_settings,omitempty"`
	Name               *string             `json:"name,omitempty"`
	Tags               []string            `json:"tags,omitempty"`
}

func NewUpdateAgentReq(agentId string, conversationConfig *ConversationConfig, platformSettings *PlatformSettings, name *string, tags []string) *UpdateAgentReq {
	return &UpdateAgentReq{
		AgentId:            agentId,
		ConversationConfig: conversationConfig,
		PlatformSettings:   platformSettings,
		Name:               name,
		Tags:               tags,
	}
}

type UpdateAgentResp struct {
	AgentId            string             `json:"agent_id"`
	Name               string             `json:"name"`
	ConversationConfig ConversationConfig `json:"conversation_config"`
	Metadata           RespAgentMetadata  `json:"metadata"`
	PlatformSettings   PlatformSettings   `json:"platform_settings,omitempty"`
	PhoneNumbers       []PhoneNumber      `json:"phone_numbers,omitempty"`
	Workflow           any                `json:"workflow,omitempty"`
	AccessInfo         *AccessInfo        `json:"access_info,omitempty"`
	Tags               []string           `json:"tags,omitempty"`
}

func (c *Client) UpdateAgent(ctx context.Context, req *UpdateAgentReq) (UpdateAgentResp, error) {
	if req == nil {
		return UpdateAgentResp{}, errors.New("request is nil")
	}

	body, err := c.patch(ctx, "/convai/agents/"+req.AgentId, req)
	if err != nil {
		return UpdateAgentResp{}, err
	}

	var resp UpdateAgentResp
	if err := c.parseResponse(body, &resp); err != nil {
		return UpdateAgentResp{}, err
	}

	return resp, nil
}

type DeleteAgentReq struct {
	AgentId string `path:"agent_id"`
}

func NewDeleteAgentReq(agentId string) *DeleteAgentReq {
	return &DeleteAgentReq{
		AgentId: agentId,
	}
}

// DeleteAgent deletes an agent.
// https://elevenlabs.io/docs/api-reference/agents/delete
func (c *Client) DeleteAgent(ctx context.Context, req *DeleteAgentReq) error {
	if req == nil {
		return errors.New("request is nil")
	}

	if err := c.delete(ctx, "/convai/agents/"+req.AgentId); err != nil {
		return err
	}

	return nil
}

type DuplicateAgentReq struct {
	AgentId string  `path:"agent_id"`
	Name    *string `json:"name,omitempty"`
}

func NewDuplicateAgentReq(agentId string) *DuplicateAgentReq {
	return &DuplicateAgentReq{
		AgentId: agentId,
	}
}

type DuplicateAgentResp struct {
	AgentId string `json:"agent_id"`
}

// DuplicateAgent creates a new agent by duplicating an existing one.
// https://elevenlabs.io/docs/api-reference/agents/duplicate
func (c *Client) DuplicateAgent(ctx context.Context, req *DuplicateAgentReq) (DuplicateAgentResp, error) {
	if req == nil {
		return DuplicateAgentResp{}, errors.New("request is nil")
	}

	body, err := c.post(ctx, "/convai/agents/"+req.AgentId+"/duplicate", req)
	if err != nil {
		return DuplicateAgentResp{}, err
	}

	var resp DuplicateAgentResp
	if err := c.parseResponse(body, &resp); err != nil {
		return DuplicateAgentResp{}, err
	}

	return resp, nil
}

type GetAgentLinkReq struct {
	AgentId string `path:"agent_id"`
}

func NewGetAgentLinkReq(agentId string) *GetAgentLinkReq {
	return &GetAgentLinkReq{
		AgentId: agentId,
	}
}

type GetAgentLinkResp struct {
	AgentId string                 `json:"agent_id"`
	Token   *GetAgentLinkRespToken `json:"token,omitempty"`
}

type GetAgentLinkRespToken struct {
	AgentId                string        `json:"agent_id"`
	ConversationToken      string        `json:"conversation_token"`
	ExpirationTimeUnixSecs *int64        `json:"expiration_time_unix_secs,omitempty"`
	ConversationId         *string       `json:"conversation_id,omitempty"`
	Purpose                *TokenPurpose `json:"purpose,omitempty"`
	TokenRequesterUserId   *string       `json:"token_requester_user_id,omitempty"`
}

// GetAgentLink gets the current link used to share the agent with others by agent id.
// https://elevenlabs.io/docs/api-reference/agents/get-link
func (c *Client) GetAgentLink(ctx context.Context, req *GetAgentLinkReq) (GetAgentLinkResp, error) {
	if req == nil {
		return GetAgentLinkResp{}, errors.New("request is nil")
	}

	body, err := c.get(ctx, "/convai/agents/"+req.AgentId+"/link")
	if err != nil {
		return GetAgentLinkResp{}, err
	}

	var resp GetAgentLinkResp
	if err := c.parseResponse(body, &resp); err != nil {
		return GetAgentLinkResp{}, err
	}

	return resp, nil
}

type SimulateConversationReq struct {
	AgentId                 string                               `path:"agent_id"`
	SimulationSpecification SimulationSpecification              `json:"simulation_specification"`
	ExtraEvaluationCriteria []SimulationExtraEvaluationCriterium `json:"extra_evaluation_criteria,omitempty"`
	NewTurnsLimit           *int                                 `json:"new_turns_limit,omitempty"`
}

func NewSimulateConversationReq(agentId string, simulationSpecification SimulationSpecification, extraEvaluationCriteria []SimulationExtraEvaluationCriterium, newTurnsLimit *int) *SimulateConversationReq {
	return &SimulateConversationReq{
		AgentId:                 agentId,
		SimulationSpecification: simulationSpecification,
		ExtraEvaluationCriteria: extraEvaluationCriteria,
		NewTurnsLimit:           newTurnsLimit,
	}
}

type SimulateConversationResp struct {
	SimulatedConversation []SimulatedConversation       `json:"simulated_conversation"`
	Analysis              SimulatedConversationAnalysis `json:"analysis"`
}

// SimulateConversation runs a conversation between the agent and a simulated user.
// https://elevenlabs.io/docs/api-reference/agents/simulate-conversation
func (c *Client) SimulateConversation(ctx context.Context, req *SimulateConversationReq) (SimulateConversationResp, error) {
	if req == nil {
		return SimulateConversationResp{}, errors.New("request is nil")
	}

	body, err := c.post(ctx, "/convai/agents/"+req.AgentId+"/simulate-conversation", req)
	if err != nil {
		return SimulateConversationResp{}, err
	}

	var resp SimulateConversationResp
	if err := c.parseResponse(body, &resp); err != nil {
		return SimulateConversationResp{}, err
	}

	return resp, nil
}

// StreamSimulateConversation runs a conversation between the agent and a simulated user and stream back the response.
// Response is streamed back as partial lists of messages that should be concatenated and once the conversation has complete a single final message with the conversation analysis will be sent.
// https://elevenlabs.io/docs/api-reference/agents/simulate-conversation-stream
func (c *Client) StreamSimulateConversation(ctx context.Context, req *SimulateConversationReq) (SimulateConversationResp, error) {
	if req == nil {
		return SimulateConversationResp{}, errors.New("request is nil")
	}

	body, err := c.post(ctx, "/convai/agents/"+req.AgentId+"/simulate-conversation/stream", req)
	if err != nil {
		return SimulateConversationResp{}, err
	}

	var resp SimulateConversationResp
	if err := c.parseResponse(body, &resp); err != nil {
		return SimulateConversationResp{}, err
	}

	return resp, nil
}

type CalculateExpectedLLMUsageReq struct {
	AgentId       string `path:"agent_id"`
	PromptLength  *int   `json:"prompt_length,omitempty"`
	NumberOfPages *int   `json:"number_of_pages,omitempty"`
	RAGEnabled    *bool  `json:"rag_enabled,omitempty"`
}

func NewCalculateExpectedLLMUsageReq(agentId string, promptLength *int, numberOfPages *int, ragEnabled *bool) *CalculateExpectedLLMUsageReq {
	return &CalculateExpectedLLMUsageReq{
		AgentId:       agentId,
		PromptLength:  promptLength,
		NumberOfPages: numberOfPages,
		RAGEnabled:    ragEnabled,
	}
}

type CalculateExpectedLLMUsageResp struct {
	LLMPrices []CalculateExpectedLLMUsageRespLLMPrice `json:"llm_prices"`
}

type CalculateExpectedLLMUsageRespLLMPrice struct {
	LLM            AgentPromptLLM `json:"llm"`
	PricePerMinute float32        `json:"price_per_minute"`
}

// CalculateExpectedLLMUsage calculates expected number of LLM tokens needed for the specified agent.
// https://elevenlabs.io/docs/api-reference/agents/calculate-expected-llm-usage
func (c *Client) CalculateExpectedLLMUsage(ctx context.Context, req *CalculateExpectedLLMUsageReq) (CalculateExpectedLLMUsageResp, error) {
	if req == nil {
		return CalculateExpectedLLMUsageResp{}, errors.New("request is nil")
	}

	body, err := c.post(ctx, "/convai/agents/"+req.AgentId+"/llm-usage/calculate", req)
	if err != nil {
		return CalculateExpectedLLMUsageResp{}, err
	}

	var resp CalculateExpectedLLMUsageResp
	if err := c.parseResponse(body, &resp); err != nil {
		return CalculateExpectedLLMUsageResp{}, err
	}

	return resp, nil
}
