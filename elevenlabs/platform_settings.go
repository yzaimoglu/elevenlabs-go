package elevenlabs

type PlatformSettings struct {
	Auth               *PlatformSettingsAuth                     `json:"auth,omitempty"`
	Evaluation         *PlatformSettingsEvaluation               `json:"evaluation,omitempty"`
	Widget             *PlatformSettingsWidget                   `json:"widget,omitempty"`
	DataCollection     map[string]PlatformSettingsDataCollection `json:"data_collection,omitempty"`
	Overrides          *PlatformSettingsOverrides                `json:"overrides,omitempty"`
	CallLimits         *PlatformSettingsCallLimits               `json:"call_limits,omitempty"`
	Privacy            *PlatformSettingsPrivacy                  `json:"privacy,omitempty"`
	WorkspaceOverrides *PlatformSettingsWorkspaceOverrides       `json:"workspace_overrides,omitempty"`
	Testing            *PlatformSettingsTesting                  `json:"testing,omitempty"`
	Safety             *PlatformSettingsSafety                   `json:"safety,omitempty"`
}

type PlatformSettingsAuth struct {
	EnableAuth     *bool                           `json:"enable_auth,omitempty"`
	AllowList      []PlatformSettingsAuthAllowList `json:"allowlist,omitempty"`
	ShareableToken *string                         `json:"shareable_token,omitempty"`
}

type PlatformSettingsAuthAllowList struct {
	Hostname string `json:"hostname"`
}

type PlatformSettingsEvaluation struct {
	Criteria []PlatformSettingsEvaluationCriteria `json:"criteria,omitempty"`
}

type PlatformSettingsEvaluationCriteria struct {
	ID                     string                                  `json:"id"`
	Name                   string                                  `json:"name"`
	ConversationGoalPrompt string                                  `json:"conversation_goal_prompt"`
	Type                   *PlatformSettingsEvaluationCriteriaType `json:"type,omitempty"`
	UseKnowledgeBase       *bool                                   `json:"use_knowledge_base,omitempty"`
}

type PlatformSettingsEvaluationCriteriaType string

const (
	PlatformSettingsEvaluationCriteriaTypePrompt PlatformSettingsEvaluationCriteriaType = "prompt"
)

type PlatformSettingsWidget struct {
	Variant                 *PlatformSettingsWidgetVariant                `json:"variant,omitempty"`
	Placement               *PlatformSettingsWidgetPlacement              `json:"placement,omitempty"`
	Expandable              *PlatformSettingsWidgetExpandable             `json:"expandable,omitempty"`
	Avatar                  *PlatformSettingsWidgetAvatar                 `json:"avatar,omitempty"`
	FeedbackMode            *PlatformSettingsWidgetFeedbackMode           `json:"feedback_mode,omitempty"`
	BgColor                 *string                                       `json:"bg_color,omitempty"`
	TextColor               *string                                       `json:"text_color,omitempty"`
	BtnColor                *string                                       `json:"btn_color,omitempty"`
	BtnTextColor            *string                                       `json:"btn_text_color,omitempty"`
	BorderColor             *string                                       `json:"border_color,omitempty"`
	FocusColor              *string                                       `json:"focus_color,omitempty"`
	BorderRadius            *int                                          `json:"border_radius,omitempty"`
	BtnRadius               *int                                          `json:"btn_radius,omitempty"`
	ActionText              *string                                       `json:"action_text,omitempty"`
	StartCallText           *string                                       `json:"start_call_text,omitempty"`
	EndCallText             *string                                       `json:"end_call_text,omitempty"`
	ExpandText              *string                                       `json:"expand_text,omitempty"`
	ListeningText           *string                                       `json:"listening_text,omitempty"`
	SpeakingText            *string                                       `json:"speaking_text,omitempty"`
	ShareablePageText       *string                                       `json:"shareable_page_text,omitempty"`
	ShareablePageShowTerms  *bool                                         `json:"shareable_page_show_terms,omitempty"`
	TermsText               *string                                       `json:"terms_text,omitempty"`
	TermsHTML               *string                                       `json:"terms_html,omitempty"`
	TermsKey                *string                                       `json:"terms_key,omitempty"`
	ShowAvatarWhenCollapsed *bool                                         `json:"show_avatar_when_collapsed,omitempty"`
	DisableBanner           *bool                                         `json:"disable_banner,omitempty"`
	OverrideLink            *string                                       `json:"override_link,omitempty"`
	MicMutingEnabled        *bool                                         `json:"mic_muting_enabled,omitempty"`
	TranscriptEnabled       *bool                                         `json:"transcript_enabled,omitempty"`
	TextInputEnabled        *bool                                         `json:"text_input_enabled,omitempty"`
	DefaultExpanded         *bool                                         `json:"default_expanded,omitempty"`
	AlwaysExpanded          *bool                                         `json:"always_expanded,omitempty"`
	TextContents            *PlatformSettingsWidgetTextContents           `json:"text_contents,omitempty"`
	Styles                  *PlatformSettingsWidgetStyles                 `json:"styles,omitempty"`
	LanguageSelector        *bool                                         `json:"language_selector,omitempty"`
	SupportsTextOnly        *bool                                         `json:"supports_text_only,omitempty"`
	CustomAvatarPath        *string                                       `json:"custom_avatar_path,omitempty"`
	LanguagePresets         map[string]PlatformSettingsWidgetTextContents `json:"language_presets,omitempty"`
}

