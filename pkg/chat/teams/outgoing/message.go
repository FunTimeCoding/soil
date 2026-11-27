package outgoing

import "github.com/funtimecoding/soil/pkg/chat/teams/response"

type Message struct {
	Body     response.Body `json:"body"`
	Mentions []Mention     `json:"mentions,omitempty"`
}
