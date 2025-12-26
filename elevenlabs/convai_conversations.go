package elevenlabs

import (
	"context"
	"errors"

	"github.com/google/go-querystring/query"
)

type ConvaiConversationsAPI interface {
	ListConversations(ctx context.Context, req *ListConversationsReq) (ListConversationsResp, error)
	GetConversation(ctx context.Context, req *GetConversationReq) (GetConversationResp, error)
	GetConversationAudio(ctx context.Context, req *GetConversationAudioReq) ([]byte, error)
	DeleteConversation(ctx context.Context, req *DeleteConversationReq) error
	GetConversationToken(ctx context.Context, req *GetConversationTokenReq) (GetConversationTokenResp, error)
	GetSignedUrl(ctx context.Context, req *GetSignedUrlReq) (GetSignedUrlResp, error)
	SendConversationFeedback(ctx context.Context, req *SendConversationFeedbackReq) error
}

// ListConversationsReq represents the request for listing conversations.
type ListConversationsReq struct {
	Cursor               *string                  `url:"cursor,omitempty"`
	AgentId              *string                  `url:"agent_id,omitempty"`
	CallSuccessful       *EvaluationSuccessResult `url:"call_successful,omitempty"`
	CallStartBeforeUnix  *int64                   `url:"call_start_before_unix,omitempty"`
	CallStartAfterUnix   *int64                   `url:"call_start_after_unix,omitempty"`
	CallDurationMinSecs  *int                     `url:"call_duration_min_secs,omitempty"`
	CallDurationMaxSecs  *int                     `url:"call_duration_max_secs,omitempty"`
	RatingMax            *int                     `url:"rating_max,omitempty"`
	RatingMin            *int                     `url:"rating_min,omitempty"`
	HasFeedbackComment   *bool                    `url:"has_feedback_comment,omitempty"`
	UserId               *string                  `url:"user_id,omitempty"`
	EvaluationParams     []string                 `url:"evaluation_params,omitempty"`
	DataCollectionParams []string                 `url:"data_collection_params,omitempty"`
	ToolNames            []string                 `url:"tool_names,omitempty"`
	MainLanguages        []string                 `url:"main_languages,omitempty"`
	PageSize             int                      `url:"page_size,omitempty"`
	SummaryMode          *ConversationSummaryMode `url:"summary_mode,omitempty"`
	Search               *string                  `url:"search,omitempty"`
}

type EvaluationSuccessResult string

const (
	EvaluationSuccessResultSuccess EvaluationSuccessResult = "success"
	EvaluationSuccessResultFailure EvaluationSuccessResult = "failure"
	EvaluationSuccessResultUnknown EvaluationSuccessResult = "unknown"
)

type ConversationSummaryMode string

const (
	ConversationSummaryModeExclude ConversationSummaryMode = "exclude"
	ConversationSummaryModeInclude ConversationSummaryMode = "include"
)

func NewListConversationsReq() *ListConversationsReq {
	return &ListConversationsReq{
		PageSize: defaultPageSize,
	}
}

func (r ListConversationsReq) QueryString() string {
	v, err := query.Values(r)
	if err != nil {
		return ""
	}
	return "?" + v.Encode()
}

type ConversationStatus string

const (
	ConversationStatusInitiated  ConversationStatus = "initiated"
	ConversationStatusInProgress ConversationStatus = "in-progress"
	ConversationStatusProcessing ConversationStatus = "processing"
	ConversationStatusDone       ConversationStatus = "done"
	ConversationStatusFailed     ConversationStatus = "failed"
)

type ConversationDirection string

const (
	ConversationDirectionInbound  ConversationDirection = "inbound"
	ConversationDirectionOutbound ConversationDirection = "outbound"
)