type PlatformSettingsWidgetVariant string

const (
	PlatformSettingsWidgetVariantTiny       PlatformSettingsWidgetVariant = "tiny"
	PlatformSettingsWidgetVariantCompact    PlatformSettingsWidgetVariant = "compact"
	PlatformSettingsWidgetVariantFull       PlatformSettingsWidgetVariant = "full"
	PlatformSettingsWidgetVariantExpandable PlatformSettingsWidgetVariant = "expandable"
)

type PlatformSettingsWidgetPlacement string

const (
	PlatformSettingsWidgetPlacementTopLeft     PlatformSettingsWidgetPlacement = "top_left"
	PlatformSettingsWidgetPlacementTop         PlatformSettingsWidgetPlacement = "top"
	PlatformSettingsWidgetPlacementTopRight    PlatformSettingsWidgetPlacement = "top_right"
	PlatformSettingsWidgetPlacementBottomLeft  PlatformSettingsWidgetPlacement = "bottom_left"
	PlatformSettingsWidgetPlacementBottom      PlatformSettingsWidgetPlacement = "bottom"
	PlatformSettingsWidgetPlacementBottomRight PlatformSettingsWidgetPlacement = "bottom_right"
)

type PlatformSettingsWidgetExpandable string

const (
	PlatformSettingsWidgetExpandableNever   PlatformSettingsWidgetExpandable = "never"
	PlatformSettingsWidgetExpandableMobile  PlatformSettingsWidgetExpandable = "mobile"
	PlatformSettingsWidgetExpandableDesktop PlatformSettingsWidgetExpandable = "desktop"
	PlatformSettingsWidgetExpandableAlways  PlatformSettingsWidgetExpandable = "always"
)

type PlatformSettingsWidgetAvatar struct {
	OrbAvatar   PlatformSettingsWidgetAvatarOrb   `json:"orb_avatar,omitempty"`
	URLAvatar   PlatformSettingsWidgetAvatarURL   `json:"url_avatar,omitempty"`
	ImageAvatar PlatformSettingsWidgetAvatarImage `json:"image_avatar,omitempty"`
}

type PlatformSettingsWidgetAvatarOrb struct {
	Type   *PlatformSettingsWidgetAvatarOrbType `json:"type,omitempty"`
	Color1 *string                              `json:"color_1,omitempty"`
	Color2 *string                              `json:"color_2,omitempty"`
}

type PlatformSettingsWidgetAvatarOrbType string

const (
	PlatformSettingsWidgetAvatarOrbTypeOrb PlatformSettingsWidgetAvatarOrbType = "orb"
)

type PlatformSettingsWidgetAvatarURL struct {
	Type      *PlatformSettingsWidgetAvatarURLType `json:"type,omitempty"`
	CustomURL *string                              `json:"custom_url,omitempty"`
}

type PlatformSettingsWidgetAvatarURLType string

const (
	PlatformSettingsWidgetAvatarURLTypeURL PlatformSettingsWidgetAvatarURLType = "url"
)

type PlatformSettingsWidgetAvatarImage struct {
	Type *PlatformSettingsWidgetAvatarImageType `json:"type,omitempty"`
	URL  *string                                `json:"url,omitempty"`
}

