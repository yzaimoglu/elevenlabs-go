package elevenlabs

import (
	"context"
	"errors"
)

type ConvaiWhatsAppAPI interface {
	WhatsAppOutboundCall(ctx context.Context, req *WhatsAppOutboundCallReq) (WhatsAppOutboundCallResp, error)
}

// WhatsAppOutboundCallReq represents the request for a WhatsApp outbound call
type WhatsAppOutboundCallReq struct {
	WhatsAppPhoneNumberId                         string                            `json:"whatsapp_phone_number_id"`
	WhatsAppUserId                                string                            `json:"whatsapp_user_id"`
	WhatsAppCallPermissionRequestTemplateName     string                            `json:"whatsapp_call_permission_request_template_name"`
	WhatsAppCallPermissionRequestTemplateLanguageCode string                            `json:"whatsapp_call_permission_request_template_language_code"`
	AgentId                                       string                            `json:"agent_id"`
	ConversationInitiationClientData              *ConversationInitiationClientData `json:"conversation_initiation_client_data,omitempty"`
}

func NewWhatsAppOutboundCallReq(whatsappPhoneNumberId, whatsappUserId, templateName, templateLanguageCode, agentId string) *WhatsAppOutboundCallReq {
	return &WhatsAppOutboundCallReq{
		WhatsAppPhoneNumberId:                         whatsappPhoneNumberId,
		WhatsAppUserId:                                whatsappUserId,
		WhatsAppCallPermissionRequestTemplateName:     templateName,
		WhatsAppCallPermissionRequestTemplateLanguageCode: templateLanguageCode,
		AgentId:                                       agentId,
	}
}

// WhatsAppOutboundCallResp represents the response from a WhatsApp outbound call
type WhatsAppOutboundCallResp struct {
	Success        bool    `json:"success"`
	Message        string  `json:"message"`
	ConversationId *string `json:"conversation_id"`
}

// WhatsAppOutboundCall makes an outbound call via WhatsApp.
// https://elevenlabs.io/docs/api-reference/whats-app/outbound-call
func (c *Client) WhatsAppOutboundCall(ctx context.Context, req *WhatsAppOutboundCallReq) (WhatsAppOutboundCallResp, error) {
	if req == nil {
		return WhatsAppOutboundCallResp{}, errors.New("request is nil")
	}

	body, err := c.post(ctx, "/convai/whatsapp/outbound-call", req)
	if err != nil {
		return WhatsAppOutboundCallResp{}, err
	}

	var resp WhatsAppOutboundCallResp
	if err := c.parseResponse(body, &resp); err != nil {
		return WhatsAppOutboundCallResp{}, err
	}

	return resp, nil
}
