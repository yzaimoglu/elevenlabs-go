package elevenlabs

import (
	"bytes"
	"context"
	"errors"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/google/go-querystring/query"
)

type ConvaiWidgetAPI interface {
	GetAgentWidget(ctx context.Context, req *GetAgentWidgetReq) (GetAgentWidgetResp, error)
	PostAgentAvatar(ctx context.Context, req *PostAgentAvatarReq) (PostAgentAvatarResp, error)
}

// EmbedVariant represents the widget variant
type EmbedVariant string

const (
	EmbedVariantTiny       EmbedVariant = "tiny"
	EmbedVariantCompact    EmbedVariant = "compact"
	EmbedVariantFull       EmbedVariant = "full"
	EmbedVariantExpandable EmbedVariant = "expandable"
)

// WidgetPlacement represents the widget placement on screen
type WidgetPlacement string

const (
	WidgetPlacementTopLeft     WidgetPlacement = "top-left"
	WidgetPlacementTop         WidgetPlacement = "top"
	WidgetPlacementTopRight    WidgetPlacement = "top-right"
	WidgetPlacementBottomLeft  WidgetPlacement = "bottom-left"
	WidgetPlacementBottom      WidgetPlacement = "bottom"
	WidgetPlacementBottomRight WidgetPlacement = "bottom-right"
)

// WidgetExpandable represents when the widget is expandable
type WidgetExpandable string

const (
	WidgetExpandableNever   WidgetExpandable = "never"
	WidgetExpandableMobile  WidgetExpandable = "mobile"
	WidgetExpandableDesktop WidgetExpandable = "desktop"
	WidgetExpandableAlways  WidgetExpandable = "always"
)

// WidgetFeedbackMode represents the feedback mode
type WidgetFeedbackMode string

const (
	WidgetFeedbackModeNone   WidgetFeedbackMode = "none"
	WidgetFeedbackModeDuring WidgetFeedbackMode = "during"
	WidgetFeedbackModeEnd    WidgetFeedbackMode = "end"
)

// WidgetEndFeedbackType represents the end feedback type
type WidgetEndFeedbackType string

const (
	WidgetEndFeedbackTypeRating WidgetEndFeedbackType = "rating"
)

// WidgetEndFeedbackConfig represents end feedback configuration
type WidgetEndFeedbackConfig struct {
	Type WidgetEndFeedbackType `json:"type"`
}

// WidgetAvatarType represents the avatar type
type WidgetAvatarType string

const (
	WidgetAvatarTypeOrb   WidgetAvatarType = "orb"
	WidgetAvatarTypeURL   WidgetAvatarType = "url"
	WidgetAvatarTypeImage WidgetAvatarType = "image"
)

// WidgetAvatar represents a widget avatar (can be orb, url, or image)
type WidgetAvatar struct {
	Type      WidgetAvatarType `json:"type"`
	Color1    string           `json:"color_1,omitempty"`
	Color2    string           `json:"color_2,omitempty"`
	CustomURL string           `json:"custom_url,omitempty"`
	URL       string           `json:"url,omitempty"`
}

// AllowlistItem represents an allowed hostname
type AllowlistItem struct {
	Hostname string `json:"hostname"`
}

// WidgetTextContents represents text contents for the widget
type WidgetTextContents struct {
	MainLabel                      *string `json:"main_label,omitempty"`
	StartCall                      *string `json:"start_call,omitempty"`
	StartChat                      *string `json:"start_chat,omitempty"`
	NewCall                        *string `json:"new_call,omitempty"`
	EndCall                        *string `json:"end_call,omitempty"`
	MuteMicrophone                 *string `json:"mute_microphone,omitempty"`
	ChangeLanguage                 *string `json:"change_language,omitempty"`
	Collapse                       *string `json:"collapse,omitempty"`
	Expand                         *string `json:"expand,omitempty"`
	Copied                         *string `json:"copied,omitempty"`
	AcceptTerms                    *string `json:"accept_terms,omitempty"`
	DismissTerms                   *string `json:"dismiss_terms,omitempty"`
	ListeningStatus                *string `json:"listening_status,omitempty"`
	SpeakingStatus                 *string `json:"speaking_status,omitempty"`
	ConnectingStatus               *string `json:"connecting_status,omitempty"`
	ChattingStatus                 *string `json:"chatting_status,omitempty"`
	InputLabel                     *string `json:"input_label,omitempty"`
	InputPlaceholder               *string `json:"input_placeholder,omitempty"`
	InputPlaceholderTextOnly       *string `json:"input_placeholder_text_only,omitempty"`
	InputPlaceholderNewConversation *string `json:"input_placeholder_new_conversation,omitempty"`
	UserEndedConversation          *string `json:"user_ended_conversation,omitempty"`
	AgentEndedConversation         *string `json:"agent_ended_conversation,omitempty"`
	ConversationId                 *string `json:"conversation_id,omitempty"`
	ErrorOccurred                  *string `json:"error_occurred,omitempty"`
	CopyId                         *string `json:"copy_id,omitempty"`
	InitiateFeedback               *string `json:"initiate_feedback,omitempty"`
	RequestFollowUpFeedback        *string `json:"request_follow_up_feedback,omitempty"`
	ThanksForFeedback              *string `json:"thanks_for_feedback,omitempty"`
	ThanksForFeedbackDetails       *string `json:"thanks_for_feedback_details,omitempty"`
	FollowUpFeedbackPlaceholder    *string `json:"follow_up_feedback_placeholder,omitempty"`
	Submit                         *string `json:"submit,omitempty"`
	GoBack                         *string `json:"go_back,omitempty"`
}

