package elevenlabs

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/go-querystring/query"
)

type ConvaiKnowledgeBaseDocumentsAPI interface {
	ListKnowledgeBaseDocuments(ctx context.Context, req *ListKnowledgeBaseDocumentsReq) (ListKnowledgeBaseDocumentsResp, error)
	DeleteKnowledgeBaseDocument(ctx context.Context, req *DeleteKnowledgeBaseDocumentReq) error
	GetKnowledgeBaseDocument(ctx context.Context, req *GetKnowledgeBaseDocumentReq) (KnowledgeBaseDocumentResponse, error)
	UpdateKnowledgeBaseDocument(ctx context.Context, req *UpdateKnowledgeBaseDocumentReq) (KnowledgeBaseDocumentResponse, error)
	CreateKnowledgeBaseDocumentFromURL(ctx context.Context, req *CreateKnowledgeBaseDocumentFromURLReq) (AddKnowledgeBaseResp, error)
	CreateKnowledgeBaseDocumentFromText(ctx context.Context, req *CreateKnowledgeBaseDocumentFromTextReq) (AddKnowledgeBaseResp, error)
	CreateKnowledgeBaseDocumentFromFile(ctx context.Context, req *CreateKnowledgeBaseDocumentFromFileReq) (AddKnowledgeBaseResp, error)
	GetDocumentContent(ctx context.Context, req *GetDocumentContentReq) (string, error)
	GetDocumentChunk(ctx context.Context, req *GetDocumentChunkReq) (KnowledgeBaseDocumentChunkResponseModel, error)
}

// SortDirection represents the direction to sort results
type SortDirection string

const (
	SortDirectionAsc  SortDirection = "asc"
	SortDirectionDesc SortDirection = "desc"
)

// KnowledgeBaseSortBy represents fields to sort knowledge base documents by
type KnowledgeBaseSortBy string

const (
	KnowledgeBaseSortByName      KnowledgeBaseSortBy = "name"
	KnowledgeBaseSortByCreatedAt KnowledgeBaseSortBy = "created_at"
	KnowledgeBaseSortByUpdatedAt KnowledgeBaseSortBy = "updated_at"
	KnowledgeBaseSortBySize      KnowledgeBaseSortBy = "size"
)

// ListKnowledgeBaseDocumentsReq represents the request for listing knowledge base documents
type ListKnowledgeBaseDocumentsReq struct {
	PageSize             int                      `url:"page_size,omitempty"`
	Search               *string                  `url:"search,omitempty"`
	ShowOnlyOwnedDocuments *bool                  `url:"show_only_owned_documents,omitempty"`
	Types                []KnowledgeBaseDocumentType `url:"types,omitempty"`
	ParentFolderID       *string                  `url:"parent_folder_id,omitempty"`
	AncestorFolderID     *string                  `url:"ancestor_folder_id,omitempty"`
	FoldersFirst         *bool                    `url:"folders_first,omitempty"`
	SortDirection        *SortDirection           `url:"sort_direction,omitempty"`
	SortBy               *KnowledgeBaseSortBy     `url:"sort_by,omitempty"`
	UseTypesense         *bool                    `url:"use_typesense,omitempty"`
	Cursor               *string                  `url:"cursor,omitempty"`
}

func NewListKnowledgeBaseDocumentsReq() *ListKnowledgeBaseDocumentsReq {
	return &ListKnowledgeBaseDocumentsReq{
		PageSize: defaultPageSize,
	}
}

