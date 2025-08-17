package elevenlabs

type ConversationConfig struct {
	ASR             *ConversationConfigASR                      `json:"asr,omitempty"`
	Turn            *ConversationConfigTurn                     `json:"turn,omitempty"`
	TTS             *ConversationConfigTTS                      `json:"tts,omitempty"`
	Conversation    *ConversationConfigConversation             `json:"conversation,omitempty"`
	LanguagePresets map[string]ConversationConfigLanguagePreset `json:"language_presets,omitempty"`
	Agent           *ConversationConfigAgent                    `json:"agent,omitempty"`
}

type ConversationConfigAgent struct {
	FirstMessage     *string                `json:"first_message,omitempty"`
	Language         *string                `json:"language,omitempty"`
	DynamicVariables *AgentDynamicVariables `json:"dynamic_variables,omitempty"`
	Prompt           *AgentPrompt           `json:"prompt,omitempty"`
}

type AgentPrompt struct {
	Prompt                   *string                    `json:"prompt"`
	LLM                      *AgentPromptLLM            `json:"llm,omitempty"`
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
}

type AgentPromptRAG struct {
	Enabled                    *bool              `json:"enabled,omitempty"`
	EmbeddingModel             *RAGEmbeddingModel `json:"embedding_model,omitempty"`
	MaxVectorDistance          *float32           `json:"max_vector_distance,omitempty"`
	MaxDocumentsLength         *int               `json:"max_documents_length,omitempty"`
	MaxRetrievedRAGChunksCount *int               `json:"max_retrieved_rag_chunks_count,omitempty"`
}

type RAGEmbeddingModel string

const (
	RAGEmbeddingModelE5Mistral      RAGEmbeddingModel = "e5_mistral_7b_instruct"
	RAGEmbeddingModelMultilingualE5 RAGEmbeddingModel = "multilingual_e5_large_instruct"
)

type AgentPromptCustomLLM struct {
	URL            string         `json:"url"`
	ModelID        *string        `json:"model_id,omitempty"`
	APIKey         *APIKey        `json:"api_key"`
	RequestHeaders map[string]any `json:"request_headers,omitempty"`
	APIVersion     *string        `json:"api_version,omitempty"`
}

type APIKey struct {
	SecretID string `json:"secret_id"`
}

type AgentPromptKnowledgeBase struct {
	Type      AgentPromptKnowledgeBaseType       `json:"type"`
	Name      string                             `json:"name"`
	ID        string                             `json:"id"`
	UsageMode *AgentPromptKnowledgeBaseUsageMode `json:"usage_mode,omitempty"`
}

type AgentPromptKnowledgeBaseType string

const (
	KnowledgeBaseTypeFile AgentPromptKnowledgeBaseType = "file"
	KnowledgeBaseTypeUrl  AgentPromptKnowledgeBaseType = "url"
	KnowledgeBaseTypeText AgentPromptKnowledgeBaseType = "text"
)

type AgentPromptKnowledgeBaseUsageMode string

const (
	KnowledgeBaseUsageModePrompt AgentPromptKnowledgeBaseUsageMode = "prompt"
	KnowledgeBaseUsageModeAuto   AgentPromptKnowledgeBaseUsageMode = "auto"
)

type AgentPromptBuiltInTools struct {
	EndCall             *AgentPromptBuiltInTool `json:"end_call,omitempty"`
	LanguageDetection   *AgentPromptBuiltInTool `json:"language_detection,omitempty"`
	TransferToAgent     *AgentPromptBuiltInTool `json:"transfer_to_agent,omitempty"`
	TransferToNumber    *AgentPromptBuiltInTool `json:"transfer_to_number,omitempty"`
	SkipTurn            *AgentPromptBuiltInTool `json:"skip_turn,omitempty"`
	PlayKeypadTouchTone *AgentPromptBuiltInTool `json:"play_keypad_touch_tone,omitempty"`
	VoicemailDetection  *AgentPromptBuiltInTool `json:"voicemail_detection,omitempty"`
}

type AgentPromptBuiltInTool struct {
	Name                 string                             `json:"name"`
	Description          string                             `json:"description"`
	Params               AgentPromptBuiltInToolParams       `json:"params"`
	ResponseTimeoutSecs  *int                               `json:"response_timeout_secs,omitempty"`
	DisableInterruptions *bool                              `json:"disable_interruptions,omitempty"`
	ForcePreToolSpeech   *bool                              `json:"force_pre_tool_speech,omitempty"`
	Assignments          []AgentPromptBuiltInToolAssignment `json:"assignments,omitempty"`
	Type                 *AgentPromptBuiltInToolType        `json:"type,omitempty"`
}

