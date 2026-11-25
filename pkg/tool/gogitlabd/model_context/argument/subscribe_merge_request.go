package argument

type SubscribeMergeRequest struct {
	Project      string `json:"project"`
	MergeRequest int64  `json:"merge_request"`
	Callsign     string `json:"callsign"`
	Alias        string `json:"alias"`
	Delivery     string `json:"delivery"`
}
