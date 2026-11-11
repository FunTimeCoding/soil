package response

type Size struct {
	Database   int64   `json:"size"`
	Repository *Stored `json:"repository"`
}
