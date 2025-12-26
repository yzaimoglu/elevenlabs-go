package elevenlabs

import (
	"context"
	"errors"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"
)

type ConvaiKnowledgeBaseAPI interface {
	GetKnowledgeBaseDependentAgents(ctx context.Context, req *GetKnowledgeBaseDependentAgentsReq) (GetKnowledgeBaseDependentAgentsResp, error)
	GetKnowledgeBaseSize(ctx context.Context, req *GetKnowledgeBaseSizeReq) (GetKnowledgeBaseSizeResp, error)
	GetKnowledgeBaseSummaries(ctx context.Context, req *GetKnowledgeBaseSummariesReq) (map[string]KnowledgeBaseSummaryBatchResponse, error)
}

// KnowledgeBaseDependentType represents the type of dependent agents to return
type KnowledgeBaseDependentType string

const (
	KnowledgeBaseDependentTypeDirect     KnowledgeBaseDependentType = "direct"
	KnowledgeBaseDependentTypeTransitive KnowledgeBaseDependentType = "transitive"
	KnowledgeBaseDependentTypeAll        KnowledgeBaseDependentType = "all"
)

// GetKnowledgeBaseDependentAgentsReq represents the request for getting dependent agents
type GetKnowledgeBaseDependentAgentsReq struct {
	DocumentationId string                      `path:"documentation_id"`
	DependentType   *KnowledgeBaseDependentType `url:"dependent_type,omitempty"`
	PageSize        int                         `url:"page_size,omitempty"`
	Cursor          *string                     `url:"cursor,omitempty"`
}

func NewGetKnowledgeBaseDependentAgentsReq(documentationId string) *GetKnowledgeBaseDependentAgentsReq {
	return &GetKnowledgeBaseDependentAgentsReq{
		DocumentationId: documentationId,
		PageSize:        defaultPageSize,
	}
}

