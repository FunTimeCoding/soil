package argument

type UnsubscribeMergeRequest struct {
	Project      string `json:"project"`
	MergeRequest int64  `json:"merge_request"`
	Callsign     string `json:"callsign"`
}
