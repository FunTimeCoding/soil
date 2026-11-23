package response

import "time"

type Chat struct {
	Identifier  string    `json:"id"`
	Topic       string    `json:"topic"`
	ChatType    string    `json:"chatType"`
	LastUpdated time.Time `json:"lastUpdatedDateTime"`
	WebLocator  string    `json:"webUrl"`
}
