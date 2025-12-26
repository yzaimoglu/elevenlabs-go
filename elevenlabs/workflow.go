package elevenlabs

// AgentWorkflow represents a workflow for an agent, defining the flow of the conversation
// and how the agent interacts with tools.
type AgentWorkflow struct {
	Edges                map[string]WorkflowEdge `json:"edges,omitempty"`
	Nodes                map[string]WorkflowNode `json:"nodes,omitempty"`
	PreventSubagentLoops *bool                   `json:"prevent_subagent_loops,omitempty"`
}

// WorkflowEdge represents a connection between two nodes in the workflow.
type WorkflowEdge struct {
	Source            string             `json:"source"`
	Target            string             `json:"target"`
	ForwardCondition  *WorkflowCondition `json:"forward_condition,omitempty"`
	BackwardCondition *WorkflowCondition `json:"backward_condition,omitempty"`
}

// WorkflowCondition represents a condition for edge traversal.
// It can be unconditional, LLM-based, result-based, or expression-based.
type WorkflowCondition struct {
	Label      *string                      `json:"label,omitempty"`
	Type       WorkflowConditionType        `json:"type"`
	Condition  *string                      `json:"condition,omitempty"`  // For LLM type
	Successful *bool                        `json:"successful,omitempty"` // For result type
	Expression *WorkflowConditionExpression `json:"expression,omitempty"` // For expression type
}

type WorkflowConditionType string

const (
	WorkflowConditionTypeUnconditional WorkflowConditionType = "unconditional"
	WorkflowConditionTypeLLM           WorkflowConditionType = "llm"
	WorkflowConditionTypeResult        WorkflowConditionType = "result"
	WorkflowConditionTypeExpression    WorkflowConditionType = "expression"
)

// WorkflowConditionExpression represents an AST expression for workflow conditions.
// This is a simplified representation - the actual AST can be deeply nested.
type WorkflowConditionExpression struct {
	Type     WorkflowExpressionNodeType `json:"type"`
	Value    any                        `json:"value,omitempty"`    // For literals
	Name     *string                    `json:"name,omitempty"`     // For dynamic_variable
	Prompt   *string                    `json:"prompt,omitempty"`   // For llm
	Left     *WorkflowConditionExpression  `json:"left,omitempty"`     // For binary operators
	Right    *WorkflowConditionExpression  `json:"right,omitempty"`    // For binary operators
	Children []WorkflowConditionExpression `json:"children,omitempty"` // For and_operator, or_operator
}

type WorkflowExpressionNodeType string

const (
	WorkflowExpressionNodeTypeStringLiteral   WorkflowExpressionNodeType = "string_literal"
	WorkflowExpressionNodeTypeNumberLiteral   WorkflowExpressionNodeType = "number_literal"
	WorkflowExpressionNodeTypeBooleanLiteral  WorkflowExpressionNodeType = "boolean_literal"
	WorkflowExpressionNodeTypeLLM             WorkflowExpressionNodeType = "llm"
	WorkflowExpressionNodeTypeDynamicVariable WorkflowExpressionNodeType = "dynamic_variable"
	WorkflowExpressionNodeTypeOrOperator      WorkflowExpressionNodeType = "or_operator"
	WorkflowExpressionNodeTypeAndOperator     WorkflowExpressionNodeType = "and_operator"
	WorkflowExpressionNodeTypeEqOperator      WorkflowExpressionNodeType = "eq_operator"
	WorkflowExpressionNodeTypeNeqOperator     WorkflowExpressionNodeType = "neq_operator"
	WorkflowExpressionNodeTypeGtOperator      WorkflowExpressionNodeType = "gt_operator"
	WorkflowExpressionNodeTypeLtOperator      WorkflowExpressionNodeType = "lt_operator"
	WorkflowExpressionNodeTypeGteOperator     WorkflowExpressionNodeType = "gte_operator"
	WorkflowExpressionNodeTypeLteOperator     WorkflowExpressionNodeType = "lte_operator"
)

