package elevenlabs

// Credential is interface of API credential
type Credential interface {
	ApiKey() string
}

// AuthCredential is type of credential for API authentication
type AuthCredential struct {
	apiKey string
}

// NewAuthCredential creates AuthCredential and returns its pointer
func NewAuthCredential(apiKey string) *AuthCredential {
	return &AuthCredential{
		apiKey: apiKey,
	}
}

// ApiKey is accessor which returns API key
func (c AuthCredential) ApiKey() string {
	return c.apiKey
}
