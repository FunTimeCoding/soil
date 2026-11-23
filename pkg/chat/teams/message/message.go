package message

import "time"

type Message struct {
	Identifier string    `json:"identifier"`
	Author     string    `json:"author"`
	Created    time.Time `json:"created"`
	Subject    string    `json:"subject,omitempty"`
	Text       string    `json:"text"`
	Replies    int       `json:"replies,omitzero"`
	Link       string    `json:"link,omitempty"`
}