type PlatformSettingsWidgetAvatarImageType string

const (
	PlatformSettingsWidgetAvatarImageTypeImage PlatformSettingsWidgetAvatarImageType = "image"
)

type PlatformSettingsWidgetFeedbackMode string

const (
	PlatformSettingsWidgetFeedbackModeNone   PlatformSettingsWidgetFeedbackMode = "none"
	PlatformSettingsWidgetFeedbackModeDuring PlatformSettingsWidgetFeedbackMode = "during"
	PlatformSettingsWidgetFeedbackModeEnd    PlatformSettingsWidgetFeedbackMode = "end"
)

type PlatformSettingsWidgetTextContents struct {
	MainLabel                       *string `json:"main_label,omitempty"`
	StartCall                       *string `json:"start_call,omitempty"`
	StartChat                       *string `json:"start_chat,omitempty"`
	NewCall                         *string `json:"new_call,omitempty"`
	EndCall                         *string `json:"end_call,omitempty"`
	MuteMicrophone                  *string `json:"mute_microphone,omitempty"`
	ChangeLanguage                  *string `json:"change_language,omitempty"`
	Collapse                        *string `json:"collapse,omitempty"`
	Expand                          *string `json:"expand,omitempty"`
	Copied                          *string `json:"copied,omitempty"`
	AcceptTerms                     *string `json:"accept_terms,omitempty"`
	DismissTerms                    *string `json:"dismiss_terms,omitempty"`
	ListeningStatus                 *string `json:"listening_status,omitempty"`
	SpeakingStatus                  *string `json:"speaking_status,omitempty"`
	ConnectingStatus                *string `json:"connecting_status,omitempty"`
	ChattingStatus                  *string `json:"chatting_status,omitempty"`
	InputLabel                      *string `json:"input_label,omitempty"`
	InputPlaceholder                *string `json:"input_placeholder,omitempty"`
	InputPlaceholderTextOnly        *string `json:"input_placeholder_text_only,omitempty"`
	InputPlaceholderNewConversation *string `json:"input_placeholder_new_conversation,omitempty"`
	UserEndedConversation           *string `json:"user_ended_conversation,omitempty"`
	AgentEndedConversation          *string `json:"agent_ended_conversation,omitempty"`
	ConversationId                  *string `json:"conversation_id,omitempty"`
	ErrorOccured                    *string `json:"error_occured,omitempty"`
	CopyId                          *string `json:"copy_id,omitempty"`
}

type PlatformSettingsWidgetStyles struct {
	Base                *string  `json:"base,omitempty"`
	BaseHover           *string  `json:"base_hover,omitempty"`
	BaseActive          *string  `json:"base_active,omitempty"`
	BaseBorder          *string  `json:"base_border,omitempty"`
	BaseSubtle          *string  `json:"base_subtle,omitempty"`
	BasePrimary         *string  `json:"base_primary,omitempty"`
	BaseError           *string  `json:"base_error,omitempty"`
	Accent              *string  `json:"accent,omitempty"`
	AccentHover         *string  `json:"accent_hover,omitempty"`
	AccentActive        *string  `json:"accent_active,omitempty"`
	AccentBorder        *string  `json:"accent_border,omitempty"`
	AccentSubtle        *string  `json:"accent_subtle,omitempty"`
	AccentPrimary       *string  `json:"accent_primary,omitempty"`
	OverlayPadding      *float32 `json:"overlay_padding,omitempty"`
	ButtonRadius        *float32 `json:"button_radius,omitempty"`
	InputRadius         *float32 `json:"input_radius,omitempty"`
	BubbleRadius        *float32 `json:"bubble_radius,omitempty"`
	SheetRadius         *float32 `json:"sheet_radius,omitempty"`
	CompactSheetRadius  *float32 `json:"compact_sheet_radius,omitempty"`
	DropdownSheetRadius *float32 `json:"dropdown_sheet_radius,omitempty"`
}

type PlatformSettingsDataCollection struct {
	Type            PlatformSettingsDataCollectionType `json:"type"`
	Description     *string                            `json:"description,omitempty"`
	DynamicVariable *string                            `json:"dynamic_variable,omitempty"`
	// string or integer or double or boolean or null
	ConstantValue any `json:"constant_value,omitempty"`
}

