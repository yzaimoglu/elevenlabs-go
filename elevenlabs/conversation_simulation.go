package elevenlabs

type SimulationExtraEvaluationCriterium struct {
	ID                     string                                  `json:"id"`
	Name                   string                                  `json:"name"`
	ConversationGoalPrompt string                                  `json:"conversation_goal_prompt"`
	Type                   *SimulationExtraEvaluationCriteriumType `json:"type,omitempty"`
	UseKnowledgeBase       *bool                                   `json:"use_knowledge_base,omitempty"`
}

type SimulationExtraEvaluationCriteriumType string

const (
	SimulationExtraEvaluationCriteriumTypePrompt SimulationExtraEvaluationCriteriumType = "prompt"
)

type SimulationSpecification struct {
	SimulatedUserConfig        SimulatedUserConfig          `json:"simulated_user_config"`
	ToolMockConfig             map[string]ToolMockConfig    `json:"tool_mock_config,omitempty"`
	PartialConversationHistory []PartialConversationHistory `json:"partial_conversation_history,omitempty"`
	DynamicVariables           map[string]any               `json:"dynamic_variables,omitempty"`
}

type SimulatedUserConfig struct {
	FirstMessage     *string                `json:"first_message,omitempty"`
	Language         *string                `json:"language,omitempty"`
	DynamicVariables *AgentDynamicVariables `json:"dynamic_variables,omitempty"`
	Prompt           *AgentPrompt           `json:"prompt,omitempty"`
}

type ToolMockConfig struct {
	DefaultReturnValue *string `json:"default_return_value,omitempty"`
	DefaultIsError     *bool   `json:"default_is_error,omitempty"`
}

type PartialConversationHistory struct {
	Role                    ConversationRole         `json:"role"`
	TimeInCallSecs          int64                    `json:"time_in_call_secs"`
	Message                 *string                  `json:"message,omitempty"`
	ToolCalls               []ConversationToolCall   `json:"tool_calls,omitempty"`
	ToolResults             []ConversationToolResult `json:"tool_results,omitempty"`
	Feedback                *ConversationFeedback    `json:"feedback,omitempty"`
	LLMOverride             *string                  `json:"llm_override,omitempty"`
	ConversationTurnMetrics *ConversationTurnMetrics `json:"conversation_turn_metrics,omitempty"`
	RAGRetrievalInfo        *RAGRetrievalInfo        `json:"rag_retrieval_info,omitempty"`
	LLMUsage                *ConversationLLMUsage    `json:"llm_usage,omitempty"`
	Interrupted             *bool                    `json:"interrupted,omitempty"`
	OriginalMessage         *string                  `json:"original_message,omitempty"`
	SourceMedium            *SourceMedium            `json:"source_medium,omitempty"`
}

type ConversationRole string

const (
	ConversationRoleUser  ConversationRole = "user"
	ConversationRoleAgent ConversationRole = "agent"
)

type ConversationToolCall struct {
	RequestId         string                    `json:"request_id"`
	ToolName          string                    `json:"tool_name"`
	ParamsAsJson      string                    `json:"params_as_json"`
	ToolHasBeenCalled bool                      `json:"tool_has_been_called"`
	Type              *ConversationToolCallType `json:"type,omitempty"`
	ToolDetails       any                       `json:"tool_details,omitempty"`
}

type ConversationToolCallType string

const (
	ConversationToolCallTypeSystem  ConversationToolCallType = "system"
	ConversationToolCallTypeWebhook ConversationToolCallType = "webhook"
	ConversationToolCallTypeClient  ConversationToolCallType = "client"
	ConversationToolCallTypeMCP     ConversationToolCallType = "mcp"
)

type ConversationToolResult struct {
	RequestId              string                        `json:"request_id"`
	ToolName               string                        `json:"tool_name"`
	ResultValue            string                        `json:"result_value"`
	IsError                bool                          `json:"is_error"`
	ToolHasBeenCalled      bool                          `json:"tool_has_been_called"`
	Type                   ConversationToolCallType      `json:"type"`
	ToolLatencySecs        *float32                      `json:"tool_latency_secs,omitempty"`
	DynamicVariableUpdates []DynamicVariableUpdate       `json:"dynamic_variable_updates,omitempty"`
	Result                 *ConversationToolResultResult `json:"result,omitempty"`
}

type DynamicVariableUpdate struct {
	VariableName  string  `json:"variable_name"`
	OldValue      string  `json:"old_value"`
	NewValue      string  `json:"new_value"`
	UpdatedAt     float32 `json:"updated_at"`
	ToolName      string  `json:"tool_name"`
	ToolRequestId string  `json:"tool_request_id"`
}

