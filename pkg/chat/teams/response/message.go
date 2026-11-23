package response

import "time"

type Message struct {
	Identifier  string     `json:"id"`
	MessageType string     `json:"messageType"`
	Created     time.Time  `json:"createdDateTime"`
	Deleted     *time.Time `json:"deletedDateTime"`
	Subject     string     `json:"subject"`
	WebLocator  string     `json:"webUrl"`
	From        *Sender    `json:"from"`
	Body        Body       `json:"body"`
	Replies     []Message  `json:"replies"`
}