// WorkflowNode represents a node in the workflow.
// The Type field determines which other fields are relevant.
type WorkflowNode struct {
	Type      WorkflowNodeType `json:"type"`
	Position  *WorkflowPosition `json:"position,omitempty"`
	EdgeOrder []string          `json:"edge_order,omitempty"`

	// For phone_number type
	TransferDestination *WorkflowTransferDestination `json:"transfer_destination,omitempty"`
	TransferType        *TransferToNumberType        `json:"transfer_type,omitempty"`

	// For override_agent type
	ConversationConfig      *ConversationConfigWorkflowOverride `json:"conversation_config,omitempty"`
	AdditionalPrompt        *string                             `json:"additional_prompt,omitempty"`
	AdditionalKnowledgeBase []AgentPromptKnowledgeBase          `json:"additional_knowledge_base,omitempty"`
	AdditionalToolIds       []string                            `json:"additional_tool_ids,omitempty"`
	Label                   *string                             `json:"label,omitempty"`

	// For standalone_agent type
	AgentId                            *string `json:"agent_id,omitempty"`
	DelayMs                            *int    `json:"delay_ms,omitempty"`
	TransferMessage                    *string `json:"transfer_message,omitempty"`
	EnableTransferredAgentFirstMessage *bool   `json:"enable_transferred_agent_first_message,omitempty"`

	// For tool type
	Tools []WorkflowToolLocator `json:"tools,omitempty"`
}

type WorkflowNodeType string

const (
	WorkflowNodeTypeStart           WorkflowNodeType = "start"
	WorkflowNodeTypeEnd             WorkflowNodeType = "end"
	WorkflowNodeTypePhoneNumber     WorkflowNodeType = "phone_number"
	WorkflowNodeTypeOverrideAgent   WorkflowNodeType = "override_agent"
	WorkflowNodeTypeStandaloneAgent WorkflowNodeType = "standalone_agent"
	WorkflowNodeTypeTool            WorkflowNodeType = "tool"
)

type WorkflowPosition struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

// WorkflowTransferDestination represents a transfer destination in a workflow node.
type WorkflowTransferDestination struct {
	Type        WorkflowTransferDestinationType `json:"type"`
	PhoneNumber *string                         `json:"phone_number,omitempty"`
	SIPUri      *string                         `json:"sip_uri,omitempty"`
}

type WorkflowTransferDestinationType string

const (
	WorkflowTransferDestinationTypePhone                WorkflowTransferDestinationType = "phone"
	WorkflowTransferDestinationTypeSIPUri               WorkflowTransferDestinationType = "sip_uri"
	WorkflowTransferDestinationTypePhoneDynamicVariable WorkflowTransferDestinationType = "phone_dynamic_variable"
	WorkflowTransferDestinationTypeSIPUriDynamicVariable WorkflowTransferDestinationType = "sip_uri_dynamic_variable"
)

type WorkflowToolLocator struct {
	ToolId string `json:"tool_id"`
}

// ConversationConfigWorkflowOverride represents configuration overrides for a workflow subagent.
type ConversationConfigWorkflowOverride struct {
	ASR             *ASRConfigWorkflowOverride   `json:"asr,omitempty"`
	Turn            *TurnConfigWorkflowOverride  `json:"turn,omitempty"`
	TTS             *TTSConfigWorkflowOverride   `json:"tts,omitempty"`
	Conversation    *ConversationWorkflowOverride `json:"conversation,omitempty"`
	LanguagePresets map[string]ConversationConfigLanguagePreset `json:"language_presets,omitempty"`
	Agent           *AgentConfigWorkflowOverride `json:"agent,omitempty"`
}

type ASRConfigWorkflowOverride struct {
	Quality              *ASRQuality  `json:"quality,omitempty"`
	Provider             *ASRProvider `json:"provider,omitempty"`
	UserInputAudioFormat *AudioFormat `json:"user_input_audio_format,omitempty"`
	Keywords             []string     `json:"keywords,omitempty"`
}

type TurnConfigWorkflowOverride struct {
	TurnTimeout           *float32                        `json:"turn_timeout,omitempty"`
	InitialWaitTime       *float32                        `json:"initial_wait_time,omitempty"`
	SilenceEndCallTimeout *float32                        `json:"silence_end_call_timeout,omitempty"`
	SoftTimeoutConfig     *SoftTimeoutConfigOverride      `json:"soft_timeout_config,omitempty"`
	TurnEagerness         *TurnEagerness                  `json:"turn_eagerness,omitempty"`
}