type ConversationSummary struct {
	AgentId           string                  `json:"agent_id"`
	BranchId          *string                 `json:"branch_id,omitempty"`
	AgentName         *string                 `json:"agent_name,omitempty"`
	ConversationId    string                  `json:"conversation_id"`
	StartTimeUnixSecs int64                   `json:"start_time_unix_secs"`
	CallDurationSecs  int                     `json:"call_duration_secs"`
	MessageCount      int                     `json:"message_count"`
	Status            ConversationStatus      `json:"status"`
	CallSuccessful    EvaluationSuccessResult `json:"call_successful"`
	TranscriptSummary *string                 `json:"transcript_summary,omitempty"`
	CallSummaryTitle  *string                 `json:"call_summary_title,omitempty"`
	Direction         *ConversationDirection  `json:"direction,omitempty"`
	Rating            *float64                `json:"rating,omitempty"`
}

type ListConversationsResp struct {
	Conversations []ConversationSummary `json:"conversations"`
	NextCursor    *string               `json:"next_cursor,omitempty"`
	HasMore       bool                  `json:"has_more"`
}

// ListConversations gets all conversations of agents that user owns.
// https://elevenlabs.io/docs/api-reference/conversations/list
func (c *Client) ListConversations(ctx context.Context, req *ListConversationsReq) (ListConversationsResp, error) {
	if req == nil {
		req = NewListConversationsReq()
	}

	body, err := c.get(ctx, "/convai/conversations"+req.QueryString())
	if err != nil {
		return ListConversationsResp{}, err
	}

	var resp ListConversationsResp
	if err := c.parseResponse(body, &resp); err != nil {
		return ListConversationsResp{}, err
	}

	return resp, nil
}

// GetConversationReq represents the request for getting a conversation.
type GetConversationReq struct {
	ConversationId string `path:"conversation_id"`
}

func NewGetConversationReq(conversationId string) *GetConversationReq {
	return &GetConversationReq{
		ConversationId: conversationId,
	}
}

type TranscriptRole string

const (
	TranscriptRoleUser  TranscriptRole = "user"
	TranscriptRoleAgent TranscriptRole = "agent"
)

type TranscriptSourceMedium string

const (
	TranscriptSourceMediumAudio TranscriptSourceMedium = "audio"
	TranscriptSourceMediumText  TranscriptSourceMedium = "text"
)

type UserFeedbackScore string

const (
	UserFeedbackScoreLike    UserFeedbackScore = "like"
	UserFeedbackScoreDislike UserFeedbackScore = "dislike"
)

type UserFeedback struct {
	Score          UserFeedbackScore `json:"score"`
	TimeInCallSecs int               `json:"time_in_call_secs"`
}

// TranscriptMessageToolType represents the type of tool used in transcript
type TranscriptMessageToolType string

const (
	TranscriptMessageToolTypeSystem                TranscriptMessageToolType = "system"
	TranscriptMessageToolTypeWebhook               TranscriptMessageToolType = "webhook"
	TranscriptMessageToolTypeClient                TranscriptMessageToolType = "client"
	TranscriptMessageToolTypeMcp                   TranscriptMessageToolType = "mcp"
	TranscriptMessageToolTypeWorkflow              TranscriptMessageToolType = "workflow"
	TranscriptMessageToolTypeApiIntegrationWebhook TranscriptMessageToolType = "api_integration_webhook"
	TranscriptMessageToolTypeApiIntegrationMcp     TranscriptMessageToolType = "api_integration_mcp"
)

// TranscriptMessageToolCallCommonModel represents a tool call in the transcript
type TranscriptMessageToolCallCommonModel struct {
	Type              *TranscriptMessageToolType                       `json:"type,omitempty"`
	RequestId         string                                           `json:"request_id"`
	ToolName          string                                           `json:"tool_name"`
	ParamsAsJson      string                                           `json:"params_as_json"`
	ToolHasBeenCalled bool                                             `json:"tool_has_been_called"`
	ToolDetails       *TranscriptMessageToolCallCommonModelToolDetails `json:"tool_details,omitempty"`
}

// TranscriptMessageToolCallCommonModelToolDetails is a union type for tool details
type TranscriptMessageToolCallCommonModelToolDetails struct {
	Webhook               *TranscriptMessageToolCallWebhookDetails
	Client                *TranscriptMessageToolCallClientDetails
	MCP                   *TranscriptMessageToolCallMCPDetails
	ApiIntegrationWebhook *TranscriptMessageToolCallApiIntegrationWebhookDetails
}