// WidgetStyles represents widget styles
type WidgetStyles struct {
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
	OverlayPadding      *float64 `json:"overlay_padding,omitempty"`
	ButtonRadius        *float64 `json:"button_radius,omitempty"`
	InputRadius         *float64 `json:"input_radius,omitempty"`
	BubbleRadius        *float64 `json:"bubble_radius,omitempty"`
	SheetRadius         *float64 `json:"sheet_radius,omitempty"`
	CompactSheetRadius  *float64 `json:"compact_sheet_radius,omitempty"`
	DropdownSheetRadius *float64 `json:"dropdown_sheet_radius,omitempty"`
}

// WidgetLanguagePreset represents a language preset for the widget
type WidgetLanguagePreset struct {
	FirstMessage *string             `json:"first_message,omitempty"`
	TextContents *WidgetTextContents `json:"text_contents,omitempty"`
	TermsText    *string             `json:"terms_text,omitempty"`
	TermsHTML    *string             `json:"terms_html,omitempty"`
	TermsKey     *string             `json:"terms_key,omitempty"`
}

// WidgetConfig represents the widget configuration
type WidgetConfig struct {
	Variant                       EmbedVariant                    `json:"variant,omitempty"`
	Placement                     WidgetPlacement                 `json:"placement,omitempty"`
	Expandable                    WidgetExpandable                `json:"expandable,omitempty"`
	Avatar                        *WidgetAvatar                   `json:"avatar,omitempty"`
	FeedbackMode                  WidgetFeedbackMode              `json:"feedback_mode,omitempty"`
	EndFeedback                   *WidgetEndFeedbackConfig        `json:"end_feedback,omitempty"`
	BgColor                       string                          `json:"bg_color,omitempty"`
	TextColor                     string                          `json:"text_color,omitempty"`
	BtnColor                      string                          `json:"btn_color,omitempty"`
	BtnTextColor                  string                          `json:"btn_text_color,omitempty"`
	BorderColor                   string                          `json:"border_color,omitempty"`
	FocusColor                    string                          `json:"focus_color,omitempty"`
	BorderRadius                  *int                            `json:"border_radius,omitempty"`
	BtnRadius                     *int                            `json:"btn_radius,omitempty"`
	ActionText                    *string                         `json:"action_text,omitempty"`
	StartCallText                 *string                         `json:"start_call_text,omitempty"`
	EndCallText                   *string                         `json:"end_call_text,omitempty"`
	ExpandText                    *string                         `json:"expand_text,omitempty"`
	ListeningText                 *string                         `json:"listening_text,omitempty"`
	SpeakingText                  *string                         `json:"speaking_text,omitempty"`
	ShareablePageText             *string                         `json:"shareable_page_text,omitempty"`
	ShareablePageShowTerms        bool                            `json:"shareable_page_show_terms"`
	TermsText                     *string                         `json:"terms_text,omitempty"`
	TermsHTML                     *string                         `json:"terms_html,omitempty"`
	TermsKey                      *string                         `json:"terms_key,omitempty"`
	ShowAvatarWhenCollapsed       *bool                           `json:"show_avatar_when_collapsed,omitempty"`
	DisableBanner                 bool                            `json:"disable_banner"`
	OverrideLink                  *string                         `json:"override_link,omitempty"`
	MarkdownLinkAllowedHosts      []AllowlistItem                 `json:"markdown_link_allowed_hosts,omitempty"`
	MarkdownLinkIncludeWww        bool                            `json:"markdown_link_include_www"`
	MarkdownLinkAllowHttp         bool                            `json:"markdown_link_allow_http"`
	MicMutingEnabled              bool                            `json:"mic_muting_enabled"`
	TranscriptEnabled             bool                            `json:"transcript_enabled"`
	TextInputEnabled              bool                            `json:"text_input_enabled"`
	ConversationModeToggleEnabled bool                            `json:"conversation_mode_toggle_enabled"`
	DefaultExpanded               bool                            `json:"default_expanded"`
	AlwaysExpanded                bool                            `json:"always_expanded"`
	TextContents                  *WidgetTextContents             `json:"text_contents,omitempty"`
	Styles                        *WidgetStyles                   `json:"styles,omitempty"`
	Language                      string                          `json:"language"`
	SupportedLanguageOverrides    []string                        `json:"supported_language_overrides,omitempty"`
	LanguagePresets               map[string]WidgetLanguagePreset `json:"language_presets,omitempty"`
	TextOnly                      bool                            `json:"text_only"`
	SupportsTextOnly              bool                            `json:"supports_text_only"`
	FirstMessage                  *string                         `json:"first_message,omitempty"`
	UseRTC                        *bool                           `json:"use_rtc,omitempty"`
}

