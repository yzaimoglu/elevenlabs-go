package elevenlabs

// API an interface containing all of the zendesk client methods
type API interface {
	BaseAPI
	ConvaiAPI
}

var _ API = (*Client)(nil)
