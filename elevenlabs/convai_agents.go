package elevenlabs

import (
	"context"
	"errors"

	"github.com/google/go-querystring/query"
)

type GetAgentResp struct {
	AgentId            string                       `json:"agent_id"`
	Name               string                       `json:"name"`
	ConversationConfig ConversationConfig           `json:"conversation_config,omitempty"`
	Metadata           GetAgentRespMetadata         `json:"metadata"`
	PlatformSettings   GetAgentRespPlatformSettings `json:"platform_settings,omitempty"`
	PhoneNumbers       GetAgentRespPhoneNumbers     `json:"phone_numbers,omitempty"`
	Workflow           any                          `json:"workflow,omitempty"`
	AccessInfo         *AccessInfo                  `json:"access_info,omitempty"`
	Tags               []string                     `json:"tags,omitempty"`
}

type GetAgentRespMetadata struct {
	CreatedAtUnixSecs int64 `json:"created_at_unix_secs"`
}

type AccessInfo struct {
	IsCreator    bool   `json:"is_creator"`
	CreatorName  string `json:"creator_name"`
	CreatorEmail string `json:"creator_email"`
	Role         Role   `json:"role"`
}

type Role string

const (
	RoleAdmin  = "admin"
	RoleEditor = "editor"
	RoleViewer = "viewer"
)

type ConvaiAgentsAPI interface {
	GetAgent(ctx context.Context, agentId string) (Agent, error)
	ListAgents(ctx context.Context, req *ListAgentsReq) (ListAgentsResp, error)
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
