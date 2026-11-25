package response

type Subscription struct {
	Reference    string `json:"reference"`
	Alias        string `json:"alias,omitempty"`
	Instance     string `json:"instance"`
	Project      int64  `json:"project"`
	MergeRequest int64  `json:"merge_request"`
	Delivery     string `json:"delivery"`
}
