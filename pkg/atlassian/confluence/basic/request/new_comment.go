package request

import "fmt"

func NewComment(
	pageIdentifier string,
	body string,
) *Comment {
	return &Comment{
		Type: "comment",
		Container: CommentContainer{
			Identifier: pageIdentifier,
			Type:       "page",
		},
		Body: CommentBody{
			Storage: CommentStorage{
				Value:          fmt.Sprintf("<p>%s</p>", body),
				Representation: "storage",
			},
		},
	}
}
