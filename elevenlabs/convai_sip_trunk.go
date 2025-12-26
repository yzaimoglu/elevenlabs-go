package elevenlabs

import (
	"context"
	"errors"
)

type ConvaiSIPTrunkAPI interface {
	SIPTrunkOutboundCall(ctx context.Context, req *SIPTrunkOutboundCallReq) (SIPTrunkOutboundCallResp, error)
}

// SIPTrunkOutboundCallReq represents the request for a SIP trunk outbound call
type SIPTrunkOutboundCallReq struct {
	AgentId                          string                            `json:"agent_id"`
	AgentPhoneNumberId               string                            `json:"agent_phone_number_id"`
	ToNumber                         string                            `json:"to_number"`
	ConversationInitiationClientData *ConversationInitiationClientData `json:"conversation_initiation_client_data,omitempty"`
}

func NewSIPTrunkOutboundCallReq(agentId, agentPhoneNumberId, toNumber string) *SIPTrunkOutboundCallReq {
	return &SIPTrunkOutboundCallReq{
		AgentId:            agentId,
		AgentPhoneNumberId: agentPhoneNumberId,
		ToNumber:           toNumber,
	}
}

// SIPTrunkOutboundCallResp represents the response from a SIP trunk outbound call
type SIPTrunkOutboundCallResp struct {
	Success        bool    `json:"success"`
	Message        string  `json:"message"`
	ConversationId *string `json:"conversation_id"`
	SIPCallId      *string `json:"sip_call_id"`
}

// SIPTrunkOutboundCall handles an outbound call via SIP trunk.
// https://elevenlabs.io/docs/api-reference/sip-trunk/outbound-call
func (c *Client) SIPTrunkOutboundCall(ctx context.Context, req *SIPTrunkOutboundCallReq) (SIPTrunkOutboundCallResp, error) {
	if req == nil {
		return SIPTrunkOutboundCallResp{}, errors.New("request is nil")
	}

	body, err := c.post(ctx, "/convai/sip-trunk/outbound-call", req)
	if err != nil {
		return SIPTrunkOutboundCallResp{}, err
	}

	var resp SIPTrunkOutboundCallResp
	if err := c.parseResponse(body, &resp); err != nil {
		return SIPTrunkOutboundCallResp{}, err
	}

	return resp, nil
}
