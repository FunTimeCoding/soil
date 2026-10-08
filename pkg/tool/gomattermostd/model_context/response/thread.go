package response

type Thread struct {
	PostIdentifier string         `json:"post_identifier"`
	Channel        string         `json:"channel,omitempty"`
	Author         string         `json:"author"`
	Message        string         `json:"message"`
	ReplyCount     int64          `json:"reply_count"`
	UnreadReplies  int64          `json:"unread_replies"`
	UnreadMentions int64          `json:"unread_mentions"`
	LastReplyAt    string         `json:"last_reply_at"`
	Participants   []*Participant `json:"participants,omitempty"`
}
