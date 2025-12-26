package elevenlabs

import (
	"context"
	"errors"

	"github.com/google/go-querystring/query"
)

type ConvaiBatchCallingAPI interface {
	SubmitBatchCall(ctx context.Context, req *SubmitBatchCallReq) (BatchCallResponse, error)
	GetBatchCall(ctx context.Context, req *GetBatchCallReq) (BatchCallDetailedResponse, error)
	ListBatchCalls(ctx context.Context, req *ListBatchCallsReq) (ListBatchCallsResp, error)
	RetryBatchCall(ctx context.Context, req *RetryBatchCallReq) (BatchCallResponse, error)
	CancelBatchCall(ctx context.Context, req *CancelBatchCallReq) (BatchCallResponse, error)
}

// TelephonyProvider represents the telephony provider type
type TelephonyProvider string

const (
	TelephonyProviderTwilio   TelephonyProvider = "twilio"
	TelephonyProviderSIPTrunk TelephonyProvider = "sip_trunk"
)

// BatchCallStatus represents the status of a batch call
type BatchCallStatus string

const (
	BatchCallStatusPending    BatchCallStatus = "pending"
	BatchCallStatusInProgress BatchCallStatus = "in_progress"
	BatchCallStatusCompleted  BatchCallStatus = "completed"
	BatchCallStatusFailed     BatchCallStatus = "failed"
	BatchCallStatusCancelled  BatchCallStatus = "cancelled"
)

// BatchCallRecipientStatus represents the status of a batch call recipient
type BatchCallRecipientStatus string

const (
	BatchCallRecipientStatusPending    BatchCallRecipientStatus = "pending"
	BatchCallRecipientStatusInitiated  BatchCallRecipientStatus = "initiated"
	BatchCallRecipientStatusInProgress BatchCallRecipientStatus = "in_progress"
	BatchCallRecipientStatusCompleted  BatchCallRecipientStatus = "completed"
	BatchCallRecipientStatusFailed     BatchCallRecipientStatus = "failed"
	BatchCallRecipientStatusCancelled  BatchCallRecipientStatus = "cancelled"
	BatchCallRecipientStatusVoicemail  BatchCallRecipientStatus = "voicemail"
)

// OutboundCallRecipient represents a recipient for outbound calls
type OutboundCallRecipient struct {
	Id                             *string                           `json:"id,omitempty"`
	PhoneNumber                    *string                           `json:"phone_number,omitempty"`
	WhatsAppUserId                 *string                           `json:"whatsapp_user_id,omitempty"`
	ConversationInitiationClientData *ConversationInitiationClientData `json:"conversation_initiation_client_data,omitempty"`
}

// BatchCallWhatsAppParams represents WhatsApp parameters for batch calls
type BatchCallWhatsAppParams struct {
	WhatsAppPhoneNumberId                         *string `json:"whatsapp_phone_number_id,omitempty"`
	WhatsAppCallPermissionRequestTemplateName     string  `json:"whatsapp_call_permission_request_template_name"`
	WhatsAppCallPermissionRequestTemplateLanguageCode string  `json:"whatsapp_call_permission_request_template_language_code"`
}

// SubmitBatchCallReq represents the request to submit a batch call
type SubmitBatchCallReq struct {
	CallName           string                   `json:"call_name"`
	AgentId            string                   `json:"agent_id"`
	Recipients         []OutboundCallRecipient  `json:"recipients"`
	ScheduledTimeUnix  *int64                   `json:"scheduled_time_unix,omitempty"`
	AgentPhoneNumberId *string                  `json:"agent_phone_number_id,omitempty"`
	WhatsAppParams     *BatchCallWhatsAppParams `json:"whatsapp_params,omitempty"`
}

func NewSubmitBatchCallReq(callName, agentId string, recipients []OutboundCallRecipient) *SubmitBatchCallReq {
	return &SubmitBatchCallReq{
		CallName:   callName,
		AgentId:    agentId,
		Recipients: recipients,
	}
}

