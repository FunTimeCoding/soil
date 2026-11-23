package chat

import "time"

type Chat struct {
	Identifier string    `json:"identifier"`
	Topic      string    `json:"topic,omitempty"`
	Kind       string    `json:"kind"`
	Updated    time.Time `json:"updated"`
	Link       string    `json:"link,omitempty"`
}