// TranscriptMessageToolCallWebhookDetails represents webhook tool call details
type TranscriptMessageToolCallWebhookDetails struct {
	Type        string            `json:"type"`
	Method      string            `json:"method"`
	URL         string            `json:"url"`
	Headers     map[string]string `json:"headers,omitempty"`
	PathParams  map[string]string `json:"path_params,omitempty"`
	QueryParams map[string]string `json:"query_params,omitempty"`
	Body        *string           `json:"body,omitempty"`
}

// TranscriptMessageToolCallClientDetails represents client tool call details
type TranscriptMessageToolCallClientDetails struct {
	Type       string `json:"type"`
	Parameters string `json:"parameters"`
}

// TranscriptMessageToolCallMCPDetails represents MCP tool call details
type TranscriptMessageToolCallMCPDetails struct {
	Type               string            `json:"type"`
	McpServerId        string            `json:"mcp_server_id"`
	McpServerName      string            `json:"mcp_server_name"`
	IntegrationType    string            `json:"integration_type"`
	Parameters         map[string]string `json:"parameters,omitempty"`
	ApprovalPolicy     string            `json:"approval_policy"`
	RequiresApproval   bool              `json:"requires_approval"`
	McpToolName        string            `json:"mcp_tool_name"`
	McpToolDescription string            `json:"mcp_tool_description"`
}

// TranscriptMessageToolCallApiIntegrationWebhookDetails represents API integration webhook details
type TranscriptMessageToolCallApiIntegrationWebhookDetails struct {
	Type                    string                                   `json:"type"`
	IntegrationId           string                                   `json:"integration_id"`
	CredentialId            string                                   `json:"credential_id"`
	IntegrationConnectionId string                                   `json:"integration_connection_id"`
	WebhookDetails          *TranscriptMessageToolCallWebhookDetails `json:"webhook_details,omitempty"`
}

// BaseToolResult contains common fields for all tool results
type BaseToolResult struct {
	RequestId              string                  `json:"request_id"`
	ToolName               string                  `json:"tool_name"`
	ResultValue            string                  `json:"result_value"`
	IsError                bool                    `json:"is_error"`
	ToolHasBeenCalled      bool                    `json:"tool_has_been_called"`
	ToolLatencySecs        *float64                `json:"tool_latency_secs,omitempty"`
	DynamicVariableUpdates []DynamicVariableUpdate `json:"dynamic_variable_updates,omitempty"`
}

// TranscriptMessageToolResultCommonModel represents a tool result in the transcript
type TranscriptMessageToolResultCommonModel struct {
	*BaseToolResult
	Type *string `json:"type,omitempty"` // "client", "webhook", "mcp", "api_integration_webhook", "workflow", or "system"
}

type TranscriptMessage struct {
	Role            TranscriptRole                           `json:"role"`
	AgentMetadata   *AgentMetadata                           `json:"agent_metadata,omitempty"`
	Message         *string                                  `json:"message,omitempty"`
	ToolCalls       []TranscriptMessageToolCallCommonModel   `json:"tool_calls,omitempty"`
	ToolResults     []TranscriptMessageToolResultCommonModel `json:"tool_results,omitempty"`
	Feedback        *UserFeedback                            `json:"feedback,omitempty"`
	LLMOverride     *string                                  `json:"llm_override,omitempty"`
	TimeInCallSecs  int                                      `json:"time_in_call_secs"`
	Interrupted     bool                                     `json:"interrupted"`
	OriginalMessage *string                                  `json:"original_message,omitempty"`
	SourceMedium    *TranscriptSourceMedium                  `json:"source_medium,omitempty"`
}

type AgentMetadata struct {
	AgentId        string  `json:"agent_id"`
	BranchId       *string `json:"branch_id,omitempty"`
	WorkflowNodeId *string `json:"workflow_node_id,omitempty"`
}