// BatchCallResponse represents the response from batch call operations
type BatchCallResponse struct {
	Id                   string                   `json:"id"`
	PhoneNumberId        *string                  `json:"phone_number_id,omitempty"`
	PhoneProvider        *TelephonyProvider       `json:"phone_provider,omitempty"`
	WhatsAppParams       *BatchCallWhatsAppParams `json:"whatsapp_params,omitempty"`
	Name                 string                   `json:"name"`
	AgentId              string                   `json:"agent_id"`
	CreatedAtUnix        int64                    `json:"created_at_unix"`
	ScheduledTimeUnix    int64                    `json:"scheduled_time_unix"`
	TotalCallsDispatched int                      `json:"total_calls_dispatched"`
	TotalCallsScheduled  int                      `json:"total_calls_scheduled"`
	LastUpdatedAtUnix    int64                    `json:"last_updated_at_unix"`
	Status               BatchCallStatus          `json:"status"`
	RetryCount           int                      `json:"retry_count"`
	AgentName            string                   `json:"agent_name"`
}

// SubmitBatchCall submits a batch call request to schedule calls for multiple recipients.
// https://elevenlabs.io/docs/api-reference/batch-calling/create
func (c *Client) SubmitBatchCall(ctx context.Context, req *SubmitBatchCallReq) (BatchCallResponse, error) {
	if req == nil {
		return BatchCallResponse{}, errors.New("request is nil")
	}

	body, err := c.post(ctx, "/convai/batch-calling/submit", req)
	if err != nil {
		return BatchCallResponse{}, err
	}

	var resp BatchCallResponse
	if err := c.parseResponse(body, &resp); err != nil {
		return BatchCallResponse{}, err
	}

	return resp, nil
}

// GetBatchCallReq represents the request to get a batch call
type GetBatchCallReq struct {
	BatchId string `path:"batch_id"`
}

func NewGetBatchCallReq(batchId string) *GetBatchCallReq {
	return &GetBatchCallReq{
		BatchId: batchId,
	}
}

// OutboundCallRecipientResponse represents a recipient in batch call responses
type OutboundCallRecipientResponse struct {
	Id                               string                            `json:"id"`
	PhoneNumber                      *string                           `json:"phone_number,omitempty"`
	WhatsAppUserId                   *string                           `json:"whatsapp_user_id,omitempty"`
	Status                           BatchCallRecipientStatus          `json:"status"`
	CreatedAtUnix                    int64                             `json:"created_at_unix"`
	UpdatedAtUnix                    int64                             `json:"updated_at_unix"`
	ConversationId                   *string                           `json:"conversation_id"`
	ConversationInitiationClientData *ConversationInitiationClientData `json:"conversation_initiation_client_data,omitempty"`
}

// BatchCallDetailedResponse represents detailed batch call response
type BatchCallDetailedResponse struct {
	Id                   string                          `json:"id"`
	PhoneNumberId        *string                         `json:"phone_number_id,omitempty"`
	PhoneProvider        *TelephonyProvider              `json:"phone_provider,omitempty"`
	WhatsAppParams       *BatchCallWhatsAppParams        `json:"whatsapp_params,omitempty"`
	Name                 string                          `json:"name"`
	AgentId              string                          `json:"agent_id"`
	CreatedAtUnix        int64                           `json:"created_at_unix"`
	ScheduledTimeUnix    int64                           `json:"scheduled_time_unix"`
	TotalCallsDispatched int                             `json:"total_calls_dispatched"`
	TotalCallsScheduled  int                             `json:"total_calls_scheduled"`
	LastUpdatedAtUnix    int64                           `json:"last_updated_at_unix"`
	Status               BatchCallStatus                 `json:"status"`
	RetryCount           int                             `json:"retry_count"`
	AgentName            string                          `json:"agent_name"`
	Recipients           []OutboundCallRecipientResponse `json:"recipients"`
}