type SoftTimeoutConfigOverride struct {
	TimeoutSeconds *float64 `json:"timeout_seconds,omitempty"`
	Message        *string  `json:"message,omitempty"`
}

type TurnEagerness string

const (
	TurnEagernessPatient TurnEagerness = "patient"
	TurnEagernessNormal  TurnEagerness = "normal"
	TurnEagernessEager   TurnEagerness = "eager"
)

type TTSConfigWorkflowOverride struct {
	ModelId                          *TTSModelId                       `json:"model_id,omitempty"`
	VoiceId                          *string                           `json:"voice_id,omitempty"`
	SupportedVoices                  []SupportedVoice                  `json:"supported_voices,omitempty"`
	AgentOutputAudioFormat           *AudioFormat                      `json:"agent_output_audio_format,omitempty"`
	OptimizeStreamingLatency         *OptimizeStreamingLatency         `json:"optimize_streaming_latency,omitempty"`
	Stability                        *float32                          `json:"stability,omitempty"`
	Speed                            *float32                          `json:"speed,omitempty"`
	SimilarityBoost                  *float32                          `json:"similarity_boost,omitempty"`
	TextNormalisationType            *TextNormalisationType            `json:"text_normalisation_type,omitempty"`
	PronounciationDictionaryLocators []PronounciationDictionaryLocator `json:"pronunciation_dictionary_locators,omitempty"`
}

type TextNormalisationType string

const (
	TextNormalisationTypeSystemPrompt TextNormalisationType = "system_prompt"
	TextNormalisationTypeElevenLabs   TextNormalisationType = "elevenlabs"
)

type ConversationWorkflowOverride struct {
	TextOnly           *bool         `json:"text_only,omitempty"`
	MaxDurationSeconds *int          `json:"max_duration_seconds,omitempty"`
	ClientEvents       []ClientEvent `json:"client_events,omitempty"`
	MonitoringEnabled  *bool         `json:"monitoring_enabled,omitempty"`
	MonitoringEvents   []ClientEvent `json:"monitoring_events,omitempty"`
}

type AgentConfigWorkflowOverride struct {
	FirstMessage                    *string                           `json:"first_message,omitempty"`
	Language                        *string                           `json:"language,omitempty"`
	HinglishMode                    *bool                             `json:"hinglish_mode,omitempty"`
	DynamicVariables                *AgentDynamicVariables            `json:"dynamic_variables,omitempty"`
	DisableFirstMessageInterruptions *bool                            `json:"disable_first_message_interruptions,omitempty"`
	Prompt                          *AgentPromptWorkflowOverride      `json:"prompt,omitempty"`
}

type AgentPromptWorkflowOverride struct {
	Prompt                   *string                    `json:"prompt,omitempty"`
	LLM                      *AgentPromptLLM            `json:"llm,omitempty"`
	ReasoningEffort          *LLMReasoningEffort        `json:"reasoning_effort,omitempty"`
	ThinkingBudget           *int                       `json:"thinking_budget,omitempty"`
	Temperature              *float32                   `json:"temperature,omitempty"`
	MaxTokens                *int                       `json:"max_tokens,omitempty"`
	ToolIds                  []string                   `json:"tool_ids,omitempty"`
	BuiltInTools             *AgentPromptBuiltInTools   `json:"built_in_tools,omitempty"`
	MCPServerIds             []string                   `json:"mcp_server_ids,omitempty"`
	NativeMCPServerIds       []string                   `json:"native_mcp_server_ids,omitempty"`
	KnowledgeBase            []AgentPromptKnowledgeBase `json:"knowledge_base,omitempty"`
	CustomLLM                *AgentPromptCustomLLM      `json:"custom_llm,omitempty"`
	IgnoreDefaultPersonality *bool                      `json:"ignore_default_personality,omitempty"`
	RAG                      *AgentPromptRAG            `json:"rag,omitempty"`
	Timezone                 *string                    `json:"timezone,omitempty"`
	BackupLLMConfig          *BackupLLMConfig           `json:"backup_llm_config,omitempty"`
	Tools                    []AgentTool                `json:"tools,omitempty"`
}

