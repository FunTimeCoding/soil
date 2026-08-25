package response

import "time"

type Issue struct {
	Identifier      string     `json:"id"`
	ShortIdentifier string     `json:"shortId"`
	Title           string     `json:"title"`
	Culprit         string     `json:"culprit"`
	Status          string     `json:"status"`
	Level           string     `json:"level"`
	Type            string     `json:"type"`
	Permalink       string     `json:"permalink"`
	FirstSeen       *time.Time `json:"firstSeen"`
	LastSeen        *time.Time `json:"lastSeen"`
	Count           string     `json:"count"`
	UserCount       int        `json:"userCount"`
	Project         Project    `json:"project"`
}