// GetAgentWidgetReq represents the request for getting an agent's widget config
type GetAgentWidgetReq struct {
	AgentId               string  `path:"agent_id"`
	ConversationSignature *string `url:"conversation_signature,omitempty"`
}

func NewGetAgentWidgetReq(agentId string) *GetAgentWidgetReq {
	return &GetAgentWidgetReq{
		AgentId: agentId,
	}
}

func (r GetAgentWidgetReq) QueryString() string {
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

// GetAgentWidgetResp represents the response from getting widget configuration
type GetAgentWidgetResp struct {
	AgentId      string       `json:"agent_id"`
	WidgetConfig WidgetConfig `json:"widget_config"`
}

// GetAgentWidget retrieves the widget configuration for an agent.
// https://elevenlabs.io/docs/api-reference/widget/get
func (c *Client) GetAgentWidget(ctx context.Context, req *GetAgentWidgetReq) (GetAgentWidgetResp, error) {
	if req == nil {
		return GetAgentWidgetResp{}, errors.New("request is nil")
	}

	body, err := c.get(ctx, "/convai/agents/"+req.AgentId+"/widget"+req.QueryString())
	if err != nil {
		return GetAgentWidgetResp{}, err
	}

	var resp GetAgentWidgetResp
	if err := c.parseResponse(body, &resp); err != nil {
		return GetAgentWidgetResp{}, err
	}

	return resp, nil
}

// PostAgentAvatarReq represents the request for posting an agent avatar
type PostAgentAvatarReq struct {
	AgentId    string
	AvatarFile io.Reader
	Filename   string
}

func NewPostAgentAvatarReq(agentId string, avatarFile io.Reader, filename string) *PostAgentAvatarReq {
	return &PostAgentAvatarReq{
		AgentId:    agentId,
		AvatarFile: avatarFile,
		Filename:   filename,
	}
}

// PostAgentAvatarResp represents the response from posting an agent avatar
type PostAgentAvatarResp struct {
	AgentId   string  `json:"agent_id"`
	AvatarURL *string `json:"avatar_url,omitempty"`
}

// PostAgentAvatar sets the avatar for an agent displayed in the widget.
// https://elevenlabs.io/docs/api-reference/widget/create
func (c *Client) PostAgentAvatar(ctx context.Context, req *PostAgentAvatarReq) (PostAgentAvatarResp, error) {
	if req == nil {
		return PostAgentAvatarResp{}, errors.New("request is nil")
	}

	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	part, err := writer.CreateFormFile("avatar_file", req.Filename)
	if err != nil {
		return PostAgentAvatarResp{}, err
	}

	if _, err := io.Copy(part, req.AvatarFile); err != nil {
		return PostAgentAvatarResp{}, err
	}

	if err := writer.Close(); err != nil {
		return PostAgentAvatarResp{}, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL.String()+"/convai/agents/"+req.AgentId+"/avatar", &buf)
	if err != nil {
		return PostAgentAvatarResp{}, err
	}

	httpReq.Header.Set("Content-Type", writer.FormDataContentType())
	httpReq = c.prepareRequest(ctx, httpReq)

	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return PostAgentAvatarResp{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return PostAgentAvatarResp{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return PostAgentAvatarResp{}, NewError(resp.StatusCode, body, resp)
	}

	var result PostAgentAvatarResp
	if err := c.parseResponse(body, &result); err != nil {
		return PostAgentAvatarResp{}, err
	}

	return result, nil
}