type AgentPromptBuiltInToolParams struct {
	EndCallToolConfig            BuiltInToolParamsEndCallToolConfig            `json:"EndCallToolConfig"`
	LanguageDetectionToolConfig  BuiltInToolParamsLanguageDetectionToolConfig  `json:"LanguageDetectionToolConfig"`
	TransferToAgentToolConfig    BuiltInToolParamsTransferToAgentToolConfig    `json:"TransferToAgentToolConfig"`
	TransferToNumberToolConfig   BuiltInToolParamsTransferToNumberToolConfig   `json:"TransferToNumberToolConfig"`
	SkipTurnToolConfig           BuiltInToolParamsSkipTurnToolConfig           `json:"SkipTurnToolConfig"`
	PlayDTMFToolConfig           BuiltInToolParamsPlayDTMFToolConfig           `json:"PlayDTMFToolConfig"`
	VoicemailDetectionToolConfig BuiltInToolParamsVoicemailDetectionToolConfig `json:"VoicemailDetectionToolConfig"`
}

type BuiltInToolParamsSkipTurnToolConfig struct {
	SystemToolType *PromptBuiltInSystemToolType `json:"system_tool_type,omitempty"`
}

type BuiltInToolParamsPlayDTMFToolConfig struct {
	SystemToolType *PromptBuiltInSystemToolType `json:"system_tool_type,omitempty"`
}

type BuiltInToolParamsVoicemailDetectionToolConfig struct {
	SystemToolType   *PromptBuiltInSystemToolType `json:"system_tool_type,omitempty"`
	VoicemailMessage *string                      `json:"voicemail_message,omitempty"`
}

type BuiltInToolParamsEndCallToolConfig struct {
	SystemToolType *PromptBuiltInSystemToolType `json:"system_tool_type,omitempty"`
}

type BuiltInToolParamsLanguageDetectionToolConfig struct {
	SystemToolType *PromptBuiltInSystemToolType `json:"system_tool_type,omitempty"`
}

type BuiltInToolParamsTransferToAgentToolConfig struct {
	SystemToolType *PromptBuiltInSystemToolType               `json:"system_tool_type,omitempty"`
	Transfers      []PromptBuiltInToolTransferToAgentTransfer `json:"transfers"`
}

type BuiltInToolParamsTransferToNumberToolConfig struct {
	SystemToolType      *PromptBuiltInSystemToolType                `json:"system_tool_type,omitempty"`
	EnableClientMessage *bool                                       `json:"enable_client_message,omitempty"`
	Transfers           []PromptBuiltInToolTransferToNumberTransfer `json:"transfers,omitempty"`
}

type PromptBuiltInToolTransferToNumberTransfer struct {
	Condition           string                       `json:"condition"`
	TransferDestination *TransferToNumberDestination `json:"transfer_destination"`
	TransferType        *TransferToNumberType        `json:"transfer_type"`
}

type TransferToNumberDestination struct {
	PhoneNumberTransferDestination *PhoneNumberTransferDestination `json:"PhoneNumberTransferDestination"`
	SIPUriTransferDestination      *SIPUriTransferDestination      `json:"SIPUriTransferDestination"`
}

type SIPUriTransferDestination struct {
	SIPUri string                         `json:"sip_uri"`
	Type   *SIPUriTransferDestinationType `json:"type"`
}

type SIPUriTransferDestinationType string

const (
	SIPUriTransferDestinationTypeSIPUri SIPUriTransferDestinationType = "sip_uri"
)

type PhoneNumberTransferDestination struct {
	PhoneNumber string                              `json:"phone_number"`
	Type        *PhoneNumberTransferDestinationType `json:"type"`
}

type PhoneNumberTransferDestinationType string

const (
	PhoneNumberTransferDestinationTypePhone PhoneNumberTransferDestinationType = "phone"
)

type TransferToNumberType string

const (
	TransferToNumberTypeConference TransferToNumberType = "conference"
	TransferToNumberTypeSipRefer   TransferToNumberType = "sip_refer"
)

type PromptBuiltInToolTransferToAgentTransfer struct {
	AgentID                            string  `json:"agent_id"`
	Condition                          string  `json:"condition"`
	DelayMs                            *int    `json:"delay_ms,omitempty"`
	TransferMessage                    *string `json:"transfer_message"`
	EnableTransferredAgentFirstMessage *bool   `json:"enable_transferred_agent_first_message,omitempty"`
}

type PromptBuiltInSystemToolType string

