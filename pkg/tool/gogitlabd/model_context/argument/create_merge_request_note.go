package argument

type CreateMergeRequestNote struct {
	Project      string `json:"project"`
	MergeRequest int64  `json:"merge_request"`
	Body         string `json:"body"`
	Callsign     string `json:"callsign"`
}
