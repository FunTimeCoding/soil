package argument

type ReplyToMergeRequestDiscussion struct {
	Project      string `json:"project"`
	MergeRequest int64  `json:"merge_request"`
	Discussion   string `json:"discussion"`
	Body         string `json:"body"`
	Callsign     string `json:"callsign"`
	Resolve      bool   `json:"resolve"`
}
