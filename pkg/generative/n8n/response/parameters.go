package response

type Parameters struct {
	ChannelIdentifier any          `json:"channelId,omitempty"`
	Message           string       `json:"message,omitempty"`
	Attachments       []any        `json:"attachments,omitempty"`
	OtherOptions      struct{}     `json:"otherOptions,omitempty"`
	Owner             ResourceLink `json:"owner,omitempty"`
	Repository        ResourceLink `json:"repository,omitempty"`
	Events            []string     `json:"events,omitempty"`
	Options           Options      `json:"options,omitempty"`
	Content           string       `json:"content,omitempty"`
	Height            int          `json:"height,omitzero"`
	Width             int          `json:"width,omitzero"`
	PromptType        string       `json:"promptType,omitempty"`
	Text              string       `json:"text,omitempty"`
	AdditionalFields  struct{}     `json:"additionalFields,omitempty"`
	Updates           []string     `json:"updates,omitempty"`
	Resource          string       `json:"resource,omitempty"`
	Operation         string       `json:"operation,omitempty"`
	GuildIdentifier   ResourceLink `json:"guildId,omitempty"`
	Authentication    string       `json:"authentication,omitempty"`
}