type ConversationMetadata struct {
	StartTimeUnixSecs    int64                          `json:"start_time_unix_secs"`
	AcceptedTimeUnixSecs *int64                         `json:"accepted_time_unix_secs,omitempty"`
	CallDurationSecs     int                            `json:"call_duration_secs"`
	Cost                 *int                           `json:"cost,omitempty"`
	DeletionSettings     *ConversationDeletionSettings  `json:"deletion_settings,omitempty"`
	Feedback             *ConversationFeedbackSummary   `json:"feedback,omitempty"`
	AuthorizationMethod  *AuthorizationMethod           `json:"authorization_method,omitempty"`
	Charging             *ConversationCharging          `json:"charging,omitempty"`
	PhoneCall            *ConversationMetadataPhoneCall `json:"phone_call,omitempty"`
	BatchCall            any                            `json:"batch_call,omitempty"`
	TerminationReason    string                         `json:"termination_reason"`
	Error                any                            `json:"error,omitempty"`
	Warnings             []string                       `json:"warnings,omitempty"`
	MainLanguage         *string                        `json:"main_language,omitempty"`
	RagUsage             any                            `json:"rag_usage,omitempty"`
	TextOnly             bool                           `json:"text_only"`
	FeaturesUsage        *FeaturesUsageCommonModel      `json:"features_usage,omitempty"`
	InitiatorId          *string                        `json:"initiator_id,omitempty"`
}

type ConversationDeletionSettings struct {
	DeletionTimeUnixSecs            *int64 `json:"deletion_time_unix_secs,omitempty"`
	DeletedLogsAtTimeUnixSecs       *int64 `json:"deleted_logs_at_time_unix_secs,omitempty"`
	DeletedAudioAtTimeUnixSecs      *int64 `json:"deleted_audio_at_time_unix_secs,omitempty"`
	DeletedTranscriptAtTimeUnixSecs *int64 `json:"deleted_transcript_at_time_unix_secs,omitempty"`
	DeleteTranscriptAndPII          bool   `json:"delete_transcript_and_pii"`
	DeleteAudio                     bool   `json:"delete_audio"`
}

type ConversationFeedbackType string

const (
	ConversationFeedbackTypeThumbs ConversationFeedbackType = "thumbs"
	ConversationFeedbackTypeRating ConversationFeedbackType = "rating"
)

type ConversationFeedbackSummary struct {
	Type         *ConversationFeedbackType `json:"type,omitempty"`
	OverallScore *UserFeedbackScore        `json:"overall_score,omitempty"`
	Likes        int                       `json:"likes"`
	Dislikes     int                       `json:"dislikes"`
	Rating       *int                      `json:"rating,omitempty"`
	Comment      *string                   `json:"comment,omitempty"`
}

type AuthorizationMethod string

const (
	AuthorizationMethodInvalid             AuthorizationMethod = "invalid"
	AuthorizationMethodPublic              AuthorizationMethod = "public"
	AuthorizationMethodAuthorizationHeader AuthorizationMethod = "authorization_header"
	AuthorizationMethodSignedUrl           AuthorizationMethod = "signed_url"
	AuthorizationMethodShareableLink       AuthorizationMethod = "shareable_link"
	AuthorizationMethodLivekitToken        AuthorizationMethod = "livekit_token"
	AuthorizationMethodLivekitTokenWebsite AuthorizationMethod = "livekit_token_website"
	AuthorizationMethodGenesysApiKey       AuthorizationMethod = "genesys_api_key"
	AuthorizationMethodWhatsapp            AuthorizationMethod = "whatsapp"
)

// LLMCategoryUsage represents LLM usage for a specific category
type LLMCategoryUsage struct {
	IrreversibleGeneration *LLMUsage `json:"irreversible_generation,omitempty"`
	InitiatedGeneration    *LLMUsage `json:"initiated_generation,omitempty"`
}

// LLMUsage represents LLM model usage information
type LLMUsage struct {
	ModelUsage map[string]LLMInputOutputTokensUsage `json:"model_usage"`
}

// LLMInputOutputTokensUsage represents token usage for LLM input/output
type LLMInputOutputTokensUsage struct {
	Input           LLMTokensCategoryUsage `json:"input"`
	InputCacheRead  LLMTokensCategoryUsage `json:"input_cache_read"`
	InputCacheWrite LLMTokensCategoryUsage `json:"input_cache_write"`
	OutputTotal     LLMTokensCategoryUsage `json:"output_total"`
}

