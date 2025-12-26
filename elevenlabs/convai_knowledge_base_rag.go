package elevenlabs

import (
	"context"
	"errors"
	"io"
	"net/http"
)

type ConvaiKnowledgeBaseRAGAPI interface {
	ComputeRAGIndex(ctx context.Context, req *ComputeRAGIndexReq) (RAGDocumentIndexResponseModel, error)
	GetRAGIndex(ctx context.Context, req *GetRAGIndexReq) (RAGDocumentIndexesResponseModel, error)
	ComputeRAGIndexBatch(ctx context.Context, req *ComputeRAGIndexBatchReq) (map[string]RAGIndexBatchResponse, error)
	GetRAGIndexOverview(ctx context.Context) (RAGIndexOverviewResponseModel, error)
	DeleteRAGIndex(ctx context.Context, req *DeleteRAGIndexReq) (RAGDocumentIndexResponseModel, error)
}

// EmbeddingModel represents the model used for embeddings
type EmbeddingModel string

const (
	EmbeddingModelE5Mistral7BInstruct         EmbeddingModel = "e5_mistral_7b_instruct"
	EmbeddingModelMultilingualE5LargeInstruct EmbeddingModel = "multilingual_e5_large_instruct"
)

// RAGIndexStatus represents the status of a RAG index
type RAGIndexStatus string

const (
	RAGIndexStatusNotStarted    RAGIndexStatus = "not_started"
	RAGIndexStatusRunning       RAGIndexStatus = "running"
	RAGIndexStatusCompleted     RAGIndexStatus = "completed"
	RAGIndexStatusFailed        RAGIndexStatus = "failed"
	RAGIndexStatusOutOfCredits  RAGIndexStatus = "out_of_credits"
	RAGIndexStatusDeleted       RAGIndexStatus = "deleted"
	RAGIndexStatusWillBeDeleted RAGIndexStatus = "will_be_deleted"
)

// ComputeRAGIndexReq represents the request for computing a RAG index
type ComputeRAGIndexReq struct {
	DocumentationId string `path:"documentation_id"`
}

func NewComputeRAGIndexReq(documentationId string) *ComputeRAGIndexReq {
	return &ComputeRAGIndexReq{
		DocumentationId: documentationId,
	}
}

// RAGDocumentIndexResponseModel represents the response for a RAG document index
type RAGDocumentIndexResponseModel struct {
	Id                     string         `json:"id"`
	EmbeddingModel         EmbeddingModel `json:"embedding_model"`
	EmbeddingModelSettings any            `json:"embedding_model_settings"`
	Status                 RAGIndexStatus `json:"status"`
	Progress               float64        `json:"progress"`
	IndexedDocumentSize    int            `json:"indexed_document_size"`
	OutputPrice            float64        `json:"output_price"`
	Error                  *string        `json:"error,omitempty"`
}

// ComputeRAGIndex triggers rag indexing task or returns current status if already indexed.
// https://elevenlabs.io/docs/api-reference/knowledge-base/compute-rag-index
func (c *Client) ComputeRAGIndex(ctx context.Context, req *ComputeRAGIndexReq) (RAGDocumentIndexResponseModel, error) {
	if req == nil {
		return RAGDocumentIndexResponseModel{}, errors.New("request is nil")
	}

	body, err := c.post(ctx, "/convai/knowledge-base/"+req.DocumentationId+"/rag-index", nil)
	if err != nil {
		return RAGDocumentIndexResponseModel{}, err
	}

	var resp RAGDocumentIndexResponseModel
	if err := c.parseResponse(body, &resp); err != nil {
		return RAGDocumentIndexResponseModel{}, err
	}

	return resp, nil
}

// GetRAGIndexReq represents the request for getting RAG index
type GetRAGIndexReq struct {
	DocumentationId string `path:"documentation_id"`
}

func NewGetRAGIndexReq(documentationId string) *GetRAGIndexReq {
	return &GetRAGIndexReq{
		DocumentationId: documentationId,
	}
}

// RAGDocumentIndexesResponseModel represents the response for getting RAG document indexes
type RAGDocumentIndexesResponseModel struct {
	Indexes []RAGDocumentIndexResponseModel `json:"indexes"`
}

// GetRAGIndex provides information about all RAG indexes of the specified knowledgebase document.
// https://elevenlabs.io/docs/api-reference/knowledge-base/get-rag-index
func (c *Client) GetRAGIndex(ctx context.Context, req *GetRAGIndexReq) (RAGDocumentIndexesResponseModel, error) {
	if req == nil {
		return RAGDocumentIndexesResponseModel{}, errors.New("request is nil")
	}

	body, err := c.get(ctx, "/convai/knowledge-base/"+req.DocumentationId+"/rag-index")
	if err != nil {
		return RAGDocumentIndexesResponseModel{}, err
	}

	var resp RAGDocumentIndexesResponseModel
	if err := c.parseResponse(body, &resp); err != nil {
		return RAGDocumentIndexesResponseModel{}, err
	}

	return resp, nil
}