func (r GetKnowledgeBaseDependentAgentsReq) QueryString() string {
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

// GetKnowledgeBaseDependentAgentsResp represents the response for getting dependent agents
type GetKnowledgeBaseDependentAgentsResp struct {
	Agents     []DependentAgent `json:"agents"`
	NextCursor *string          `json:"next_cursor,omitempty"`
	HasMore    bool             `json:"has_more"`
}

// GetKnowledgeBaseDependentAgents gets a list of agents depending on this knowledge base document.
// https://elevenlabs.io/docs/api-reference/knowledge-base/get-agents
func (c *Client) GetKnowledgeBaseDependentAgents(ctx context.Context, req *GetKnowledgeBaseDependentAgentsReq) (GetKnowledgeBaseDependentAgentsResp, error) {
	if req == nil {
		return GetKnowledgeBaseDependentAgentsResp{}, errors.New("request is nil")
	}

	body, err := c.get(ctx, "/convai/knowledge-base/"+req.DocumentationId+"/dependent-agents"+req.QueryString())
	if err != nil {
		return GetKnowledgeBaseDependentAgentsResp{}, err
	}

	var resp GetKnowledgeBaseDependentAgentsResp
	if err := c.parseResponse(body, &resp); err != nil {
		return GetKnowledgeBaseDependentAgentsResp{}, err
	}

	return resp, nil
}

// GetKnowledgeBaseSizeReq represents the request for getting knowledge base size
type GetKnowledgeBaseSizeReq struct {
	AgentId string `path:"agent_id"`
}

func NewGetKnowledgeBaseSizeReq(agentId string) *GetKnowledgeBaseSizeReq {
	return &GetKnowledgeBaseSizeReq{
		AgentId: agentId,
	}
}

// GetKnowledgeBaseSizeResp represents the response for getting knowledge base size
type GetKnowledgeBaseSizeResp struct {
	NumberOfPages float64 `json:"number_of_pages"`
}

// GetKnowledgeBaseSize returns the number of pages in the agent's knowledge base.
// https://elevenlabs.io/docs/api-reference/knowledge-base/size
func (c *Client) GetKnowledgeBaseSize(ctx context.Context, req *GetKnowledgeBaseSizeReq) (GetKnowledgeBaseSizeResp, error) {
	if req == nil {
		return GetKnowledgeBaseSizeResp{}, errors.New("request is nil")
	}

	body, err := c.get(ctx, "/convai/agent/"+req.AgentId+"/knowledge-base/size")
	if err != nil {
		return GetKnowledgeBaseSizeResp{}, err
	}

	var resp GetKnowledgeBaseSizeResp
	if err := c.parseResponse(body, &resp); err != nil {
		return GetKnowledgeBaseSizeResp{}, err
	}

	return resp, nil
}

// GetKnowledgeBaseSummariesReq represents the request for getting knowledge base summaries
type GetKnowledgeBaseSummariesReq struct {
	DocumentIds []string `url:"document_ids"`
}

func NewGetKnowledgeBaseSummariesReq(documentIds []string) *GetKnowledgeBaseSummariesReq {
	return &GetKnowledgeBaseSummariesReq{
		DocumentIds: documentIds,
	}
}

func (r GetKnowledgeBaseSummariesReq) QueryString() string {
	if len(r.DocumentIds) == 0 {
		return ""
	}
	params := url.Values{}
	for _, id := range r.DocumentIds {
		params.Add("document_ids", id)
	}
	return "?" + params.Encode()
}

// DocumentUsageMode represents how a document can be used
type DocumentUsageMode string

const (
	DocumentUsageModePrompt DocumentUsageMode = "prompt"
	DocumentUsageModeAuto   DocumentUsageMode = "auto"
)

// KnowledgeBaseDocumentType represents the type of knowledge base document
type KnowledgeBaseDocumentType string

const (
	KnowledgeBaseDocumentTypeURL    KnowledgeBaseDocumentType = "url"
	KnowledgeBaseDocumentTypeFile   KnowledgeBaseDocumentType = "file"
	KnowledgeBaseDocumentTypeText   KnowledgeBaseDocumentType = "text"
	KnowledgeBaseDocumentTypeFolder KnowledgeBaseDocumentType = "folder"
)

// KnowledgeBaseDocumentMetadata contains metadata about a knowledge base document
type KnowledgeBaseDocumentMetadata struct {
	CreatedAtUnixSecs     int64 `json:"created_at_unix_secs"`
	LastUpdatedAtUnixSecs int64 `json:"last_updated_at_unix_secs"`
	SizeBytes             int64 `json:"size_bytes"`
}

// KnowledgeBaseSummary represents a knowledge base document summary
type KnowledgeBaseSummary struct {
	Id              string                        `json:"id"`
	Name            string                        `json:"name"`
	Type            KnowledgeBaseDocumentType     `json:"type"`
	Metadata        KnowledgeBaseDocumentMetadata `json:"metadata"`
	SupportedUsages []DocumentUsageMode           `json:"supported_usages"`
	AccessInfo      ResourceAccessInfo            `json:"access_info"`
	FolderParentId  *string                       `json:"folder_parent_id,omitempty"`
	DependentAgents []DependentAgent              `json:"dependent_agents"`
	// URL-specific field
	URL *string `json:"url,omitempty"`
	// Folder-specific field
	ChildrenCount *int `json:"children_count,omitempty"`
}

// BatchResponseStatus represents the status of a batch response
type BatchResponseStatus string

const (
	BatchResponseStatusSuccess BatchResponseStatus = "success"
	BatchResponseStatusFailure BatchResponseStatus = "failure"
)

// KnowledgeBaseSummaryBatchResponse represents a batch response for a knowledge base summary
type KnowledgeBaseSummaryBatchResponse struct {
	Status       BatchResponseStatus   `json:"status"`
	Data         *KnowledgeBaseSummary `json:"data,omitempty"`
	ErrorCode    *int                  `json:"error_code,omitempty"`
	ErrorStatus  *string               `json:"error_status,omitempty"`
	ErrorMessage *string               `json:"error_message,omitempty"`
}

// GetKnowledgeBaseSummaries gets multiple knowledge base document summaries by their IDs.
// https://elevenlabs.io/docs/api-reference/knowledge-base/get-summaries
func (c *Client) GetKnowledgeBaseSummaries(ctx context.Context, req *GetKnowledgeBaseSummariesReq) (map[string]KnowledgeBaseSummaryBatchResponse, error) {
	if req == nil {
		return nil, errors.New("request is nil")
	}

	if len(req.DocumentIds) == 0 {
		return nil, errors.New("document_ids is required")
	}

	// Build query string with repeated document_ids parameters
	params := make([]string, 0, len(req.DocumentIds))
	for _, id := range req.DocumentIds {
		params = append(params, "document_ids="+url.QueryEscape(id))
	}
	queryString := "?" + strings.Join(params, "&")

	body, err := c.get(ctx, "/convai/knowledge-base/summaries"+queryString)
	if err != nil {
		return nil, err
	}

	var resp map[string]KnowledgeBaseSummaryBatchResponse
	if err := c.parseResponse(body, &resp); err != nil {
		return nil, err
	}

	return resp, nil
}
