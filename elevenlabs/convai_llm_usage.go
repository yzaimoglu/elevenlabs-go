package elevenlabs

import (
	"context"
	"errors"
)

type ConvaiLLMUsageAPI interface {
	CalculateLLMUsage(ctx context.Context, req *CalculateLLMUsageReq) (CalculateLLMUsageResp, error)
}

// LLMModel represents an LLM model
type LLMModel string

const (
	LLMModelGPT4oMini       LLMModel = "gpt-4o-mini"
	LLMModelGPT4o           LLMModel = "gpt-4o"
	LLMModelGPT4            LLMModel = "gpt-4"
	LLMModelGPT4Turbo       LLMModel = "gpt-4-turbo"
	LLMModelGPT41           LLMModel = "gpt-4.1"
	LLMModelGPT41Mini       LLMModel = "gpt-4.1-mini"
	LLMModelGPT41Nano       LLMModel = "gpt-4.1-nano"
	LLMModelGPT5            LLMModel = "gpt-5"
	LLMModelGPT51           LLMModel = "gpt-5.1"
	LLMModelGPT52           LLMModel = "gpt-5.2"
	LLMModelGPT5Mini        LLMModel = "gpt-5-mini"
	LLMModelGPT5Nano        LLMModel = "gpt-5-nano"
	LLMModelGPT35Turbo      LLMModel = "gpt-3.5-turbo"
	LLMModelGemini15Pro     LLMModel = "gemini-1.5-pro"
	LLMModelGemini15Flash   LLMModel = "gemini-1.5-flash"
	LLMModelGemini20Flash   LLMModel = "gemini-2.0-flash"
	LLMModelGemini20FlashLite LLMModel = "gemini-2.0-flash-lite"
	LLMModelGemini25FlashLite LLMModel = "gemini-2.5-flash-lite"
	LLMModelGemini25Flash   LLMModel = "gemini-2.5-flash"
	LLMModelClaudeSonnet45  LLMModel = "claude-sonnet-4-5"
	LLMModelClaudeSonnet4   LLMModel = "claude-sonnet-4"
	LLMModelClaudeHaiku45   LLMModel = "claude-haiku-4-5"
	LLMModelClaude37Sonnet  LLMModel = "claude-3-7-sonnet"
	LLMModelClaude35Sonnet  LLMModel = "claude-3-5-sonnet"
	LLMModelClaude3Haiku    LLMModel = "claude-3-haiku"
	LLMModelGrokBeta        LLMModel = "grok-beta"
	LLMModelCustomLLM       LLMModel = "custom-llm"
)

// CalculateLLMUsageReq represents the request for calculating LLM usage
type CalculateLLMUsageReq struct {
	PromptLength  int  `json:"prompt_length"`
	NumberOfPages int  `json:"number_of_pages"`
	RAGEnabled    bool `json:"rag_enabled"`
}

func NewCalculateLLMUsageReq(promptLength, numberOfPages int, ragEnabled bool) *CalculateLLMUsageReq {
	return &CalculateLLMUsageReq{
		PromptLength:  promptLength,
		NumberOfPages: numberOfPages,
		RAGEnabled:    ragEnabled,
	}
}

// LLMUsagePrice represents the price for an LLM model
type LLMUsagePrice struct {
	LLM            LLMModel `json:"llm"`
	PricePerMinute float64  `json:"price_per_minute"`
}

// CalculateLLMUsageResp represents the response for calculating LLM usage
type CalculateLLMUsageResp struct {
	LLMPrices []LLMUsagePrice `json:"llm_prices"`
}

// CalculateLLMUsage returns a list of LLM models and the expected cost for using them.
// https://elevenlabs.io/docs/api-reference/llm-usage/calculate
func (c *Client) CalculateLLMUsage(ctx context.Context, req *CalculateLLMUsageReq) (CalculateLLMUsageResp, error) {
	if req == nil {
		return CalculateLLMUsageResp{}, errors.New("request is nil")
	}

	body, err := c.post(ctx, "/convai/llm-usage/calculate", req)
	if err != nil {
		return CalculateLLMUsageResp{}, err
	}

	var resp CalculateLLMUsageResp
	if err := c.parseResponse(body, &resp); err != nil {
		return CalculateLLMUsageResp{}, err
	}

	return resp, nil
}
