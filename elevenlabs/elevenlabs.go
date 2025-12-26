package elevenlabs

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	baseURLFormat = "https://%s.elevenlabs.io/v1"
)

type Environment string

const (
	EnvironmentProduction      = "api"
	EnvironmentProductionUS    = "api.us"
	EnvironmentProductionEU    = "api.eu.residency"
	EnvironmentProductionIndia = "api.in.residency"
)

var defaultHeaders = map[string]string{
	"User-Agent":   "yzaimoglu/elevenlabs-go",
	"Content-Type": "application/json",
}

// Client of ElevenLabs API
type Client struct {
	baseURL    *url.URL
	httpClient *http.Client
	credential Credential
	headers    map[string]string
}

// BaseAPI encapsulates base methods for elevenlabs client
type BaseAPI interface {
	Get(ctx context.Context, path string) ([]byte, error)
	Post(ctx context.Context, path string, data any) ([]byte, error)
	Put(ctx context.Context, path string, data any) ([]byte, error)
	Delete(ctx context.Context, path string) error
}

// NewClient creates new Elevenlabs API client
// apiKey is your ElevenLabs API Key
// environment is the ElevenLabs API environment (Production, Production US, Production EU, Production India)
// customClient is a custom http client that can be provided
func NewClient(apiKey string, environment Environment, customClient ...*http.Client) (*Client, error) {
	var httpClient *http.Client
	if len(customClient) > 0 {
		httpClient = customClient[0]
	} else {
		httpClient = http.DefaultClient
	}

	client := &Client{httpClient: httpClient}
	client.headers = defaultHeaders
	client.SetEnvironment(environment)
	client.SetCredential(NewAuthCredential(apiKey))
	return client, nil
}

// SetHeader saves HTTP header in client. It will be included all API request
func (c *Client) SetHeader(key string, value string) {
	c.headers[key] = value
}

// SetEnvironment saves the environment in the client. It will be used
// when the API is called
func (c *Client) SetEnvironment(environment Environment) error {
	baseURLString := fmt.Sprintf(baseURLFormat, environment)
	baseURL, err := url.Parse(baseURLString)
	if err != nil {
		return err
	}

	c.baseURL = baseURL
	return nil
}

// SetCredential saves credential in client. It will be set
// to request header when call API
func (c *Client) SetCredential(cred Credential) {
	c.credential = cred
}

// get get JSON data from API and returns its body as []bytes
func (c *Client) get(ctx context.Context, path string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, c.baseURL.String()+path, nil)
	if err != nil {
		return nil, err
	}

	req = c.prepareRequest(ctx, req)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, ElevenlabsError{
			body: body,
			resp: resp,
		}
	}
	return body, nil
}

// post send data to API and returns response body as []bytes
func (c *Client) post(ctx context.Context, path string, data any) ([]byte, error) {
	bytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, c.baseURL.String()+path, strings.NewReader(string(bytes)))
	if err != nil {
		return nil, err
	}

	req = c.prepareRequest(ctx, req)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, ElevenlabsError{
			body: body,
			resp: resp,
		}
	}

	return body, nil
}

// put sends data to API and returns response body as []bytes
func (c *Client) put(ctx context.Context, path string, data any) ([]byte, error) {
	bytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPut, c.baseURL.String()+path, strings.NewReader(string(bytes)))
	if err != nil {
		return nil, err
	}

	req = c.prepareRequest(ctx, req)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, ElevenlabsError{
			body: body,
			resp: resp,
		}
	}

	return body, nil
}

// patch sends data to API and returns response body as []bytes
func (c *Client) patch(ctx context.Context, path string, data any) ([]byte, error) {
	bytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPatch, c.baseURL.String()+path, strings.NewReader(string(bytes)))
	if err != nil {
		return nil, err
	}

	req = c.prepareRequest(ctx, req)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// NOTE: some webhook mutation APIs return status No Content.
	if resp.StatusCode != http.StatusOK {
		return nil, ElevenlabsError{
			body: body,
			resp: resp,
		}
	}

	return body, nil
}

// delete sends data to API and returns an error if unsuccessful
func (c *Client) delete(ctx context.Context, path string) error {
	req, err := http.NewRequest(http.MethodDelete, c.baseURL.String()+path, nil)
	if err != nil {
		return err
	}

	req = c.prepareRequest(ctx, req)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return ElevenlabsError{
			body: body,
			resp: resp,
		}
	}

	return nil
}

// prepare request sets common request variables such as authn and user agent
func (c *Client) prepareRequest(ctx context.Context, req *http.Request) *http.Request {
	out := req.WithContext(ctx)
	c.includeHeaders(out)
	if c.credential != nil {
		out.Header.Add("xi-api-key", c.credential.ApiKey())
	}

	return out
}

// includeHeaders set HTTP headers from client.headers to *http.Request
func (c *Client) includeHeaders(req *http.Request) {
	for key, value := range c.headers {
		req.Header.Set(key, value)
	}
}

// parseResponse is a generic helper function that retrieves and unmarshals JSON data from a specified URL.
// It takes four parameters:
// - a pointer to a Client (c) which is used to execute the GET request,
// - a context (ctx) for managing the request's lifecycle,
// - a string (url) representing the endpoint from which data should be retrieved,
// - and an empty interface (data) where the retrieved data will be stored after being unmarshalled from JSON.
//
// The function starts by sending a GET request to the specified URL. If the request is successful,
// the returned body in the form of a byte slice is unmarshalled into the provided empty interface using the json.Unmarshal function.
//
// If an error occurs during either the GET request or the JSON unmarshalling, the function will return this error.
func (c *Client) parseResponse(res []byte, v any) error {
	return json.Unmarshal(res, v)
}

// Get allows users to send requests not yet implemented
func (c *Client) Get(ctx context.Context, path string) ([]byte, error) {
	return c.get(ctx, path)
}

// Post allows users to send requests not yet implemented
func (c *Client) Post(ctx context.Context, path string, data any) ([]byte, error) {
	return c.post(ctx, path, data)
}

// Put allows users to send requests not yet implemented
func (c *Client) Put(ctx context.Context, path string, data any) ([]byte, error) {
	return c.put(ctx, path, data)
}

// Delete allows users to send requests not yet implemented
func (c *Client) Delete(ctx context.Context, path string) error {
	return c.delete(ctx, path)
}