// ComputeRAGIndexBatchReq represents the request for computing RAG indexes in batch
type ComputeRAGIndexBatchReq struct {
	DocumentationIds []string `json:"documentation_ids"`
}

func NewComputeRAGIndexBatchReq(documentationIds []string) *ComputeRAGIndexBatchReq {
	return &ComputeRAGIndexBatchReq{
		DocumentationIds: documentationIds,
	}
}

// RAGIndexBatchResponse represents the response for batch RAG index operations
type RAGIndexBatchResponse struct {
	Status       BatchResponseStatus            `json:"status"`
	Data         *RAGDocumentIndexResponseModel `json:"data,omitempty"`
	ErrorCode    *int                           `json:"error_code,omitempty"`
	ErrorStatus  *string                        `json:"error_status,omitempty"`
	ErrorMessage *string                        `json:"error_message,omitempty"`
}

// ComputeRAGIndexBatch retrieves and/or creates RAG indexes for multiple documents in a single request.
// https://elevenlabs.io/docs/api-reference/knowledge-base/compute-rag-index-batch
func (c *Client) ComputeRAGIndexBatch(ctx context.Context, req *ComputeRAGIndexBatchReq) (map[string]RAGIndexBatchResponse, error) {
	if req == nil {
		return nil, errors.New("request is nil")
	}

	if len(req.DocumentationIds) == 0 {
		return nil, errors.New("documentation_ids is required")
	}

	body, err := c.post(ctx, "/convai/knowledge-base/rag-index", req)
	if err != nil {
		return nil, err
	}

	var resp map[string]RAGIndexBatchResponse
	if err := c.parseResponse(body, &resp); err != nil {
		return nil, err
	}

	return resp, nil
}

// GetRAGIndexOverview provides total size and other information of RAG indexes used by knowledgebase documents.
// https://elevenlabs.io/docs/api-reference/knowledge-base/rag-index-overview
func (c *Client) GetRAGIndexOverview(ctx context.Context) (RAGIndexOverviewResponseModel, error) {
	body, err := c.get(ctx, "/convai/knowledge-base/rag-index")
	if err != nil {
		return RAGIndexOverviewResponseModel{}, err
	}

	var resp RAGIndexOverviewResponseModel
	if err := c.parseResponse(body, &resp); err != nil {
		return RAGIndexOverviewResponseModel{}, err
	}

	return resp, nil
}

// DeleteRAGIndexReq represents the request for deleting a RAG index
type DeleteRAGIndexReq struct {
	DocumentationId string `path:"documentation_id"`
	RAGIndexId      string `path:"rag_index_id"`
}

func NewDeleteRAGIndexReq(documentationId, ragIndexId string) *DeleteRAGIndexReq {
	return &DeleteRAGIndexReq{
		DocumentationId: documentationId,
		RAGIndexId:      ragIndexId,
	}
}

// DeleteRAGIndex deletes a RAG index for the knowledgebase document.
// https://elevenlabs.io/docs/api-reference/knowledge-base/delete-rag-index
func (c *Client) DeleteRAGIndex(ctx context.Context, req *DeleteRAGIndexReq) (RAGDocumentIndexResponseModel, error) {
	if req == nil {
		return RAGDocumentIndexResponseModel{}, errors.New("request is nil")
	}

	body, err := c.deleteWithResponse(ctx, "/convai/knowledge-base/"+req.DocumentationId+"/rag-index/"+req.RAGIndexId)
	if err != nil {
		return RAGDocumentIndexResponseModel{}, err
	}

	var resp RAGDocumentIndexResponseModel
	if err := c.parseResponse(body, &resp); err != nil {
		return RAGDocumentIndexResponseModel{}, err
	}

	return resp, nil
}

// deleteWithResponse sends a DELETE request and returns the response body
func (c *Client) deleteWithResponse(ctx context.Context, path string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodDelete, c.baseURL.String()+path, nil)
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

// RAGIndexOverviewEmbeddingModelResponseModel represents embedding model information in the overview
type RAGIndexOverviewEmbeddingModelResponseModel struct {
	Name                    EmbeddingModel `json:"name"`
	TotalSize               int            `json:"total_size"`
	TotalPages              int            `json:"total_pages"`
	TotalDocuments          int            `json:"total_documents"`
	CreditsUsed             float64        `json:"credits_used"`
	CreditsUsedAcrossModels float64        `json:"credits_used_across_models"`
}

// RAGIndexOverviewResponseModel represents the overview of RAG indexes
type RAGIndexOverviewResponseModel struct {
	EmbeddingModels []RAGIndexOverviewEmbeddingModelResponseModel `json:"embedding_models"`
}
