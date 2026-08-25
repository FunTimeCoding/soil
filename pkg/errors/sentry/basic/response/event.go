package response

import "time"

type Event struct {
	Identifier      string                    `json:"id"`
	EventIdentifier string                    `json:"eventID"`
	Title           string                    `json:"title"`
	Culprit         string                    `json:"culprit"`
	DateCreated     *time.Time                `json:"dateCreated"`
	Project         string                    `json:"project"`
	Message         string                    `json:"message"`
	Platform        string                    `json:"platform"`
	Tags            []EventTag                `json:"tags"`
	Contexts        map[string]map[string]any `json:"contexts"`
	Entries         []EventEntry              `json:"entries"`
}
