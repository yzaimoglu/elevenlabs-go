package elevenlabs

import (
	"context"
	"errors"

	"github.com/google/go-querystring/query"
	"github.com/k0kubun/pp/v3"
)

type ConvaiAgentsAPI interface {
	GetAgent(ctx context.Context, req *GetAgentReq) (GetAgentResp, error)
	ListAgents(ctx context.Context, req *ListAgentsReq) (ListAgentsResp, error)
}

type GetAgentReq struct {
	AgentId string `path:"agent_id"`
}

func NewGetAgentReq(agentId string) *GetAgentReq {
	return &GetAgentReq{
		AgentId: agentId,
	}
}

type GetAgentRespMetadata struct {
	CreatedAtUnixSecs int64 `json:"created_at_unix_secs"`
}

type GetAgentResp struct {
	AgentId            string               `json:"agent_id"`
	Name               string               `json:"name"`
	ConversationConfig ConversationConfig   `json:"conversation_config,omitempty"`
	Metadata           GetAgentRespMetadata `json:"metadata"`
	PlatformSettings   PlatformSettings     `json:"platform_settings,omitempty"`
	PhoneNumbers       []PhoneNumber        `json:"phone_numbers,omitempty"`
	Workflow           any                  `json:"workflow,omitempty"`
	AccessInfo         *AccessInfo          `json:"access_info,omitempty"`
	Tags               []string             `json:"tags,omitempty"`
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

	var resp map[string]any
	if err := c.parseResponse(body, &resp); err != nil {
		return GetAgentResp{}, err
	}
	pp.Println(resp)

	var respRet GetAgentResp
	if err := c.parseResponse(body, &respRet); err != nil {
		return GetAgentResp{}, err
	}

	return respRet, nil
}

type ListAgentsReq struct {
	Cursor   *string `url:"cursor,omitempty"`
	PageSize int     `url:"page_size,omitempty"`
	Search   *string `url:"search,omitempty"`
}

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