const (
	SystemToolTypeEndCall            PromptBuiltInSystemToolType = "end_call"
	SystemToolTypeLanguageDetection  PromptBuiltInSystemToolType = "language_detection"
	SystemToolTypeTransferToAgent    PromptBuiltInSystemToolType = "transfer_to_agent"
	SystemToolTypeTransferToNumber   PromptBuiltInSystemToolType = "transfer_to_number"
	SystemToolTypeSkipTurn           PromptBuiltInSystemToolType = "skip_turn"
	SystemToolTypePlayDTMF           PromptBuiltInSystemToolType = "play_dtmf"
	SystemToolTypeVoicemailDetection PromptBuiltInSystemToolType = "voicemail_detection"
)

type AgentPromptBuiltInToolAssignment struct {
	DynamicVariable string                                                    `json:"dynamic_variable"`
	ValuePath       string                                                    `json:"value_path"`
	Source          *ConversationConfigAgentPromptBuiltInToolAssignmentSource `json:"source"`
}

type ConversationConfigAgentPromptBuiltInToolAssignmentSource string

const (
	SourceTypeResponse ConversationConfigAgentPromptBuiltInToolAssignmentSource = "response"
)

type AgentPromptBuiltInToolType string

const (
	ToolTypeSystem AgentPromptBuiltInToolType = "system"
)

type AgentPromptLLM string

const (
	GPT4oMini         AgentPromptLLM = "gpt-4o-mini"
	GPT4o             AgentPromptLLM = "gpt-4o"
	GPT4              AgentPromptLLM = "gpt-4"
	GPT4Turbo         AgentPromptLLM = "gpt-4-turbo"
	GPT41             AgentPromptLLM = "gpt-4.1"
	GPT41Mini         AgentPromptLLM = "gpt-4.1-mini"
	GPT41Nano         AgentPromptLLM = "gpt-4.1-nano"
	GPT5              AgentPromptLLM = "gpt-5"
	GPT5Mini          AgentPromptLLM = "gpt-5-mini"
	GPT5Nano          AgentPromptLLM = "gpt-5-nano"
	GPT35Turbo        AgentPromptLLM = "gpt-3.5-turbo"
	Gemini15Pro       AgentPromptLLM = "gemini-1.5-pro"
	Gemini15Flash     AgentPromptLLM = "gemini-1.5-flash"
	Gemini20Flash     AgentPromptLLM = "gemini-2.0-flash"
	Gemini20FlashLite AgentPromptLLM = "gemini-2.0-flash-lite"
	Gemini25FlashLite AgentPromptLLM = "gemini-2.5-flash-lite"
	Gemini25Flash     AgentPromptLLM = "gemini-2.5-flash"
	ClaudeSonnet4     AgentPromptLLM = "claude-sonnet-4"
	Claude37Sonnet    AgentPromptLLM = "claude-3-7-sonnet"
	Claude35Sonnet    AgentPromptLLM = "claude-3-5-sonnet"
	Claude35SonnetV1  AgentPromptLLM = "claude-3-5-sonnet-v1"
	Claude3Haiku      AgentPromptLLM = "claude-3-haiku"
	GrokBeta          AgentPromptLLM = "grok-beta"
	CustomLLM         AgentPromptLLM = "custom-llm"
	Qwen34B           AgentPromptLLM = "qwen3-4b"
	Qwen330BA3B       AgentPromptLLM = "qwen3-30b-a3b"
	WattTool8B        AgentPromptLLM = "watt-tool-8b"
	WattTool70B       AgentPromptLLM = "watt-tool-70b"

	// dated variants
	Gemini25FlashPreview0520     AgentPromptLLM = "gemini-2.5-flash-preview-05-20"
	Gemini25FlashPreview0417     AgentPromptLLM = "gemini-2.5-flash-preview-04-17"
	Gemini25FlashLitePreview0617 AgentPromptLLM = "gemini-2.5-flash-lite-preview-06-17"
	Gemini20FlashLite001         AgentPromptLLM = "gemini-2.0-flash-lite-001"
	Gemini20Flash001             AgentPromptLLM = "gemini-2.0-flash-001"
	Gemini15Flash002             AgentPromptLLM = "gemini-1.5-flash-002"
	Gemini15Flash001             AgentPromptLLM = "gemini-1.5-flash-001"
	Gemini15Pro002               AgentPromptLLM = "gemini-1.5-pro-002"
	Gemini15Pro001               AgentPromptLLM = "gemini-1.5-pro-001"
	ClaudeSonnet4_20250514       AgentPromptLLM = "claude-sonnet-4@20250514"
	Claude37Sonnet_20250219      AgentPromptLLM = "claude-3-7-sonnet@20250219"
	Claude35Sonnet_20240620      AgentPromptLLM = "claude-3-5-sonnet@20240620"
	Claude35SonnetV2_20241022    AgentPromptLLM = "claude-3-5-sonnet-v2@20241022"
	Claude3Haiku_20240307        AgentPromptLLM = "claude-3-haiku@20240307"
	GPT5_20250807                AgentPromptLLM = "gpt-5-2025-08-07"
	GPT5Mini_20250807            AgentPromptLLM = "gpt-5-mini-2025-08-07"
	GPT5Nano_20250807            AgentPromptLLM = "gpt-5-nano-2025-08-07"
	GPT41_20250414               AgentPromptLLM = "gpt-4.1-2025-04-14"
	GPT41Mini_20250414           AgentPromptLLM = "gpt-4.1-mini-2025-04-14"
	GPT41Nano_20250414           AgentPromptLLM = "gpt-4.1-nano-2025-04-14"
	GPT4oMini_20240718           AgentPromptLLM = "gpt-4o-mini-2024-07-18"
	GPT4o_20241120               AgentPromptLLM = "gpt-4o-2024-11-20"
	GPT4o_20240806               AgentPromptLLM = "gpt-4o-2024-08-06"
	GPT4o_20240513               AgentPromptLLM = "gpt-4o-2024-05-13"
	GPT40613                     AgentPromptLLM = "gpt-4-0613"
	GPT40314                     AgentPromptLLM = "gpt-4-0314"
	GPT4Turbo_20240409           AgentPromptLLM = "gpt-4-turbo-2024-04-09"
	GPT35Turbo_0125              AgentPromptLLM = "gpt-3.5-turbo-0125"
	GPT35Turbo_1106              AgentPromptLLM = "gpt-3.5-turbo-1106"
)