// GetBatchCall gets detailed information about a batch call including all recipients.
// https://elevenlabs.io/docs/api-reference/batch-calling/get
func (c *Client) GetBatchCall(ctx context.Context, req *GetBatchCallReq) (BatchCallDetailedResponse, error) {
	if req == nil {
		return BatchCallDetailedResponse{}, errors.New("request is nil")
	}

	body, err := c.get(ctx, "/convai/batch-calling/"+req.BatchId)
	if err != nil {
		return BatchCallDetailedResponse{}, err
	}

	var resp BatchCallDetailedResponse
	if err := c.parseResponse(body, &resp); err != nil {
		return BatchCallDetailedResponse{}, err
	}

	return resp, nil
}

// ListBatchCallsReq represents the request to list batch calls
type ListBatchCallsReq struct {
	Limit   int     `url:"limit,omitempty"`
	LastDoc *string `url:"last_doc,omitempty"`
}

func NewListBatchCallsReq() *ListBatchCallsReq {
	return &ListBatchCallsReq{
		Limit: 100,
	}
}

func (r ListBatchCallsReq) QueryString() string {
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

// ListBatchCallsResp represents the response from listing batch calls
type ListBatchCallsResp struct {
	BatchCalls []BatchCallResponse `json:"batch_calls"`
	NextDoc    *string             `json:"next_doc,omitempty"`
	HasMore    bool                `json:"has_more"`
}

// ListBatchCalls gets all batch calls for the current workspace.
// https://elevenlabs.io/docs/api-reference/batch-calling/list
func (c *Client) ListBatchCalls(ctx context.Context, req *ListBatchCallsReq) (ListBatchCallsResp, error) {
	if req == nil {
		req = NewListBatchCallsReq()
	}

	body, err := c.get(ctx, "/convai/batch-calling/workspace"+req.QueryString())
	if err != nil {
		return ListBatchCallsResp{}, err
	}

	var resp ListBatchCallsResp
	if err := c.parseResponse(body, &resp); err != nil {
		return ListBatchCallsResp{}, err
	}

	return resp, nil
}

// RetryBatchCallReq represents the request to retry a batch call
type RetryBatchCallReq struct {
	BatchId string `path:"batch_id"`
}

func NewRetryBatchCallReq(batchId string) *RetryBatchCallReq {
	return &RetryBatchCallReq{
		BatchId: batchId,
	}
}

// RetryBatchCall retries a batch call, calling failed and no-response recipients again.
// https://elevenlabs.io/docs/api-reference/batch-calling/retry
func (c *Client) RetryBatchCall(ctx context.Context, req *RetryBatchCallReq) (BatchCallResponse, error) {
	if req == nil {
		return BatchCallResponse{}, errors.New("request is nil")
	}

	body, err := c.post(ctx, "/convai/batch-calling/"+req.BatchId+"/retry", nil)
	if err != nil {
		return BatchCallResponse{}, err
	}

	var resp BatchCallResponse
	if err := c.parseResponse(body, &resp); err != nil {
		return BatchCallResponse{}, err
	}

	return resp, nil
}

// CancelBatchCallReq represents the request to cancel a batch call
type CancelBatchCallReq struct {
	BatchId string `path:"batch_id"`
}

func NewCancelBatchCallReq(batchId string) *CancelBatchCallReq {
	return &CancelBatchCallReq{
		BatchId: batchId,
	}
}

// CancelBatchCall cancels a running batch call and sets all recipients to cancelled status.
// https://elevenlabs.io/docs/api-reference/batch-calling/cancel
func (c *Client) CancelBatchCall(ctx context.Context, req *CancelBatchCallReq) (BatchCallResponse, error) {
	if req == nil {
		return BatchCallResponse{}, errors.New("request is nil")
	}

	body, err := c.post(ctx, "/convai/batch-calling/"+req.BatchId+"/cancel", nil)
	if err != nil {
		return BatchCallResponse{}, err
	}

	var resp BatchCallResponse
	if err := c.parseResponse(body, &resp); err != nil {
		return BatchCallResponse{}, err
	}

	return resp, nil
}