// LLMTokensCategoryUsage represents token usage for a specific category
type LLMTokensCategoryUsage struct {
	Tokens int     `json:"tokens"`
	Price  float64 `json:"price"`
}

// ConversationMetadataPhoneCall represents phone call information in conversation metadata
type ConversationMetadataPhoneCall struct {
	*ConversationHistoryTwilioPhoneCallModel
	*ConversationHistorySIPTrunkingPhoneCallModel
}

// ConversationHistoryTwilioPhoneCallModel represents Twilio phone call details
type ConversationHistoryTwilioPhoneCallModel struct {
	Direction      ConversationDirection `json:"direction"`
	PhoneNumberId  string                `json:"phone_number_id"`
	AgentNumber    string                `json:"agent_number"`
	ExternalNumber string                `json:"external_number"`
	Type           string                `json:"type"` // Always "twilio"
	StreamSid      string                `json:"stream_sid"`
	CallSid        string                `json:"call_sid"`
}

// ConversationHistorySIPTrunkingPhoneCallModel represents SIP trunking phone call details
type ConversationHistorySIPTrunkingPhoneCallModel struct {
	Direction      ConversationDirection `json:"direction"`
	PhoneNumberId  string                `json:"phone_number_id"`
	AgentNumber    string                `json:"agent_number"`
	ExternalNumber string                `json:"external_number"`
	Type           string                `json:"type"` // Always "sip_trunking"
	CallSid        string                `json:"call_sid"`
}

// FeaturesUsageCommonModel represents feature usage information
type FeaturesUsageCommonModel struct {
	LanguageDetection          *FeatureStatusCommonModel         `json:"language_detection,omitempty"`
	TransferToAgent            *FeatureStatusCommonModel         `json:"transfer_to_agent,omitempty"`
	TransferToNumber           *FeatureStatusCommonModel         `json:"transfer_to_number,omitempty"`
	Multivoice                 *FeatureStatusCommonModel         `json:"multivoice,omitempty"`
	DtmfTones                  *FeatureStatusCommonModel         `json:"dtmf_tones,omitempty"`
	ExternalMcpServers         *FeatureStatusCommonModel         `json:"external_mcp_servers,omitempty"`
	PiiZrmWorkspace            bool                              `json:"pii_zrm_workspace"`
	PiiZrmAgent                bool                              `json:"pii_zrm_agent"`
	ToolDynamicVariableUpdates *FeatureStatusCommonModel         `json:"tool_dynamic_variable_updates,omitempty"`
	IsLivekit                  bool                              `json:"is_livekit"`
	VoicemailDetection         *FeatureStatusCommonModel         `json:"voicemail_detection,omitempty"`
	Workflow                   *WorkflowFeaturesUsageCommonModel `json:"workflow,omitempty"`
	AgentTesting               *TestsFeatureUsageCommonModel     `json:"agent_testing,omitempty"`
	Versioning                 *FeatureStatusCommonModel         `json:"versioning,omitempty"`
}

// FeatureStatusCommonModel represents the status of a feature
type FeatureStatusCommonModel struct {
	Enabled bool `json:"enabled"`
	Used    bool `json:"used"`
}

// WorkflowFeaturesUsageCommonModel represents workflow feature usage
type WorkflowFeaturesUsageCommonModel struct {
	Enabled             bool                      `json:"enabled"`
	ToolNode            *FeatureStatusCommonModel `json:"tool_node,omitempty"`
	StandaloneAgentNode *FeatureStatusCommonModel `json:"standalone_agent_node,omitempty"`
	PhoneNumberNode     *FeatureStatusCommonModel `json:"phone_number_node,omitempty"`
	EndNode             *FeatureStatusCommonModel `json:"end_node,omitempty"`
}

// TestsFeatureUsageCommonModel represents test feature usage
type TestsFeatureUsageCommonModel struct {
	Enabled                       bool `json:"enabled"`
	TestsRanAfterLastModification bool `json:"tests_ran_after_last_modification"`
	TestsRanInLast7Days           bool `json:"tests_ran_in_last_7_days"`
}

