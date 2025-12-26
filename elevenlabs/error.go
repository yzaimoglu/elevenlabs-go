package elevenlabs

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

// ElevenlabsError an error type containing the http response from elevenlabs
type ElevenlabsError struct {
	body []byte
	resp *http.Response
}

// NewError is a function to initialize the Error type. This function will be useful
// for unit testing and mocking purposes in the client side
// to test their behavior by the API response.
func NewError(statusCode int, body []byte, resp *http.Response) ElevenlabsError {
	return ElevenlabsError{
		body: body,
		resp: resp,
	}
}

// Error the error string for this error
func (e ElevenlabsError) Error() string {
	msg := string(e.body)
	if msg == "" {
		msg = http.StatusText(e.Status())
	}

	return fmt.Sprintf("%d: %s", e.resp.StatusCode, msg)
}

// Body is the Body of the HTTP response
func (e ElevenlabsError) Body() io.ReadCloser {
	return io.NopCloser(bytes.NewBuffer(e.body))
}

// Headers the HTTP headers returned from elevenlabs
func (e ElevenlabsError) Headers() http.Header {
	return e.resp.Header
}

// Status the HTTP status code returned from elevenlabs
func (e ElevenlabsError) Status() int {
	return e.resp.StatusCode
}

// OptionsError is an error type for invalid option argument.
type OptionsError struct {
	opts any
}

func (e *OptionsError) Error() string {
	return fmt.Sprintf("invalid options: %v", e.opts)
}
