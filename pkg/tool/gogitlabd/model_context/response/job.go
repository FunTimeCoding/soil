package response

import "time"

type Job struct {
	Identifier int64      `json:"identifier"`
	Name       string     `json:"name"`
	Status     string     `json:"status"`
	Stage      string     `json:"stage"`
	Create     *time.Time `json:"create,omitempty"`
	Link       string     `json:"link"`
	Trace      string     `json:"trace,omitempty"`
}