type ConversationCharging struct {
	DevDiscount            bool              `json:"dev_discount"`
	IsBurst                bool              `json:"is_burst"`
	Tier                   *string           `json:"tier,omitempty"`
	LLMUsage               *LLMCategoryUsage `json:"llm_usage,omitempty"`
	LLMPrice               *float64          `json:"llm_price,omitempty"`
	LLMCharge              *int              `json:"llm_charge,omitempty"`
	CallCharge             *int              `json:"call_charge,omitempty"`
	FreeMinutesConsumed    float64           `json:"free_minutes_consumed"`
	FreeLLMDollarsConsumed float64           `json:"free_llm_dollars_consumed"`
}

type ConversationAnalysis struct {
	EvaluationCriteriaResults map[string]ConversationEvaluationCriteriaResult `json:"evaluation_criteria_results,omitempty"`
	DataCollectionResults     map[string]ConversationDataCollectionResult     `json:"data_collection_results,omitempty"`
	CallSuccessful            EvaluationSuccessResult                         `json:"call_successful"`
	TranscriptSummary         string                                          `json:"transcript_summary"`
	CallSummaryTitle          *string                                         `json:"call_summary_title,omitempty"`
}

type ConversationEvaluationCriteriaResult struct {
	CriteriaId string                  `json:"criteria_id"`
	Result     EvaluationSuccessResult `json:"result"`
	Rationale  string                  `json:"rationale"`
}

type ConversationDataCollectionResult struct {
	DataCollectionId string `json:"data_collection_id"`
	Value            any    `json:"value,omitempty"`
	Rationale        string `json:"rationale"`
}

type GetConversationResp struct {
	AgentId                          string                            `json:"agent_id"`
	ConversationId                   string                            `json:"conversation_id"`
	Status                           ConversationStatus                `json:"status"`
	UserId                           *string                           `json:"user_id,omitempty"`
	BranchId                         *string                           `json:"branch_id,omitempty"`
	Transcript                       []TranscriptMessage               `json:"transcript"`
	Metadata                         ConversationMetadata              `json:"metadata"`
	Analysis                         *ConversationAnalysis             `json:"analysis,omitempty"`
	ConversationInitiationClientData *ConversationInitiationClientData `json:"conversation_initiation_client_data,omitempty"`
	HasAudio                         bool                              `json:"has_audio"`
	HasUserAudio                     bool                              `json:"has_user_audio"`
	HasResponseAudio                 bool                              `json:"has_response_audio"`
}

// ConversationInitiationClientData represents client data for conversation initiation
type ConversationInitiationClientData struct {
	ConversationConfigOverride *ConversationConfigClientOverride `json:"conversation_config_override,omitempty"`
	CustomLlmExtraBody         map[string]any                    `json:"custom_llm_extra_body,omitempty"`
	UserId                     *string                           `json:"user_id,omitempty"`
	SourceInfo                 *ConversationInitiationSourceInfo `json:"source_info,omitempty"`
	DynamicVariables           map[string]any                    `json:"dynamic_variables,omitempty"` // Can be string, number, integer, or boolean
}

// ConversationConfigClientOverride represents client-side conversation config overrides
type ConversationConfigClientOverride struct {
	Turn         *TurnConfigOverride              `json:"turn,omitempty"`
	Tts          *TTSConversationalConfigOverride `json:"tts,omitempty"`
	Conversation *ConversationConfigOverride      `json:"conversation,omitempty"`
	Agent        *AgentConfigOverride             `json:"agent,omitempty"`
}

// TurnConfigOverride represents turn configuration override
type TurnConfigOverride struct {
	SoftTimeoutConfig *SoftTimeoutConfig `json:"soft_timeout_config,omitempty"`
}

// TTSConversationalConfigOverride represents TTS config override
type TTSConversationalConfigOverride struct {
	VoiceId         *string  `json:"voice_id,omitempty"`
	Stability       *float64 `json:"stability,omitempty"`
	Speed           *float64 `json:"speed,omitempty"`
	SimilarityBoost *float64 `json:"similarity_boost,omitempty"`
}

