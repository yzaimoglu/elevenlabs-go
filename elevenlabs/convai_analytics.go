package elevenlabs

import (
	"context"

	"github.com/google/go-querystring/query"
)

type ConvaiAnalyticsAPI interface {
	GetLiveCount(ctx context.Context, req *GetLiveCountReq) (GetLiveCountResp, error)
}

// GetLiveCountReq represents the request for getting live count
type GetLiveCountReq struct {
	AgentId *string `url:"agent_id,omitempty"`
}

func NewGetLiveCountReq() *GetLiveCountReq {
	return &GetLiveCountReq{}
}

func (r GetLiveCountReq) QueryString() string {
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

// GetLiveCountResp represents the response for getting live count
type GetLiveCountResp struct {
	Count int `json:"count"`
}

// GetLiveCount gets the live count of ongoing conversations.
// https://elevenlabs.io/docs/api-reference/analytics/get
func (c *Client) GetLiveCount(ctx context.Context, req *GetLiveCountReq) (GetLiveCountResp, error) {
	if req == nil {
		req = NewGetLiveCountReq()
	}

	body, err := c.get(ctx, "/convai/analytics/live-count"+req.QueryString())
	if err != nil {
		return GetLiveCountResp{}, err
	}

	var resp GetLiveCountResp
	if err := c.parseResponse(body, &resp); err != nil {
		return GetLiveCountResp{}, err
	}

	return resp, nil
}