type PlatformSettingsDataCollectionType string

const (
	PlatformSettingsDataCollectionTypeBoolean PlatformSettingsDataCollectionType = "boolean"
	PlatformSettingsDataCollectionTypeNumber  PlatformSettingsDataCollectionType = "number"
	PlatformSettingsDataCollectionTypeString  PlatformSettingsDataCollectionType = "string"
	PlatformSettingsDataCollectionTypeInteger PlatformSettingsDataCollectionType = "integer"
)

type PlatformSettingsOverrides struct {
	ConversationConfigOverride                        *PlatformSettingsOverridesConversationConfigOverride `json:"conversation_config_override,omitempty"`
	CustomLLMExtraBody                                *bool                                                `json:"custom_llm_extra_body,omitempty"`
	EnableConversationInitiationClientDataFromWebhook *bool                                                `json:"enable_conversation_initiation_client_data_from_webhook,omitempty"`
}

type PlatformSettingsOverridesConversationConfigOverride struct {
	TTS          *PlatformSettingsOverridesConversationConfigOverrideTTS          `json:"tts,omitempty"`
	Conversation *PlatformSettingsOverridesConversationConfigOverrideConversation `json:"conversation,omitempty"`
	Agent        *PlatformSettingsOverridesConversationConfigOverrideAgent        `json:"agent,omitempty"`
}

type PlatformSettingsOverridesConversationConfigOverrideTTS struct {
	VoiceId *bool `json:"voice_id,omitempty"`
}

type PlatformSettingsOverridesConversationConfigOverrideConversation struct {
	TextOnly *bool `json:"text_only,omitempty"`
}

type PlatformSettingsOverridesConversationConfigOverrideAgent struct {
	FirstMessage *bool                                                           `json:"first_message,omitempty"`
	Language     *bool                                                           `json:"language,omitempty"`
	Prompt       *PlatformSettingsOverridesConversationConfigOverrideAgentPrompt `json:"prompt,omitempty"`
}

type PlatformSettingsOverridesConversationConfigOverrideAgentPrompt struct {
	Prompt             *bool `json:"prompt,omitempty"`
	NativeMCPServerIds *bool `json:"native_mcp_server_ids,omitempty"`
}

type PlatformSettingsCallLimits struct {
	AgentConcurrencyLimit *int  `json:"agent_concurrency_limit,omitempty"`
	DailyLimit            *int  `json:"daily_limit,omitempty"`
	BurstingEnabled       *bool `json:"bursting_enabled,omitempty"`
}

type PlatformSettingsPrivacy struct {
	RecordVoice                  *bool `json:"record_voice,omitempty"`
	RetentionDays                *int  `json:"retention_days,omitempty"`
	DeleteTranscriptAndPII       *bool `json:"delete_transcript_and_pii,omitempty"`
	DeleteAudio                  *bool `json:"delete_audio,omitempty"`
	ApplyToExistingConversations *bool `json:"apply_to_existing_conversations,omitempty"`
	ZeroRetentionMode            *bool `json:"zero_retention_mode,omitempty"`
}

type PlatformSettingsWorkspaceOverrides struct {
	ConversationInitiationClientDataWebhook *PlatformSettingsWorkspaceOverridesConversationInitiationClientDataWebhook `json:"conversation_initiation_client_data_webhook,omitempty"`
	Webhooks                                *PlatformSettingsWorkspaceOverridesWebhooks                                `json:"webhooks,omitempty"`
}

type PlatformSettingsWorkspaceOverridesConversationInitiationClientDataWebhook struct {
	URL            string         `json:"url,omitempty"`
	RequestHeaders map[string]any `json:"request_headers,omitempty"`
}

type PlatformSettingsWorkspaceOverridesWebhooks struct {
	PostCallWebhookId *string `json:"post_call_webhook_id,omitempty"`
	SendAudio         *bool   `json:"send_audio,omitempty"`
}

type PlatformSettingsTesting struct {
	TestIds []string `json:"test_ids,omitempty"`
}

type PlatformSettingsSafety struct {
	IsBlockedIVC           *bool `json:"is_blocked_ivc,omitempty"`
	IsBlockedNonIVC        *bool `json:"is_blocked_non_ivc,omitempty"`
	IgnoreSafetyEvaluation *bool `json:"ignore_safety_evaluation,omitempty"`
}
