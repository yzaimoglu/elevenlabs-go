package elevenlabs

import (
	"context"
)

type ConvaiSettingsAPI interface {
	GetSettings(ctx context.Context) (ConvAISettings, error)
	UpdateSettings(ctx context.Context, req *UpdateSettingsReq) (ConvAISettings, error)
}

// WebhookEventType represents the type of webhook event
type WebhookEventType string

const (
	WebhookEventTypeTranscript           WebhookEventType = "transcript"
	WebhookEventTypeAudio                WebhookEventType = "audio"
	WebhookEventTypeCallInitiationFailure WebhookEventType = "call_initiation_failure"
)

// WebhookRequestHeader represents a webhook request header (can be string or secret locator)
type WebhookRequestHeader struct {
	Value         string               `json:"-"`
	SecretLocator *ConvAISecretLocator `json:"-"`
}

// ConversationInitiationClientDataWebhook represents a webhook configuration
type ConversationInitiationClientDataWebhook struct {
	URL            string                          `json:"url"`
	RequestHeaders map[string]WebhookRequestHeader `json:"request_headers"`
}

// ConvAIWebhooks represents webhook settings
type ConvAIWebhooks struct {
	PostCallWebhookId *string            `json:"post_call_webhook_id,omitempty"`
	Events            []WebhookEventType `json:"events,omitempty"`
	SendAudio         *bool              `json:"send_audio,omitempty"`
}

// ConvAISettings represents ConvAI settings
type ConvAISettings struct {
	ConversationInitiationClientDataWebhook *ConversationInitiationClientDataWebhook `json:"conversation_initiation_client_data_webhook,omitempty"`
	Webhooks                                *ConvAIWebhooks                          `json:"webhooks,omitempty"`
	CanUseMCPServers                        bool                                     `json:"can_use_mcp_servers"`
	RAGRetentionPeriodDays                  int                                      `json:"rag_retention_period_days"`
	DefaultLivekitStack                     LivekitStackType                         `json:"default_livekit_stack,omitempty"`
}

// GetSettings retrieves ConvAI settings for the workspace.
// https://elevenlabs.io/docs/api-reference/workspace/get
func (c *Client) GetSettings(ctx context.Context) (ConvAISettings, error) {
	body, err := c.get(ctx, "/convai/settings")
	if err != nil {
		return ConvAISettings{}, err
	}

	var resp ConvAISettings
	if err := c.parseResponse(body, &resp); err != nil {
		return ConvAISettings{}, err
	}

	return resp, nil
}

// UpdateSettingsReq represents the request for updating settings
type UpdateSettingsReq struct {
	ConversationInitiationClientDataWebhook *ConversationInitiationClientDataWebhook `json:"conversation_initiation_client_data_webhook,omitempty"`
	Webhooks                                *ConvAIWebhooks                          `json:"webhooks,omitempty"`
	CanUseMCPServers                        *bool                                    `json:"can_use_mcp_servers,omitempty"`
	RAGRetentionPeriodDays                  *int                                     `json:"rag_retention_period_days,omitempty"`
	DefaultLivekitStack                     *LivekitStackType                        `json:"default_livekit_stack,omitempty"`
}

func NewUpdateSettingsReq() *UpdateSettingsReq {
	return &UpdateSettingsReq{}
}

// UpdateSettings updates ConvAI settings for the workspace.
// https://elevenlabs.io/docs/api-reference/workspace/update
func (c *Client) UpdateSettings(ctx context.Context, req *UpdateSettingsReq) (ConvAISettings, error) {
	if req == nil {
		req = NewUpdateSettingsReq()
	}

	body, err := c.patch(ctx, "/convai/settings", req)
	if err != nil {
		return ConvAISettings{}, err
	}

	var resp ConvAISettings
	if err := c.parseResponse(body, &resp); err != nil {
		return ConvAISettings{}, err
	}

	return resp, nil
}
