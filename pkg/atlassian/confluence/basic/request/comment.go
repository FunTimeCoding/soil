package request

type Comment struct {
	Type      string           `json:"type"`
	Container CommentContainer `json:"container"`
	Body      CommentBody      `json:"body"`
}