func (r ListKnowledgeBaseDocumentsReq) QueryString() string {
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

// KnowledgeBaseFolderPathSegment represents a segment in the folder path
type KnowledgeBaseFolderPathSegment struct {
	Id   string  `json:"id"`
	Name *string `json:"name"`
}

// KnowledgeBaseURLDocument represents a URL-type knowledge base document
type KnowledgeBaseURLDocument struct {
	Id              string                        `json:"id"`
	Name            string                        `json:"name"`
	Type            KnowledgeBaseDocumentType     `json:"type"`
	Metadata        KnowledgeBaseDocumentMetadata `json:"metadata"`
	SupportedUsages []DocumentUsageMode           `json:"supported_usages"`
	AccessInfo      ResourceAccessInfo            `json:"access_info"`
	FolderParentId  *string                       `json:"folder_parent_id,omitempty"`
	FolderPath      []KnowledgeBaseFolderPathSegment `json:"folder_path"`
	URL             string                        `json:"url"`
	ExtractedInnerHTML string                     `json:"extracted_inner_html"`
}

// KnowledgeBaseFileDocument represents a file-type knowledge base document
type KnowledgeBaseFileDocument struct {
	Id              string                        `json:"id"`
	Name            string                        `json:"name"`
	Type            KnowledgeBaseDocumentType     `json:"type"`
	Metadata        KnowledgeBaseDocumentMetadata `json:"metadata"`
	SupportedUsages []DocumentUsageMode           `json:"supported_usages"`
	AccessInfo      ResourceAccessInfo            `json:"access_info"`
	FolderParentId  *string                       `json:"folder_parent_id,omitempty"`
	FolderPath      []KnowledgeBaseFolderPathSegment `json:"folder_path"`
	ExtractedInnerHTML string                     `json:"extracted_inner_html"`
}

// KnowledgeBaseTextDocument represents a text-type knowledge base document
type KnowledgeBaseTextDocument struct {
	Id              string                        `json:"id"`
	Name            string                        `json:"name"`
	Type            KnowledgeBaseDocumentType     `json:"type"`
	Metadata        KnowledgeBaseDocumentMetadata `json:"metadata"`
	SupportedUsages []DocumentUsageMode           `json:"supported_usages"`
	AccessInfo      ResourceAccessInfo            `json:"access_info"`
	FolderParentId  *string                       `json:"folder_parent_id,omitempty"`
	FolderPath      []KnowledgeBaseFolderPathSegment `json:"folder_path"`
	ExtractedInnerHTML string                     `json:"extracted_inner_html"`
}

// KnowledgeBaseFolderDocument represents a folder-type knowledge base document
type KnowledgeBaseFolderDocument struct {
	Id              string                        `json:"id"`
	Name            string                        `json:"name"`
	Type            KnowledgeBaseDocumentType     `json:"type"`
	Metadata        KnowledgeBaseDocumentMetadata `json:"metadata"`
	SupportedUsages []DocumentUsageMode           `json:"supported_usages"`
	AccessInfo      ResourceAccessInfo            `json:"access_info"`
	FolderParentId  *string                       `json:"folder_parent_id,omitempty"`
	FolderPath      []KnowledgeBaseFolderPathSegment `json:"folder_path"`
	ChildrenCount   int                           `json:"children_count"`
}

// KnowledgeBaseDocumentResponse is a union type for all document response types
type KnowledgeBaseDocumentResponse struct {
	URLDocument    *KnowledgeBaseURLDocument
	FileDocument   *KnowledgeBaseFileDocument
	TextDocument   *KnowledgeBaseTextDocument
	FolderDocument *KnowledgeBaseFolderDocument
}

// UnmarshalJSON implements custom unmarshaling for the union type
func (k *KnowledgeBaseDocumentResponse) UnmarshalJSON(data []byte) error {
	// Try each type until one succeeds
	var urlDoc KnowledgeBaseURLDocument
	if err := json.Unmarshal(data, &urlDoc); err == nil && urlDoc.Type == KnowledgeBaseDocumentTypeURL {
		k.URLDocument = &urlDoc
		return nil
	}

	var fileDoc KnowledgeBaseFileDocument
	if err := json.Unmarshal(data, &fileDoc); err == nil && fileDoc.Type == KnowledgeBaseDocumentTypeFile {
		k.FileDocument = &fileDoc
		return nil
	}

	var textDoc KnowledgeBaseTextDocument
	if err := json.Unmarshal(data, &textDoc); err == nil && textDoc.Type == KnowledgeBaseDocumentTypeText {
		k.TextDocument = &textDoc
		return nil
	}

	var folderDoc KnowledgeBaseFolderDocument
	if err := json.Unmarshal(data, &folderDoc); err == nil && folderDoc.Type == KnowledgeBaseDocumentTypeFolder {
		k.FolderDocument = &folderDoc
		return nil
	}

	return errors.New("unable to unmarshal as any knowledge base document type")
}

// ListKnowledgeBaseDocumentsResp represents the response for listing knowledge base documents
type ListKnowledgeBaseDocumentsResp struct {
	Documents  []KnowledgeBaseDocumentResponse `json:"documents"`
	NextCursor *string                         `json:"next_cursor,omitempty"`
	HasMore    bool                            `json:"has_more"`
}

// ListKnowledgeBaseDocuments gets a list of available knowledge base documents.
// https://elevenlabs.io/docs/api-reference/knowledge-base/list
func (c *Client) ListKnowledgeBaseDocuments(ctx context.Context, req *ListKnowledgeBaseDocumentsReq) (ListKnowledgeBaseDocumentsResp, error) {
	if req == nil {
		req = NewListKnowledgeBaseDocumentsReq()
	}

	body, err := c.get(ctx, "/convai/knowledge-base"+req.QueryString())
	if err != nil {
		return ListKnowledgeBaseDocumentsResp{}, err
	}

	var resp ListKnowledgeBaseDocumentsResp
	if err := c.parseResponse(body, &resp); err != nil {
		return ListKnowledgeBaseDocumentsResp{}, err
	}

	return resp, nil
}

// DeleteKnowledgeBaseDocumentReq represents the request for deleting a knowledge base document
type DeleteKnowledgeBaseDocumentReq struct {
	DocumentationId string `path:"documentation_id"`
	Force           *bool  `url:"force,omitempty"`
}

func NewDeleteKnowledgeBaseDocumentReq(documentationId string) *DeleteKnowledgeBaseDocumentReq {
	return &DeleteKnowledgeBaseDocumentReq{
		DocumentationId: documentationId,
	}
}

func (r DeleteKnowledgeBaseDocumentReq) QueryString() string {
	v, err := query.Values(r)
	if err != nil {
		return ""
	}
	// Remove documentation_id from query params as it's a path param
	v.Del("documentation_id")
	qs := v.Encode()
	if qs == "" {
		return ""
	}
	return "?" + qs
}

// DeleteKnowledgeBaseDocument deletes a document from the knowledge base.
// https://elevenlabs.io/docs/api-reference/knowledge-base/delete
func (c *Client) DeleteKnowledgeBaseDocument(ctx context.Context, req *DeleteKnowledgeBaseDocumentReq) error {
	if req == nil {
		return errors.New("request is nil")
	}

	if err := c.delete(ctx, "/convai/knowledge-base/"+req.DocumentationId+req.QueryString()); err != nil {
		return err
	}

	return nil
}

// GetKnowledgeBaseDocumentReq represents the request for getting a knowledge base document
type GetKnowledgeBaseDocumentReq struct {
	DocumentationId string `path:"documentation_id"`
	AgentId         *string `url:"agent_id,omitempty"`
}

func NewGetKnowledgeBaseDocumentReq(documentationId string) *GetKnowledgeBaseDocumentReq {
	return &GetKnowledgeBaseDocumentReq{
		DocumentationId: documentationId,
	}
}

func (r GetKnowledgeBaseDocumentReq) QueryString() string {
	v, err := query.Values(r)
	if err != nil {
		return ""
	}
	// Remove documentation_id from query params as it's a path param
	v.Del("documentation_id")
	qs := v.Encode()
	if qs == "" {
		return ""
	}
	return "?" + qs
}

// GetKnowledgeBaseDocument gets details about a specific knowledge base document.
// https://elevenlabs.io/docs/api-reference/knowledge-base/get-document
func (c *Client) GetKnowledgeBaseDocument(ctx context.Context, req *GetKnowledgeBaseDocumentReq) (KnowledgeBaseDocumentResponse, error) {
	if req == nil {
		return KnowledgeBaseDocumentResponse{}, errors.New("request is nil")
	}

	body, err := c.get(ctx, "/convai/knowledge-base/"+req.DocumentationId+req.QueryString())
	if err != nil {
		return KnowledgeBaseDocumentResponse{}, err
	}

	var resp KnowledgeBaseDocumentResponse
	if err := c.parseResponse(body, &resp); err != nil {
		return KnowledgeBaseDocumentResponse{}, err
	}

	return resp, nil
}

// UpdateKnowledgeBaseDocumentReq represents the request for updating a knowledge base document
type UpdateKnowledgeBaseDocumentReq struct {
	DocumentationId string `path:"documentation_id"`
	Name            string `json:"name"`
}

func NewUpdateKnowledgeBaseDocumentReq(documentationId string, name string) *UpdateKnowledgeBaseDocumentReq {
	return &UpdateKnowledgeBaseDocumentReq{
		DocumentationId: documentationId,
		Name:            name,
	}
}

// UpdateKnowledgeBaseDocument updates the name of a document.
// https://elevenlabs.io/docs/api-reference/knowledge-base/update
func (c *Client) UpdateKnowledgeBaseDocument(ctx context.Context, req *UpdateKnowledgeBaseDocumentReq) (KnowledgeBaseDocumentResponse, error) {
	if req == nil {
		return KnowledgeBaseDocumentResponse{}, errors.New("request is nil")
	}

	body, err := c.patch(ctx, "/convai/knowledge-base/"+req.DocumentationId, req)
	if err != nil {
		return KnowledgeBaseDocumentResponse{}, err
	}

	var resp KnowledgeBaseDocumentResponse
	if err := c.parseResponse(body, &resp); err != nil {
		return KnowledgeBaseDocumentResponse{}, err
	}

	return resp, nil
}

// CreateKnowledgeBaseDocumentFromURLReq represents the request for creating a knowledge base document from URL
type CreateKnowledgeBaseDocumentFromURLReq struct {
	URL            string  `json:"url"`
	Name           *string `json:"name,omitempty"`
	ParentFolderID *string `json:"parent_folder_id,omitempty"`
}

func NewCreateKnowledgeBaseDocumentFromURLReq(url string) *CreateKnowledgeBaseDocumentFromURLReq {
	return &CreateKnowledgeBaseDocumentFromURLReq{
		URL: url,
	}
}

// AddKnowledgeBaseResp represents the response for adding a knowledge base document
type AddKnowledgeBaseResp struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// CreateKnowledgeBaseDocumentFromURL creates a knowledge base document by scraping a webpage.
// https://elevenlabs.io/docs/api-reference/knowledge-base/create-from-url
func (c *Client) CreateKnowledgeBaseDocumentFromURL(ctx context.Context, req *CreateKnowledgeBaseDocumentFromURLReq) (AddKnowledgeBaseResp, error) {
	if req == nil {
		return AddKnowledgeBaseResp{}, errors.New("request is nil")
	}

	body, err := c.post(ctx, "/convai/knowledge-base/url", req)
	if err != nil {
		return AddKnowledgeBaseResp{}, err
	}

	var resp AddKnowledgeBaseResp
	if err := c.parseResponse(body, &resp); err != nil {
		return AddKnowledgeBaseResp{}, err
	}

	return resp, nil
}

// CreateKnowledgeBaseDocumentFromTextReq represents the request for creating a knowledge base document from text
type CreateKnowledgeBaseDocumentFromTextReq struct {
	Text           string  `json:"text"`
	Name           *string `json:"name,omitempty"`
	ParentFolderID *string `json:"parent_folder_id,omitempty"`
}

func NewCreateKnowledgeBaseDocumentFromTextReq(text string) *CreateKnowledgeBaseDocumentFromTextReq {
	return &CreateKnowledgeBaseDocumentFromTextReq{
		Text: text,
	}
}

// CreateKnowledgeBaseDocumentFromText creates a knowledge base document containing the provided text.
// https://elevenlabs.io/docs/api-reference/knowledge-base/create-from-text
func (c *Client) CreateKnowledgeBaseDocumentFromText(ctx context.Context, req *CreateKnowledgeBaseDocumentFromTextReq) (AddKnowledgeBaseResp, error) {
	if req == nil {
		return AddKnowledgeBaseResp{}, errors.New("request is nil")
	}

	body, err := c.post(ctx, "/convai/knowledge-base/text", req)
	if err != nil {
		return AddKnowledgeBaseResp{}, err
	}

	var resp AddKnowledgeBaseResp
	if err := c.parseResponse(body, &resp); err != nil {
		return AddKnowledgeBaseResp{}, err
	}

	return resp, nil
}

// CreateKnowledgeBaseDocumentFromFileReq represents the request for creating a knowledge base document from a file
type CreateKnowledgeBaseDocumentFromFileReq struct {
	File           *os.File
	Name           *string
	ParentFolderID *string
}

// CreateKnowledgeBaseDocumentFromFile creates a knowledge base document from an uploaded file.
// https://elevenlabs.io/docs/api-reference/knowledge-base/create-from-file
func (c *Client) CreateKnowledgeBaseDocumentFromFile(ctx context.Context, req *CreateKnowledgeBaseDocumentFromFileReq) (AddKnowledgeBaseResp, error) {
	if req == nil {
		return AddKnowledgeBaseResp{}, errors.New("request is nil")
	}

	if req.File == nil {
		return AddKnowledgeBaseResp{}, errors.New("file is required")
	}

	// Create multipart form
	var b strings.Builder
	writer := multipart.NewWriter(&b)

	// Add file field
	fileWriter, err := writer.CreateFormFile("file", filepath.Base(req.File.Name()))
	if err != nil {
		return AddKnowledgeBaseResp{}, err
	}

	_, err = io.Copy(fileWriter, req.File)
	if err != nil {
		return AddKnowledgeBaseResp{}, err
	}

	// Add name field if provided
	if req.Name != nil {
		writer.WriteField("name", *req.Name)
	}

	// Add parent_folder_id field if provided
	if req.ParentFolderID != nil {
		writer.WriteField("parent_folder_id", *req.ParentFolderID)
	}

	contentType := writer.FormDataContentType()
	writer.Close()

	httpReq, err := http.NewRequestWithContext(ctx, "POST", c.baseURL.String()+"/v1/convai/knowledge-base/file", strings.NewReader(b.String()))
	if err != nil {
		return AddKnowledgeBaseResp{}, err
	}

	httpReq.Header.Set("Content-Type", contentType)
	httpReq = c.prepareRequest(ctx, httpReq)

	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return AddKnowledgeBaseResp{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		body, _ := io.ReadAll(resp.Body)
		return AddKnowledgeBaseResp{}, fmt.Errorf("API error: %d - %s", resp.StatusCode, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return AddKnowledgeBaseResp{}, err
	}

	var result AddKnowledgeBaseResp
	if err := c.parseResponse(body, &result); err != nil {
		return AddKnowledgeBaseResp{}, err
	}

	return result, nil
}

// GetDocumentContentReq represents the request for getting document content
type GetDocumentContentReq struct {
	DocumentationId string `path:"documentation_id"`
}

func NewGetDocumentContentReq(documentationId string) *GetDocumentContentReq {
	return &GetDocumentContentReq{
		DocumentationId: documentationId,
	}
}

// GetDocumentContent gets the entire content of a document from the knowledge base.
// https://elevenlabs.io/docs/api-reference/knowledge-base/get-content
func (c *Client) GetDocumentContent(ctx context.Context, req *GetDocumentContentReq) (string, error) {
	if req == nil {
		return "", errors.New("request is nil")
	}

	body, err := c.get(ctx, "/convai/knowledge-base/"+req.DocumentationId+"/content")
	if err != nil {
		return "", err
	}

	// The response is streaming content, just return as string
	return string(body), nil
}

// GetDocumentChunkReq represents the request for getting a document chunk
type GetDocumentChunkReq struct {
	DocumentationId string `path:"documentation_id"`
	ChunkId         string `path:"chunk_id"`
}

func NewGetDocumentChunkReq(documentationId, chunkId string) *GetDocumentChunkReq {
	return &GetDocumentChunkReq{
		DocumentationId: documentationId,
		ChunkId:         chunkId,
	}
}

// KnowledgeBaseDocumentChunkResponseModel represents a knowledge base document chunk
type KnowledgeBaseDocumentChunkResponseModel struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

// GetDocumentChunk gets details about a specific documentation part used by RAG.
// https://elevenlabs.io/docs/api-reference/knowledge-base/get-chunk
func (c *Client) GetDocumentChunk(ctx context.Context, req *GetDocumentChunkReq) (KnowledgeBaseDocumentChunkResponseModel, error) {
	if req == nil {
		return KnowledgeBaseDocumentChunkResponseModel{}, errors.New("request is nil")
	}

	body, err := c.get(ctx, "/convai/knowledge-base/"+req.DocumentationId+"/chunk/"+req.ChunkId)
	if err != nil {
		return KnowledgeBaseDocumentChunkResponseModel{}, err
	}

	var resp KnowledgeBaseDocumentChunkResponseModel
	if err := c.parseResponse(body, &resp); err != nil {
		return KnowledgeBaseDocumentChunkResponseModel{}, err
	}

	return resp, nil
}
