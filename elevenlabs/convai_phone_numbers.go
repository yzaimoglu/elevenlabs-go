package elevenlabs

import (
	"context"
	"errors"
)

type ConvaiPhoneNumbersAPI interface {
	ListPhoneNumbers(ctx context.Context) ([]PhoneNumber, error)
	GetPhoneNumber(ctx context.Context, req *GetPhoneNumberReq) (PhoneNumber, error)
	CreatePhoneNumber(ctx context.Context, req *CreatePhoneNumberReq) (CreatePhoneNumberResp, error)
	UpdatePhoneNumber(ctx context.Context, req *UpdatePhoneNumberReq) (PhoneNumber, error)
	DeletePhoneNumber(ctx context.Context, req *DeletePhoneNumberReq) error
}

// TwilioRegionId represents a Twilio region ID
type TwilioRegionId string

const (
	TwilioRegionIdUS1 TwilioRegionId = "us1"
	TwilioRegionIdIE1 TwilioRegionId = "ie1"
	TwilioRegionIdAU1 TwilioRegionId = "au1"
)

// TwilioEdgeLocation represents a Twilio edge location
type TwilioEdgeLocation string

const (
	TwilioEdgeLocationAshburn   TwilioEdgeLocation = "ashburn"
	TwilioEdgeLocationDublin    TwilioEdgeLocation = "dublin"
	TwilioEdgeLocationFrankfurt TwilioEdgeLocation = "frankfurt"
	TwilioEdgeLocationSaoPaulo  TwilioEdgeLocation = "sao-paulo"
	TwilioEdgeLocationSingapore TwilioEdgeLocation = "singapore"
	TwilioEdgeLocationSydney    TwilioEdgeLocation = "sydney"
	TwilioEdgeLocationTokyo     TwilioEdgeLocation = "tokyo"
	TwilioEdgeLocationUmatilla  TwilioEdgeLocation = "umatilla"
	TwilioEdgeLocationRoaming   TwilioEdgeLocation = "roaming"
)

// LivekitStackType represents the type of Livekit stack
type LivekitStackType string

const (
	LivekitStackTypeStandard LivekitStackType = "standard"
	LivekitStackTypeStatic   LivekitStackType = "static"
)

// TwilioRegionConfig represents Twilio region configuration
type TwilioRegionConfig struct {
	RegionId     TwilioRegionId     `json:"region_id"`
	Token        string             `json:"token"`
	EdgeLocation TwilioEdgeLocation `json:"edge_location"`
}

// SIPTrunkCredentials represents SIP trunk authentication credentials
type SIPTrunkCredentials struct {
	Username string  `json:"username"`
	Password *string `json:"password,omitempty"`
}

// InboundSIPTrunkConfig represents inbound SIP trunk configuration for requests
type InboundSIPTrunkConfig struct {
	AllowedAddresses []string             `json:"allowed_addresses,omitempty"`
	AllowedNumbers   []string             `json:"allowed_numbers,omitempty"`
	MediaEncryption  *MediaEncryption     `json:"media_encryption,omitempty"`
	Credentials      *SIPTrunkCredentials `json:"credentials,omitempty"`
	RemoteDomains    []string             `json:"remote_domains,omitempty"`
}

// OutboundSIPTrunkConfig represents outbound SIP trunk configuration for requests
type OutboundSIPTrunkConfig struct {
	Address         string               `json:"address"`
	Transport       *TrunkTransport      `json:"transport,omitempty"`
	MediaEncryption *MediaEncryption     `json:"media_encryption,omitempty"`
	Headers         map[string]string    `json:"headers,omitempty"`
	Credentials     *SIPTrunkCredentials `json:"credentials,omitempty"`
}

// CreateTwilioPhoneNumberReq represents a request to create a Twilio phone number
type CreateTwilioPhoneNumberReq struct {
	PhoneNumber      string              `json:"phone_number"`
	Label            string              `json:"label"`
	SupportsInbound  *bool               `json:"supports_inbound,omitempty"`
	SupportsOutbound *bool               `json:"supports_outbound,omitempty"`
	Provider         PhoneNumberProvider `json:"provider"`
	SID              string              `json:"sid"`
	Token            string              `json:"token"`
	RegionConfig     *TwilioRegionConfig `json:"region_config,omitempty"`
}

// CreateSIPTrunkPhoneNumberReq represents a request to create a SIP trunk phone number
type CreateSIPTrunkPhoneNumberReq struct {
	PhoneNumber         string                  `json:"phone_number"`
	Label               string                  `json:"label"`
	SupportsInbound     *bool                   `json:"supports_inbound,omitempty"`
	SupportsOutbound    *bool                   `json:"supports_outbound,omitempty"`
	Provider            PhoneNumberProvider     `json:"provider"`
	InboundTrunkConfig  *InboundSIPTrunkConfig  `json:"inbound_trunk_config,omitempty"`
	OutboundTrunkConfig *OutboundSIPTrunkConfig `json:"outbound_trunk_config,omitempty"`
}

// CreatePhoneNumberReq represents a request to create a phone number (either Twilio or SIP trunk)
type CreatePhoneNumberReq struct {
	PhoneNumber         string                  `json:"phone_number"`
	Label               string                  `json:"label"`
	SupportsInbound     *bool                   `json:"supports_inbound,omitempty"`
	SupportsOutbound    *bool                   `json:"supports_outbound,omitempty"`
	Provider            PhoneNumberProvider     `json:"provider,omitempty"`
	SID                 string                  `json:"sid,omitempty"`
	Token               string                  `json:"token,omitempty"`
	RegionConfig        *TwilioRegionConfig     `json:"region_config,omitempty"`
	InboundTrunkConfig  *InboundSIPTrunkConfig  `json:"inbound_trunk_config,omitempty"`
	OutboundTrunkConfig *OutboundSIPTrunkConfig `json:"outbound_trunk_config,omitempty"`
}