type LLMReasoningEffort string

const (
	LLMReasoningEffortNone    LLMReasoningEffort = "none"
	LLMReasoningEffortMinimal LLMReasoningEffort = "minimal"
	LLMReasoningEffortLow     LLMReasoningEffort = "low"
	LLMReasoningEffortMedium  LLMReasoningEffort = "medium"
	LLMReasoningEffortHigh    LLMReasoningEffort = "high"
)

type BackupLLMConfig struct {
	Preference BackupLLMPreference `json:"preference"`
	Order      []AgentPromptLLM    `json:"order,omitempty"` // For override preference
}

type BackupLLMPreference string

const (
	BackupLLMPreferenceDefault  BackupLLMPreference = "default"
	BackupLLMPreferenceDisabled BackupLLMPreference = "disabled"
	BackupLLMPreferenceOverride BackupLLMPreference = "override"
)

// AgentTool represents a tool configuration that can be used by an agent.
type AgentTool struct {
	Type                  AgentToolType                 `json:"type"`
	Name                  string                        `json:"name"`
	Description           string                        `json:"description"`
	ResponseTimeoutSecs   *int                          `json:"response_timeout_secs,omitempty"`
	DisableInterruptions  *bool                         `json:"disable_interruptions,omitempty"`
	ForcePreToolSpeech    *bool                         `json:"force_pre_tool_speech,omitempty"`
	Assignments           []AgentPromptBuiltInToolAssignment `json:"assignments,omitempty"`
	ToolCallSound         *ToolCallSoundType            `json:"tool_call_sound,omitempty"`
	ToolCallSoundBehavior *ToolCallSoundBehavior        `json:"tool_call_sound_behavior,omitempty"`
	DynamicVariables      *AgentDynamicVariables        `json:"dynamic_variables,omitempty"`
	ExecutionMode         *ToolExecutionMode            `json:"execution_mode,omitempty"`

	// For webhook type
	APISchema *WebhookAPISchema `json:"api_schema,omitempty"`

	// For client type
	Parameters      *ClientToolParameters `json:"parameters,omitempty"`
	ExpectsResponse *bool                 `json:"expects_response,omitempty"`

	// For system type
	Params *AgentPromptBuiltInToolParams `json:"params,omitempty"`

	// For api_integration_webhook type
	ToolVersion               *string                       `json:"tool_version,omitempty"`
	APIIntegrationId          *string                       `json:"api_integration_id,omitempty"`
	APIIntegrationConnectionId *string                       `json:"api_integration_connection_id,omitempty"`
	APISchemaOverrides        *APIIntegrationWebhookOverrides `json:"api_schema_overrides,omitempty"`
}

type AgentToolType string

const (
	AgentToolTypeWebhook              AgentToolType = "webhook"
	AgentToolTypeClient               AgentToolType = "client"
	AgentToolTypeSystem               AgentToolType = "system"
	AgentToolTypeAPIIntegrationWebhook AgentToolType = "api_integration_webhook"
)

type ToolCallSoundType string

const (
	ToolCallSoundTypeTyping    ToolCallSoundType = "typing"
	ToolCallSoundTypeElevator1 ToolCallSoundType = "elevator1"
	ToolCallSoundTypeElevator2 ToolCallSoundType = "elevator2"
	ToolCallSoundTypeElevator3 ToolCallSoundType = "elevator3"
	ToolCallSoundTypeElevator4 ToolCallSoundType = "elevator4"
)

type ToolCallSoundBehavior string

const (
	ToolCallSoundBehaviorAuto   ToolCallSoundBehavior = "auto"
	ToolCallSoundBehaviorAlways ToolCallSoundBehavior = "always"
)

type ToolExecutionMode string

const (
	ToolExecutionModeImmediate      ToolExecutionMode = "immediate"
	ToolExecutionModePostToolSpeech ToolExecutionMode = "post_tool_speech"
	ToolExecutionModeAsync          ToolExecutionMode = "async"
)

