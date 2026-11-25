package response

type Watched struct {
	Reference      string `json:"reference"`
	Alias          string `json:"alias,omitempty"`
	Title          string `json:"title"`
	State          string `json:"state"`
	Draft          bool   `json:"draft"`
	MergeStatus    string `json:"merge_status"`
	PipelineStatus string `json:"pipeline_status,omitempty"`
	Delivery       string `json:"delivery"`
}
