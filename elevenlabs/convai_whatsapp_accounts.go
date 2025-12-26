package elevenlabs

import (
	"context"
	"errors"
)

type ConvaiWhatsAppAccountsAPI interface {
	ListWhatsAppAccounts(ctx context.Context) ([]WhatsAppAccount, error)
	GetWhatsAppAccount(ctx context.Context, req *GetWhatsAppAccountReq) (WhatsAppAccount, error)
	ImportWhatsAppAccount(ctx context.Context, req *ImportWhatsAppAccountReq) (ImportWhatsAppAccountResp, error)
	UpdateWhatsAppAccount(ctx context.Context, req *UpdateWhatsAppAccountReq) error
	DeleteWhatsAppAccount(ctx context.Context, req *DeleteWhatsAppAccountReq) error
}

// WhatsAppAccount represents a WhatsApp account
type WhatsAppAccount struct {
	BusinessAccountId   string  `json:"business_account_id"`
	PhoneNumberId       string  `json:"phone_number_id"`
	BusinessAccountName string  `json:"business_account_name"`
	PhoneNumberName     string  `json:"phone_number_name"`
	PhoneNumber         string  `json:"phone_number"`
	AssignedAgentId     *string `json:"assigned_agent_id,omitempty"`
	AssignedAgentName   *string `json:"assigned_agent_name"`
}

// ListWhatsAppAccountsResp represents the response from listing WhatsApp accounts
type ListWhatsAppAccountsResp struct {
	Items []WhatsAppAccount `json:"items"`
}

// ListWhatsAppAccounts retrieves all WhatsApp accounts in the workspace.
// https://elevenlabs.io/docs/api-reference/whats-app-accounts/list
func (c *Client) ListWhatsAppAccounts(ctx context.Context) ([]WhatsAppAccount, error) {
	body, err := c.get(ctx, "/convai/whatsapp-accounts")
	if err != nil {
		return nil, err
	}

	var resp ListWhatsAppAccountsResp
	if err := c.parseResponse(body, &resp); err != nil {
		return nil, err
	}

	return resp.Items, nil
}

// GetWhatsAppAccountReq represents the request for getting a WhatsApp account
type GetWhatsAppAccountReq struct {
	PhoneNumberId string `path:"phone_number_id"`
}

func NewGetWhatsAppAccountReq(phoneNumberId string) *GetWhatsAppAccountReq {
	return &GetWhatsAppAccountReq{
		PhoneNumberId: phoneNumberId,
	}
}

// GetWhatsAppAccount retrieves a WhatsApp account by phone number ID.
// https://elevenlabs.io/docs/api-reference/whats-app-accounts/get
func (c *Client) GetWhatsAppAccount(ctx context.Context, req *GetWhatsAppAccountReq) (WhatsAppAccount, error) {
	if req == nil {
		return WhatsAppAccount{}, errors.New("request is nil")
	}

	body, err := c.get(ctx, "/convai/whatsapp-accounts/"+req.PhoneNumberId)
	if err != nil {
		return WhatsAppAccount{}, err
	}

	var resp WhatsAppAccount
	if err := c.parseResponse(body, &resp); err != nil {
		return WhatsAppAccount{}, err
	}

	return resp, nil
}

// ImportWhatsAppAccountReq represents the request for importing a WhatsApp account
type ImportWhatsAppAccountReq struct {
	BusinessAccountId string `json:"business_account_id"`
	PhoneNumberId     string `json:"phone_number_id"`
	TokenCode         string `json:"token_code"`
}

func NewImportWhatsAppAccountReq(businessAccountId, phoneNumberId, tokenCode string) *ImportWhatsAppAccountReq {
	return &ImportWhatsAppAccountReq{
		BusinessAccountId: businessAccountId,
		PhoneNumberId:     phoneNumberId,
		TokenCode:         tokenCode,
	}
}

// ImportWhatsAppAccountResp represents the response from importing a WhatsApp account
type ImportWhatsAppAccountResp struct {
	PhoneNumberId string `json:"phone_number_id"`
}

// ImportWhatsAppAccount imports a WhatsApp account.
// https://elevenlabs.io/docs/api-reference/whats-app-accounts/import
func (c *Client) ImportWhatsAppAccount(ctx context.Context, req *ImportWhatsAppAccountReq) (ImportWhatsAppAccountResp, error) {
	if req == nil {
		return ImportWhatsAppAccountResp{}, errors.New("request is nil")
	}

	body, err := c.post(ctx, "/convai/whatsapp-accounts", req)
	if err != nil {
		return ImportWhatsAppAccountResp{}, err
	}

	var resp ImportWhatsAppAccountResp
	if err := c.parseResponse(body, &resp); err != nil {
		return ImportWhatsAppAccountResp{}, err
	}

	return resp, nil
}

// UpdateWhatsAppAccountReq represents the request for updating a WhatsApp account
type UpdateWhatsAppAccountReq struct {
	PhoneNumberId   string  `path:"phone_number_id"`
	AssignedAgentId *string `json:"assigned_agent_id,omitempty"`
}

func NewUpdateWhatsAppAccountReq(phoneNumberId string) *UpdateWhatsAppAccountReq {
	return &UpdateWhatsAppAccountReq{
		PhoneNumberId: phoneNumberId,
	}
}

// UpdateWhatsAppAccount updates a WhatsApp account.
// https://elevenlabs.io/docs/api-reference/whats-app-accounts/update
func (c *Client) UpdateWhatsAppAccount(ctx context.Context, req *UpdateWhatsAppAccountReq) error {
	if req == nil {
		return errors.New("request is nil")
	}

	_, err := c.patch(ctx, "/convai/whatsapp-accounts/"+req.PhoneNumberId, req)
	return err
}

// DeleteWhatsAppAccountReq represents the request for deleting a WhatsApp account
type DeleteWhatsAppAccountReq struct {
	PhoneNumberId string `path:"phone_number_id"`
}

func NewDeleteWhatsAppAccountReq(phoneNumberId string) *DeleteWhatsAppAccountReq {
	return &DeleteWhatsAppAccountReq{
		PhoneNumberId: phoneNumberId,
	}
}

// DeleteWhatsAppAccount deletes a WhatsApp account.
// https://elevenlabs.io/docs/api-reference/whats-app-accounts/delete
func (c *Client) DeleteWhatsAppAccount(ctx context.Context, req *DeleteWhatsAppAccountReq) error {
	if req == nil {
		return errors.New("request is nil")
	}

	return c.delete(ctx, "/convai/whatsapp-accounts/"+req.PhoneNumberId)
}