func NewCreateTwilioPhoneNumberReq(phoneNumber, label, sid, token string) *CreatePhoneNumberReq {
	return &CreatePhoneNumberReq{
		PhoneNumber: phoneNumber,
		Label:       label,
		Provider:    PhoneNumberProviderTwilio,
		SID:         sid,
		Token:       token,
	}
}

func NewCreateSIPTrunkPhoneNumberReq(phoneNumber, label string) *CreatePhoneNumberReq {
	return &CreatePhoneNumberReq{
		PhoneNumber: phoneNumber,
		Label:       label,
		Provider:    PhoneNumberProviderSIPTrunk,
	}
}

// CreatePhoneNumberResp represents the response from creating a phone number
type CreatePhoneNumberResp struct {
	PhoneNumberId string `json:"phone_number_id"`
}

// CreatePhoneNumber imports a phone number from a provider configuration (Twilio or SIP trunk).
// https://elevenlabs.io/docs/api-reference/phone-numbers/create
func (c *Client) CreatePhoneNumber(ctx context.Context, req *CreatePhoneNumberReq) (CreatePhoneNumberResp, error) {
	if req == nil {
		return CreatePhoneNumberResp{}, errors.New("request is nil")
	}

	body, err := c.post(ctx, "/convai/phone-numbers", req)
	if err != nil {
		return CreatePhoneNumberResp{}, err
	}

	var resp CreatePhoneNumberResp
	if err := c.parseResponse(body, &resp); err != nil {
		return CreatePhoneNumberResp{}, err
	}

	return resp, nil
}

// ListPhoneNumbers retrieves all phone numbers in the workspace.
// https://elevenlabs.io/docs/api-reference/phone-numbers/list
func (c *Client) ListPhoneNumbers(ctx context.Context) ([]PhoneNumber, error) {
	body, err := c.get(ctx, "/convai/phone-numbers")
	if err != nil {
		return nil, err
	}

	var resp []PhoneNumber
	if err := c.parseResponse(body, &resp); err != nil {
		return nil, err
	}

	return resp, nil
}

// GetPhoneNumberReq represents the request for getting a phone number
type GetPhoneNumberReq struct {
	PhoneNumberId string `path:"phone_number_id"`
}

func NewGetPhoneNumberReq(phoneNumberId string) *GetPhoneNumberReq {
	return &GetPhoneNumberReq{
		PhoneNumberId: phoneNumberId,
	}
}

// GetPhoneNumber retrieves phone number details by ID.
// https://elevenlabs.io/docs/api-reference/phone-numbers/get
func (c *Client) GetPhoneNumber(ctx context.Context, req *GetPhoneNumberReq) (PhoneNumber, error) {
	if req == nil {
		return PhoneNumber{}, errors.New("request is nil")
	}

	body, err := c.get(ctx, "/convai/phone-numbers/"+req.PhoneNumberId)
	if err != nil {
		return PhoneNumber{}, err
	}

	var resp PhoneNumber
	if err := c.parseResponse(body, &resp); err != nil {
		return PhoneNumber{}, err
	}

	return resp, nil
}

// UpdatePhoneNumberReq represents the request for updating a phone number
type UpdatePhoneNumberReq struct {
	PhoneNumberId       string                  `path:"phone_number_id"`
	AgentId             *string                 `json:"agent_id,omitempty"`
	InboundTrunkConfig  *InboundSIPTrunkConfig  `json:"inbound_trunk_config,omitempty"`
	OutboundTrunkConfig *OutboundSIPTrunkConfig `json:"outbound_trunk_config,omitempty"`
	LivekitStack        *LivekitStackType       `json:"livekit_stack,omitempty"`
}

func NewUpdatePhoneNumberReq(phoneNumberId string) *UpdatePhoneNumberReq {
	return &UpdatePhoneNumberReq{
		PhoneNumberId: phoneNumberId,
	}
}

// UpdatePhoneNumber updates a phone number's configuration.
// https://elevenlabs.io/docs/api-reference/phone-numbers/update
func (c *Client) UpdatePhoneNumber(ctx context.Context, req *UpdatePhoneNumberReq) (PhoneNumber, error) {
	if req == nil {
		return PhoneNumber{}, errors.New("request is nil")
	}

	body, err := c.patch(ctx, "/convai/phone-numbers/"+req.PhoneNumberId, req)
	if err != nil {
		return PhoneNumber{}, err
	}

	var resp PhoneNumber
	if err := c.parseResponse(body, &resp); err != nil {
		return PhoneNumber{}, err
	}

	return resp, nil
}

// DeletePhoneNumberReq represents the request for deleting a phone number
type DeletePhoneNumberReq struct {
	PhoneNumberId string `path:"phone_number_id"`
}

func NewDeletePhoneNumberReq(phoneNumberId string) *DeletePhoneNumberReq {
	return &DeletePhoneNumberReq{
		PhoneNumberId: phoneNumberId,
	}
}

// DeletePhoneNumber deletes a phone number by ID.
// https://elevenlabs.io/docs/api-reference/phone-numbers/delete
func (c *Client) DeletePhoneNumber(ctx context.Context, req *DeletePhoneNumberReq) error {
	if req == nil {
		return errors.New("request is nil")
	}

	if err := c.delete(ctx, "/convai/phone-numbers/"+req.PhoneNumberId); err != nil {
		return err
	}

	return nil
}
