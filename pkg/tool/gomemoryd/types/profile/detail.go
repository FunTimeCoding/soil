package profile

type Detail struct {
	Budget             int `json:"budget"`
	AlwaysBytes        int `json:"always_bytes"`
	IndexBytes         int `json:"index_bytes"`
	IndexTrimmed       int `json:"index_trimmed"`
	CompletionBytes    int `json:"completion_bytes"`
	CompletionsTrimmed int `json:"completions_trimmed"`
	ImpressionBytes    int `json:"impression_bytes"`
	ImpressionsTrimmed int `json:"impressions_trimmed"`
	RelevantBytes      int `json:"relevant_bytes"`
	RelevantTrimmed    int `json:"relevant_trimmed"`
	TotalBytes         int `json:"total_bytes"`
}