// ConversationConfigOverride represents conversation config override
type ConversationConfigOverride struct {
	TextOnly *bool `json:"text_only,omitempty"`
}

// AgentConfigOverride represents agent config override
type AgentConfigOverride struct {
	FirstMessage *string                      `json:"first_message,omitempty"`
	Language     *string                      `json:"language,omitempty"`
	Prompt       *PromptAgentAPIModelOverride `json:"prompt,omitempty"`
}

// PromptAgentAPIModelOverride represents prompt agent API model override
type PromptAgentAPIModelOverride struct {
	Prompt             *string  `json:"prompt,omitempty"`
	Llm                *string  `json:"llm,omitempty"`
	NativeMcpServerIds []string `json:"native_mcp_server_ids,omitempty"`
}

// ConversationInitiationSourceInfo represents source information for conversation initiation
type ConversationInitiationSourceInfo struct {
	Source  *ConversationInitiationSource `json:"source,omitempty"`
	Version *string                       `json:"version,omitempty"`
}

// ConversationInitiationSource represents the source of conversation initiation
type ConversationInitiationSource string

const (
	ConversationInitiationSourceUnknown        ConversationInitiationSource = "unknown"
	ConversationInitiationSourceAndroidSDK     ConversationInitiationSource = "android_sdk"
	ConversationInitiationSourceNodeJSSDK      ConversationInitiationSource = "node_js_sdk"
	ConversationInitiationSourceReactNativeSDK ConversationInitiationSource = "react_native_sdk"
	ConversationInitiationSourceReactSDK       ConversationInitiationSource = "react_sdk"
	ConversationInitiationSourceJSSDK          ConversationInitiationSource = "js_sdk"
	ConversationInitiationSourcePythonSDK      ConversationInitiationSource = "python_sdk"
	ConversationInitiationSourceWidget         ConversationInitiationSource = "widget"
	ConversationInitiationSourceSipTrunk       ConversationInitiationSource = "sip_trunk"
	ConversationInitiationSourceTwilio         ConversationInitiationSource = "twilio"
	ConversationInitiationSourceGenesys        ConversationInitiationSource = "genesys"
	ConversationInitiationSourceSwiftSDK       ConversationInitiationSource = "swift_sdk"
	ConversationInitiationSourceWhatsapp       ConversationInitiationSource = "whatsapp"
	ConversationInitiationSourceFlutterSDK     ConversationInitiationSource = "flutter_sdk"
)

// GetConversation gets the details of a particular conversation.
// https://elevenlabs.io/docs/api-reference/conversations/get
func (c *Client) GetConversation(ctx context.Context, req *GetConversationReq) (GetConversationResp, error) {
	if req == nil {
		return GetConversationResp{}, errors.New("request is nil")
	}

	body, err := c.get(ctx, "/convai/conversations/"+req.ConversationId)
	if err != nil {
		return GetConversationResp{}, err
	}

	var resp GetConversationResp
	if err := c.parseResponse(body, &resp); err != nil {
		return GetConversationResp{}, err
	}

	return resp, nil
}

// GetConversationAudioReq represents the request for getting conversation audio.
type GetConversationAudioReq struct {
	ConversationId string `path:"conversation_id"`
}

func NewGetConversationAudioReq(conversationId string) *GetConversationAudioReq {
	return &GetConversationAudioReq{
		ConversationId: conversationId,
	}
}

// GetConversationAudio gets the audio recording of a particular conversation.
// https://elevenlabs.io/docs/api-reference/conversations/get-audio
func (c *Client) GetConversationAudio(ctx context.Context, req *GetConversationAudioReq) ([]byte, error) {
	if req == nil {
		return nil, errors.New("request is nil")
	}

	body, err := c.get(ctx, "/convai/conversations/"+req.ConversationId+"/audio")
	if err != nil {
		return nil, err
	}

	return body, nil
}

// DeleteConversationReq represents the request for deleting a conversation.
type DeleteConversationReq struct {
	ConversationId string `path:"conversation_id"`
}

func NewDeleteConversationReq(conversationId string) *DeleteConversationReq {
	return &DeleteConversationReq{
		ConversationId: conversationId,
	}
}

