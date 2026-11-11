package response

type Stanza struct {
	Name   string    `json:"name"`
	Status *Status   `json:"status"`
	Backup []*Backup `json:"backup"`
}
