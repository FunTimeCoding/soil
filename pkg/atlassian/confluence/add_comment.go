package confluence

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/notation"
)

func (c *Client) AddComment(
	pageIdentifier string,
	body string,
) error {
	payload := commentPayload{
		Type: "comment",
		Container: commentContainer{
			Identifier: pageIdentifier,
			Type:       "page",
		},
		Body: commentBody{
			Storage: commentStorage{
				Value:          fmt.Sprintf("<p>%s</p>", body),
				Representation: "storage",
			},
		},
	}
	_, e := c.basic.PostOldPath("/content", notation.Encode(payload, false))

	return e
}