// DeleteConversation deletes a particular conversation.
// https://elevenlabs.io/docs/api-reference/conversations/delete
func (c *Client) DeleteConversation(ctx context.Context, req *DeleteConversationReq) error {
	if req == nil {
		return errors.New("request is nil")
	}

	if err := c.delete(ctx, "/convai/conversations/"+req.ConversationId); err != nil {
		return err
	}

	return nil
}

// GetConversationTokenReq represents the request for getting a conversation token.
type GetConversationTokenReq struct {
	AgentId         string  `url:"agent_id"`
	ParticipantName *string `url:"participant_name,omitempty"`
}

func NewGetConversationTokenReq(agentId string) *GetConversationTokenReq {
	return &GetConversationTokenReq{
		AgentId: agentId,
	}
}

func (r GetConversationTokenReq) QueryString() string {
	v, err := query.Values(r)
	if err != nil {
		return ""
	}
	return "?" + v.Encode()
}

type GetConversationTokenResp struct {
	Token string `json:"token"`
}

// GetConversationToken gets a WebRTC session token for real-time communication.
// https://elevenlabs.io/docs/api-reference/conversations/get-webrtc-token
func (c *Client) GetConversationToken(ctx context.Context, req *GetConversationTokenReq) (GetConversationTokenResp, error) {
	if req == nil {
		return GetConversationTokenResp{}, errors.New("request is nil")
	}

	body, err := c.get(ctx, "/convai/conversation/token"+req.QueryString())
	if err != nil {
		return GetConversationTokenResp{}, err
	}

	var resp GetConversationTokenResp
	if err := c.parseResponse(body, &resp); err != nil {
		return GetConversationTokenResp{}, err
	}

	return resp, nil
}

// GetSignedUrlReq represents the request for getting a signed URL.
type GetSignedUrlReq struct {
	AgentId               string `url:"agent_id"`
	IncludeConversationId *bool  `url:"include_conversation_id,omitempty"`
}

func NewGetSignedUrlReq(agentId string) *GetSignedUrlReq {
	return &GetSignedUrlReq{
		AgentId: agentId,
	}
}

func (r GetSignedUrlReq) QueryString() string {
	v, err := query.Values(r)
	if err != nil {
		return ""
	}
	return "?" + v.Encode()
}

type GetSignedUrlResp struct {
	SignedUrl string `json:"signed_url"`
}

// GetSignedUrl gets a signed URL to start a conversation with an agent that requires authorization.
// https://elevenlabs.io/docs/api-reference/conversations/get-signed-url
func (c *Client) GetSignedUrl(ctx context.Context, req *GetSignedUrlReq) (GetSignedUrlResp, error) {
	if req == nil {
		return GetSignedUrlResp{}, errors.New("request is nil")
	}

	body, err := c.get(ctx, "/convai/conversation/get-signed-url"+req.QueryString())
	if err != nil {
		return GetSignedUrlResp{}, err
	}

	var resp GetSignedUrlResp
	if err := c.parseResponse(body, &resp); err != nil {
		return GetSignedUrlResp{}, err
	}

	return resp, nil
}

// SendConversationFeedbackReq represents the request for sending conversation feedback.
type SendConversationFeedbackReq struct {
	ConversationId string             `path:"conversation_id"`
	Feedback       *UserFeedbackScore `json:"feedback,omitempty"`
}

func NewSendConversationFeedbackReq(conversationId string, feedback *UserFeedbackScore) *SendConversationFeedbackReq {
	return &SendConversationFeedbackReq{
		ConversationId: conversationId,
		Feedback:       feedback,
	}
}

// SendConversationFeedback sends feedback for the given conversation.
// https://elevenlabs.io/docs/api-reference/conversations/create
func (c *Client) SendConversationFeedback(ctx context.Context, req *SendConversationFeedbackReq) error {
	if req == nil {
		return errors.New("request is nil")
	}

	_, err := c.post(ctx, "/convai/conversations/"+req.ConversationId+"/feedback", req)
	if err != nil {
		return err
	}

	return nil
}
