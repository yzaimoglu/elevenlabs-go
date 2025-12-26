package elevenlabs

import (
	"context"
	"errors"
)

type ConvaiSecretsAPI interface {
	ListSecrets(ctx context.Context) (ListSecretsResp, error)
	CreateSecret(ctx context.Context, req *CreateSecretReq) (CreateSecretResp, error)
	UpdateSecret(ctx context.Context, req *UpdateSecretReq) (CreateSecretResp, error)
	DeleteSecret(ctx context.Context, req *DeleteSecretReq) error
}

// SecretDependencyType represents a type of secret dependency
type SecretDependencyType string

const (
	SecretDependencyTypeConversationInitiationWebhook SecretDependencyType = "conversation_initiation_webhook"
)

// DependentToolAccessLevel represents the access level for a dependent tool
type DependentToolAccessLevel string

const (
	DependentToolAccessLevelAdmin     DependentToolAccessLevel = "admin"
	DependentToolAccessLevelEditor    DependentToolAccessLevel = "editor"
	DependentToolAccessLevelCommenter DependentToolAccessLevel = "commenter"
	DependentToolAccessLevelViewer    DependentToolAccessLevel = "viewer"
)

// DependentToolIdentifier represents a dependent tool
type DependentToolIdentifier struct {
	Id              string                   `json:"id,omitempty"`
	Name            string                   `json:"name,omitempty"`
	Type            string                   `json:"type,omitempty"`
	CreatedAtUnixSecs int64                  `json:"created_at_unix_secs,omitempty"`
	AccessLevel     DependentToolAccessLevel `json:"access_level,omitempty"`
}

// DependentPhoneNumberIdentifier represents a dependent phone number
type DependentPhoneNumberIdentifier struct {
	PhoneNumberId string            `json:"phone_number_id"`
	PhoneNumber   string            `json:"phone_number"`
	Label         string            `json:"label"`
	Provider      TelephonyProvider `json:"provider"`
}

// SecretDependencies represents the dependencies of a secret
type SecretDependencies struct {
	Tools        []DependentToolIdentifier        `json:"tools"`
	Agents       []DependentAgentIdentifier       `json:"agents"`
	Others       []SecretDependencyType           `json:"others"`
	PhoneNumbers []DependentPhoneNumberIdentifier `json:"phone_numbers,omitempty"`
}

// WorkspaceSecret represents a workspace secret
type WorkspaceSecret struct {
	Type     string             `json:"type"`
	SecretId string             `json:"secret_id"`
	Name     string             `json:"name"`
	UsedBy   SecretDependencies `json:"used_by"`
}

// ListSecretsResp represents the response from listing secrets
type ListSecretsResp struct {
	Secrets []WorkspaceSecret `json:"secrets"`
}

// ListSecrets retrieves all workspace secrets for the user.
// https://elevenlabs.io/docs/api-reference/workspace/secrets/list
func (c *Client) ListSecrets(ctx context.Context) (ListSecretsResp, error) {
	body, err := c.get(ctx, "/convai/secrets")
	if err != nil {
		return ListSecretsResp{}, err
	}

	var resp ListSecretsResp
	if err := c.parseResponse(body, &resp); err != nil {
		return ListSecretsResp{}, err
	}

	return resp, nil
}

// CreateSecretReq represents the request for creating a secret
type CreateSecretReq struct {
	Type  string `json:"type"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

func NewCreateSecretReq(name, value string) *CreateSecretReq {
	return &CreateSecretReq{
		Type:  "new",
		Name:  name,
		Value: value,
	}
}

// CreateSecretResp represents the response from creating a secret
type CreateSecretResp struct {
	Type     string `json:"type"`
	SecretId string `json:"secret_id"`
	Name     string `json:"name"`
}

// CreateSecret creates a new secret for the workspace.
// https://elevenlabs.io/docs/api-reference/workspace/secrets/create
func (c *Client) CreateSecret(ctx context.Context, req *CreateSecretReq) (CreateSecretResp, error) {
	if req == nil {
		return CreateSecretResp{}, errors.New("request is nil")
	}

	body, err := c.post(ctx, "/convai/secrets", req)
	if err != nil {
		return CreateSecretResp{}, err
	}

	var resp CreateSecretResp
	if err := c.parseResponse(body, &resp); err != nil {
		return CreateSecretResp{}, err
	}

	return resp, nil
}

// UpdateSecretReq represents the request for updating a secret
type UpdateSecretReq struct {
	SecretId string `path:"secret_id"`
	Type     string `json:"type"`
	Name     string `json:"name"`
	Value    string `json:"value"`
}

func NewUpdateSecretReq(secretId, name, value string) *UpdateSecretReq {
	return &UpdateSecretReq{
		SecretId: secretId,
		Type:     "update",
		Name:     name,
		Value:    value,
	}
}

// UpdateSecret updates a secret for the workspace.
// https://elevenlabs.io/docs/api-reference/workspace/secrets/update
func (c *Client) UpdateSecret(ctx context.Context, req *UpdateSecretReq) (CreateSecretResp, error) {
	if req == nil {
		return CreateSecretResp{}, errors.New("request is nil")
	}

	body, err := c.patch(ctx, "/convai/secrets/"+req.SecretId, req)
	if err != nil {
		return CreateSecretResp{}, err
	}

	var resp CreateSecretResp
	if err := c.parseResponse(body, &resp); err != nil {
		return CreateSecretResp{}, err
	}

	return resp, nil
}

// DeleteSecretReq represents the request for deleting a secret
type DeleteSecretReq struct {
	SecretId string `path:"secret_id"`
}

func NewDeleteSecretReq(secretId string) *DeleteSecretReq {
	return &DeleteSecretReq{
		SecretId: secretId,
	}
}

// DeleteSecret deletes a secret from the workspace.
// https://elevenlabs.io/docs/api-reference/workspace/secrets/delete
func (c *Client) DeleteSecret(ctx context.Context, req *DeleteSecretReq) error {
	if req == nil {
		return errors.New("request is nil")
	}

	return c.delete(ctx, "/convai/secrets/"+req.SecretId)
}
