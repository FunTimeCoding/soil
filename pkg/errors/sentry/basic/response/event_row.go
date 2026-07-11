package response

import "time"

type EventRow struct {
	Identifier string     `json:"id"`
	Title      string     `json:"title"`
	Message    string     `json:"message"`
	Project    string     `json:"project"`
	Timestamp  *time.Time `json:"timestamp"`
}
