package elevenlabs

type TokenPurpose string

const (
	TokenPurposeSignedUrl     TokenPurpose = "signed_url"
	TokenPurposeShareableLink TokenPurpose = "shareable_link"
)
