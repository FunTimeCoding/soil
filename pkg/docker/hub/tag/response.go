package tag

import "time"

type Response struct {
	Name        string     `json:"name"`
	LastUpdated *time.Time `json:"last_updated"`
}