type AgentDynamicVariables struct {
	DynamicVariablePlaceholders map[string]any `json:"dynamic_variable_placeholders,omitempty"`
}

type ASRQuality string

const (
	ConversationConfigASRQualityHigh ASRQuality = "high"
)

type ASRProvider string

const (
	ConversationConfigASRProviderElevenlabs ASRProvider = "elevenlabs"
)

type AudioFormat string

const (
	AudioFormatPcm8000  AudioFormat = "pcm_8000"
	AudioFormatPcm22050 AudioFormat = "pcm_22050"
	AudioFormatPcm16000 AudioFormat = "pcm_16000"
	AudioFormatPcm24000 AudioFormat = "pcm_24000"
	AudioFormatPcm44100 AudioFormat = "pcm_44100"
	AudioFormatPcm48000 AudioFormat = "pcm_48000"
	AudioFormatUlaw8000 AudioFormat = "ulaw_8000"
)

type ConversationConfigASR struct {
	Quality              *ASRQuality  `json:"quality,omitempty"`
	Provider             *ASRProvider `json:"provider,omitempty"`
	UserInputAudioFormat *AudioFormat `json:"user_input_audio_format,omitempty"`
	Keywords             []string     `json:"keywords,omitempty"`
}

type TurnMode string

const (
	TurnModeSilence TurnMode = "silence"
	TurnModeTurn    TurnMode = "turn"
)

type ConversationConfigTurn struct {
	TurnTimeout           *float32  `json:"turn_timeout,omitempty"`
	SilenceEndCallTimeout *float32  `json:"silence_end_call_timeout,omitempty"`
	Mode                  *TurnMode `json:"mode,omitempty"`
}

type ConversationConfigTTS struct {
	ModelId                          *TTSModelId                       `json:"model_id,omitempty"`
	Voice                            *string                           `json:"voice,omitempty"`
	SupportedVoices                  []SupportedVoice                  `json:"supported_voices,omitempty"`
	AgentOutputAudioFormat           *AudioFormat                      `json:"agent_output_audio_format,omitempty"`
	OptimizeStreamingLatency         *OptimizeStreamingLatency         `json:"optimize_streaming_latency,omitempty"`
	Stability                        *float32                          `json:"stability,omitempty"`
	Speed                            *float32                          `json:"speed,omitempty"`
	SimilarityBoost                  *float32                          `json:"similarity_boost,omitempty"`
	PronounciationDictionaryLocators []PronounciationDictionaryLocator `json:"pronounciation_dictionary_locators,omitempty"`
}

type TTSModelId string

