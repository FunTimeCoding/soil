package response

type Backup struct {
	Label     string     `json:"label"`
	Type      string     `json:"type"`
	Error     bool       `json:"error"`
	Size      *Size      `json:"info"`
	Timestamp *Timestamp `json:"timestamp"`
}
