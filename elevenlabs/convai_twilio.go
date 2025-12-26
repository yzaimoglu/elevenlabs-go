package elevenlabs

import (
	"context"
	"errors"
)

type ConvaiTwilioAPI interface {
	RegisterTwilioCall(ctx context.Context, req *RegisterTwilioCallReq) (map[string]any, error)
	TwilioOutboundCall(ctx context.Context, req *TwilioOutboundCallReq) (TwilioOutboundCallResp, error)
}

// TwilioCallDirection represents the direction of a Twilio call
type TwilioCallDirection string

const (
	TwilioCallDirectionInbound  TwilioCallDirection = "inbound"
	TwilioCallDirectionOutbound TwilioCallDirection = "outbound"
)

// RegisterTwilioCallReq represents the request to register a Twilio call
type RegisterTwilioCallReq struct {
	AgentId                          string                            `json:"agent_id"`
	FromNumber                       string                            `json:"from_number"`
	ToNumber                         string                            `json:"to_number"`
	Direction                        *TwilioCallDirection              `json:"direction,omitempty"`
	ConversationInitiationClientData *ConversationInitiationClientData `json:"conversation_initiation_client_data,omitempty"`
}

func NewRegisterTwilioCallReq(agentId, fromNumber, toNumber string) *RegisterTwilioCallReq {
	return &RegisterTwilioCallReq{
		AgentId:    agentId,
		FromNumber: fromNumber,
		ToNumber:   toNumber,
	}
}

// RegisterTwilioCall registers a Twilio call and returns TwiML to connect the call.
// https://elevenlabs.io/docs/api-reference/twilio/register-call
func (c *Client) RegisterTwilioCall(ctx context.Context, req *RegisterTwilioCallReq) (map[string]any, error) {
	if req == nil {
		return nil, errors.New("request is nil")
	}

	body, err := c.post(ctx, "/convai/twilio/register-call", req)
	if err != nil {
		return nil, err
	}

	var resp map[string]any
	if err := c.parseResponse(body, &resp); err != nil {
		return nil, err
	}

	return resp, nil
}

// TwilioOutboundCallReq represents the request for a Twilio outbound call
type TwilioOutboundCallReq struct {
	AgentId                          string                            `json:"agent_id"`
	AgentPhoneNumberId               string                            `json:"agent_phone_number_id"`
	ToNumber                         string                            `json:"to_number"`
	ConversationInitiationClientData *ConversationInitiationClientData `json:"conversation_initiation_client_data,omitempty"`
}

func NewTwilioOutboundCallReq(agentId, agentPhoneNumberId, toNumber string) *TwilioOutboundCallReq {
	return &TwilioOutboundCallReq{
		AgentId:            agentId,
		AgentPhoneNumberId: agentPhoneNumberId,
		ToNumber:           toNumber,
	}
}

// TwilioOutboundCallResp represents the response from a Twilio outbound call
type TwilioOutboundCallResp struct {
	Success        bool    `json:"success"`
	Message        string  `json:"message"`
	ConversationId *string `json:"conversation_id"`
	CallSid        *string `json:"callSid"`
}

// TwilioOutboundCall handles an outbound call via Twilio.
// https://elevenlabs.io/docs/api-reference/twilio/outbound-call
func (c *Client) TwilioOutboundCall(ctx context.Context, req *TwilioOutboundCallReq) (TwilioOutboundCallResp, error) {
	if req == nil {
		return TwilioOutboundCallResp{}, errors.New("request is nil")
	}

	body, err := c.post(ctx, "/convai/twilio/outbound-call", req)
	if err != nil {
		return TwilioOutboundCallResp{}, err
	}

	var resp TwilioOutboundCallResp
	if err := c.parseResponse(body, &resp); err != nil {
		return TwilioOutboundCallResp{}, err
	}

	return resp, nil
}