type WebhookAPISchema struct {
	URL               string                 `json:"url"`
	Method            *WebhookMethod         `json:"method,omitempty"`
	RequestHeaders    map[string]any         `json:"request_headers,omitempty"`
	PathParamsSchema  map[string]any         `json:"path_params_schema,omitempty"`
	QueryParamsSchema *QueryParamsSchema     `json:"query_params_schema,omitempty"`
	RequestBodySchema *ObjectJsonSchemaProperty `json:"request_body_schema,omitempty"`
	ContentType       *WebhookContentType    `json:"content_type,omitempty"`
	AuthConnection    *AuthConnectionLocator `json:"auth_connection,omitempty"`
}

type WebhookMethod string

const (
	WebhookMethodGET    WebhookMethod = "GET"
	WebhookMethodPOST   WebhookMethod = "POST"
	WebhookMethodPUT    WebhookMethod = "PUT"
	WebhookMethodPATCH  WebhookMethod = "PATCH"
	WebhookMethodDELETE WebhookMethod = "DELETE"
)

type WebhookContentType string

const (
	WebhookContentTypeJSON           WebhookContentType = "application/json"
	WebhookContentTypeFormURLEncoded WebhookContentType = "application/x-www-form-urlencoded"
)

type QueryParamsSchema struct {
	Properties map[string]LiteralJsonSchemaProperty `json:"properties"`
	Required   []string                             `json:"required,omitempty"`
}

type LiteralJsonSchemaProperty struct {
	Type             LiteralJsonSchemaPropertyType `json:"type"`
	Description      *string                       `json:"description,omitempty"`
	Enum             []string                      `json:"enum,omitempty"`
	IsSystemProvided *bool                         `json:"is_system_provided,omitempty"`
	DynamicVariable  *string                       `json:"dynamic_variable,omitempty"`
	ConstantValue    any                           `json:"constant_value,omitempty"`
}

type LiteralJsonSchemaPropertyType string

const (
	LiteralJsonSchemaPropertyTypeBoolean LiteralJsonSchemaPropertyType = "boolean"
	LiteralJsonSchemaPropertyTypeString  LiteralJsonSchemaPropertyType = "string"
	LiteralJsonSchemaPropertyTypeInteger LiteralJsonSchemaPropertyType = "integer"
	LiteralJsonSchemaPropertyTypeNumber  LiteralJsonSchemaPropertyType = "number"
)

type ObjectJsonSchemaProperty struct {
	Type        string                           `json:"type"` // "object" or "array"
	Description *string                          `json:"description,omitempty"`
	Properties  map[string]any                   `json:"properties,omitempty"`
	Required    []string                         `json:"required,omitempty"`
	Items       any                              `json:"items,omitempty"` // For array type
}

type AuthConnectionLocator struct {
	AuthConnectionId string `json:"auth_connection_id"`
}

type ClientToolParameters struct {
	Type        string         `json:"type"` // "object"
	Description *string        `json:"description,omitempty"`
	Properties  map[string]any `json:"properties,omitempty"`
	Required    []string       `json:"required,omitempty"`
}

type APIIntegrationWebhookOverrides struct {
	PathParamsSchema    map[string]LiteralOverride `json:"path_params_schema,omitempty"`
	QueryParamsSchema   *QueryOverride             `json:"query_params_schema,omitempty"`
	RequestBodySchema   *ObjectOverride            `json:"request_body_schema,omitempty"`
	RequestHeaders      map[string]any             `json:"request_headers,omitempty"`
}

type LiteralOverride struct {
	Description     *string `json:"description,omitempty"`
	DynamicVariable *string `json:"dynamic_variable,omitempty"`
	ConstantValue   any     `json:"constant_value,omitempty"`
}

type QueryOverride struct {
	Properties map[string]LiteralOverride `json:"properties,omitempty"`
	Required   []string                   `json:"required,omitempty"`
}

type ObjectOverride struct {
	Description *string        `json:"description,omitempty"`
	Properties  map[string]any `json:"properties,omitempty"`
	Required    []string       `json:"required,omitempty"`
}