type ConversationToolResultResult struct {
	ResultType *string `json:"result_type,omitempty"`
	Status     *string `json:"status,omitempty"`
	Reason     *string `json:"reason,omitempty"`
	Error      *string `json:"error,omitempty"`
	Details    *string `json:"details,omitempty"`
	// End call
	Message *string `json:"message,omitempty"`
	// Language
	Language *string `json:"language,omitempty"`
	// TransferToAgent
	FromAgent                          *string `json:"from_agent,omitempty"`
	ToAgent                            *string `json:"to_agent,omitempty"`
	Condition                          *string `json:"condition,omitempty"`
	DelayMs                            *int    `json:"delay_ms,omitempty"`
	TransferMessage                    *string `json:"transfer_message,omitempty"`
	EnableTransferredAgentFirstMessage *bool   `json:"enable_transferred_agent_first_message,omitempty"`
	// TransferToNumber
	TransferNumber *string `json:"transfer_number,omitempty"`
	Note           *string `json:"note,omitempty"`
	// TransferToNumberTwilio
	AgentMessage   *string `json:"agent_message,omitempty"`
	ConferenceName *string `json:"conference_name,omitempty"`
	ClientMessage  *string `json:"client_message,omitempty"`
	// PlayDTMF
	DTMFTones *string `json:"dtmf_tones,omitempty"`
	// VoiceMailDetection
	VoicemailMessage *string `json:"voicemail_message,omitempty"`
}

type ConversationFeedback struct {
	Score          ConversationFeedbackScore `json:"score"`
	TimeInCallSecs int64                     `json:"time_in_call_secs"`
}

type ConversationFeedbackScore string

const (
	ConversationFeedbackScoreLike    ConversationFeedbackScore = "like"
	ConversationFeedbackScoreDislike ConversationFeedbackScore = "dislike"
)

type ConversationTurnMetrics struct {
	Metrics map[string]ConversationTurnMetric `json:"metrics,omitempty"`
}

type ConversationTurnMetric struct {
	ElapsedTime float32 `json:"elapsed_time"`
}

type RAGRetrievalInfo struct {
	Chunks         []RAGRetrievalInfoChunk `json:"chunks,omitempty"`
	EmbeddingModel RAGEmbeddingModel       `json:"embedding_model,omitempty"`
	RetrievalQuery string                  `json:"retrieval_query,omitempty"`
	RAGLatencySecs float32                 `json:"rag_latency_secs,omitempty"`
}

type RAGRetrievalInfoChunk struct {
	DocumentId     string  `json:"document_id"`
	ChunkId        string  `json:"chunk_id"`
	VectorDistance float32 `json:"vector_distance"`
}

type ConversationLLMUsage struct {
	ModelUsage map[string]LLMModelUsage `json:"model_usage,omitempty"`
}

type LLMModelUsage struct {
	Input           *LLMModelUsageTokens `json:"input,omitempty"`
	InputCacheRead  *LLMModelUsageTokens `json:"input_cache_read,omitempty"`
	InputCacheWrite *LLMModelUsageTokens `json:"input_cache_write,omitempty"`
	OutputTotal     *LLMModelUsageTokens `json:"output_total,omitempty"`
}

type LLMModelUsageTokens struct {
	Tokens int     `json:"tokens"`
	Price  float32 `json:"price"`
}

type SourceMedium string

const (
	SourceMediumAudio SourceMedium = "audio"
	SourceMediumText  SourceMedium = "text"
)

type SimulatedConversation struct {
	Role                    ConversationRole         `json:"role"`
	TimeInCallSecs          int64                    `json:"time_in_call_secs"`
	Message                 *string                  `json:"message,omitempty"`
	ToolCalls               []ConversationToolCall   `json:"tool_calls,omitempty"`
	ToolResults             []ConversationToolResult `json:"tool_results,omitempty"`
	Feedback                *ConversationFeedback    `json:"feedback,omitempty"`
	LLMOverride             *string                  `json:"llm_override,omitempty"`
	ConversationTurnMetrics *ConversationTurnMetrics `json:"conversation_turn_metrics,omitempty"`
	RAGRetrievalInfo        *RAGRetrievalInfo        `json:"rag_retrieval_info,omitempty"`
	LLMUsage                *ConversationLLMUsage    `json:"llm_usage,omitempty"`
	Interrupted             *bool                    `json:"interrupted,omitempty"`
	OriginalMessage         *string                  `json:"original_message,omitempty"`
	SourceMedium            *SourceMedium            `json:"source_medium,omitempty"`
}

type SimulatedConversationAnalysis struct {
	CallSuccessful            SimulatedConversationSuccessStatus  `json:"call_successful"`
	TranscriptSummary         string                              `json:"transcript_summary"`
	EvaluationCriteriaResults map[string]EvaluationCriteriaResult `json:"evaluation_criteria_results,omitempty"`
	DataCollectionResults     map[string]DataCollectionResult     `json:"data_collection_results,omitempty"`
	CallSummaryTitle          *string                             `json:"call_summary_title,omitempty"`
}

type SimulatedConversationSuccessStatus string

const (
	SimulatedConversationSuccessStatusSuccess SimulatedConversationSuccessStatus = "success"
	SimulatedConversationSuccessStatusFailure SimulatedConversationSuccessStatus = "failure"
	SimulatedConversationSuccessStatusUnknown SimulatedConversationSuccessStatus = "unknown"
)

type EvaluationCriteriaResult struct {
	CriteriaId string                             `json:"criteria_id"`
	Result     SimulatedConversationSuccessStatus `json:"result"`
	Rationale  string                             `json:"rationale"`
}

type DataCollectionResult struct {
	DataCollectionId string                `json:"data_collection_id"`
	Rationale        string                `json:"rationale"`
	Value            any                   `json:"value"`
	JsonSchema       *DataCollectionSchema `json:"json_schema,omitempty"`
}
