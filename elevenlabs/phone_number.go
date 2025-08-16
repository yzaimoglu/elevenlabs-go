package elevenlabs

type PhoneNumber struct {
	PhoneNumber      string                            `json:"phone_number"`
	Label            string                            `json:"label"`
	PhoneNumberId    string                            `json:"phone_number_id"`
	SupportsInbound  *bool                             `json:"supports_inbound,omitempty"`
	SupportsOutbound *bool                             `json:"supports_outbound,omitempty"`
	AssignedAgent    *PhoneNumberAssignedAgent         `json:"assigned_agent,omitempty"`
	Provider         *PhoneNumberProvider              `json:"provider,omitempty"`
	OutboundTrunk    *PhoneNumberSIPTrunkOutboundTrunk `json:"outbound_trunk,omitempty"`
	InboundTrunk     *PhoneNumberSIPTrunkInboundTrunk  `json:"inbound_trunk,omitempty"`
}

type PhoneNumberAssignedAgent struct {
	AgentId   string `json:"agent_id"`
	AgentName string `json:"agent_name"`
}

type PhoneNumberProvider string

const (
	PhoneNumberProviderTwilio   PhoneNumberProvider = "twilio"
	PhoneNumberProviderSIPTrunk PhoneNumberProvider = "sip_trunk"
)

type GetPhoneNumberSIPTrunkResponseModel struct {
}

type PhoneNumberSIPTrunkOutboundTrunk struct {
	Address            string            `json:"address"`
	Transport          TrunkTransport    `json:"transport"`
	MediaEncryption    MediaEncryption   `json:"media_encryption"`
	HasAuthCredentials bool              `json:"has_auth_credentials"`
	Headers            map[string]string `json:"headers,omitempty"`
	Username           *string           `json:"username,omitempty"`
	HasOutboundTrunk   *bool             `json:"has_outbound_trunk,omitempty"`
}

type PhoneNumberSIPTrunkInboundTrunk struct {
	AllowedAddresses   []string        `json:"allowed_addresses"`
	AllowedNumbers     []string        `json:"allowed_numbers,omitempty"`
	MediaEncryption    MediaEncryption `json:"media_encryption"`
	HasAuthCredentials *bool           `json:"has_auth_credentials,omitempty"`
	Username           *string         `json:"username,omitempty"`
}

type TrunkTransport string

const (
	TrunkTransportAuto TrunkTransport = "auto"
	TrunkTransportUDP  TrunkTransport = "udp"
	TrunkTransportTCP  TrunkTransport = "tcp"
	TrunkTransportTLS  TrunkTransport = "tls"
)

type MediaEncryption string

const (
	MediaEncryptionDisabled MediaEncryption = "disabled"
	MediaEncryptionAllowed  MediaEncryption = "allowed"
	MediaEncryptionRequired MediaEncryption = "required"
)