const (
	TTSModelIdElevenTurboV2   TTSModelId = "eleven_turbo_v2"
	TTSModelIdElevenTurboV2_5 TTSModelId = "eleven_turbo_v2_5"
	TTSModelIdElevenFlashV2   TTSModelId = "eleven_flash_v2"
	TTSModelIdElevenFlashV2_5 TTSModelId = "eleven_flash_v2_5"
)

type SupportedVoice struct {
	Label                    string                    `json:"label"`
	VoiceId                  string                    `json:"voice_id"`
	Description              *string                   `json:"description,omitempty"`
	Language                 *string                   `json:"language,omitempty"`
	ModelFamily              *ModelFamily              `json:"model_family,omitempty"`
	OptimizeStreamingLatency *OptimizeStreamingLatency `json:"optimize_streaming_latency,omitempty"`
	Stability                *float32                  `json:"stability,omitempty"`
	Speed                    *float32                  `json:"speed,omitempty"`
	SimilarityBoost          *float32                  `json:"similarity_boost,omitempty"`
}

type OptimizeStreamingLatency int

const (
	OptimizeStreamingLatencyNone     OptimizeStreamingLatency = 0
	OptimizeStreamingLatencyLow      OptimizeStreamingLatency = 1
	OptimizeStreamingLatencyMedium   OptimizeStreamingLatency = 2
	OptimizeStreamingLatencyHigh     OptimizeStreamingLatency = 3
	OptimizeStreamingLatencyVeryHigh OptimizeStreamingLatency = 4
)

type ModelFamily string

const (
	ModelFamilyTurbo        ModelFamily = "turbo"
	ModelFamilyFlash        ModelFamily = "flash"
	ModelFamilyMultilingual ModelFamily = "multilingual"
)

type PronounciationDictionaryLocator struct {
	PronounciationDictionaryId string  `json:"pronounciation_dictionary_id"`
	VersionId                  *string `json:"version_id,omitempty"`
}

type ConversationConfigConversation struct {
	TextOnly           *bool         `json:"text_only,omitempty"`
	MaxDurationSeconds *int          `json:"max_duration_seconds,omitempty"`
	ClientEvents       []ClientEvent `json:"client_events,omitempty"`
}

type ClientEvent string

const (
	ClientEventConversationInitiationMetadata    ClientEvent = "conversation_initiation_metadata"
	ClientEventConversationASRInitiationMetadata ClientEvent = "asr_initiation_metadata"
	ClientEventPing                              ClientEvent = "ping"
	ClientEventAudio                             ClientEvent = "audio"
	ClientEventInterruption                      ClientEvent = "interruption"
	ClientEventUserTranscript                    ClientEvent = "user_transcript"
	ClientEventAgentResponse                     ClientEvent = "agent_response"
	ClientEventAgentResponseCorrection           ClientEvent = "agent_response_correction"
	ClientEventClientToolCall                    ClientEvent = "client_tool_call"
	ClientEventMCPToolCall                       ClientEvent = "mcp_tool_call"
	ClientEventMCPConnectionStatus               ClientEvent = "mcp_connection_status"
	ClientEventAgentToolResponse                 ClientEvent = "agent_tool_response"
	ClientEventVADScore                          ClientEvent = "vad_score"
	ClientEventInternalTurnProbability           ClientEvent = "internal_turn_probability"
	ClientEventInternalTentativeAgentResponse    ClientEvent = "internal_tentative_agent_response"
)

type ConversationConfigLanguagePreset struct {
	Overrides               LanguagePresetOverrides                `json:"overrides"`
	FirstMessageTranslation *LanguagePresetFirstMessageTranslation `json:"first_message_translation,omitempty"`
}

type LanguagePresetFirstMessageTranslation struct {
	SourceHash string `json:"source_hash"`
	Text       string `json:"text"`
}

type LanguagePresetOverrides struct {
	TTS          *LanguagePresetOverridesTTS `json:"tts,omitempty"`
	Conversation *LanguagePresetConversation `json:"conversation,omitempty"`
	Agent        *LanguagePresetAgent        `json:"agent,omitempty"`
}

type LanguagePresetConversation struct {
	TextOnly *bool `json:"text_only,omitempty"`
}

type LanguagePresetOverridesTTS struct {
	VoiceId *string `json:"voice_id"`
}

type LanguagePresetAgent struct {
	FirstMessage *string                    `json:"first_message,omitempty"`
	Language     *string                    `json:"language,omitempty"`
	Prompt       *LanguagePresetAgentPrompt `json:"prompt,omitempty"`
}

type LanguagePresetAgentPrompt struct {
	Prompt             *string  `json:"prompt,omitempty"`
	NativeMCPServerIds []string `json:"native_mcp_server_ids,omitempty"`
}
