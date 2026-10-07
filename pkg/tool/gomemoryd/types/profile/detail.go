package profile

type Detail struct {
	Budget             int `json:"budget"`
	AlwaysTokens       int `json:"always_tokens"`
	IndexTokens        int `json:"index_tokens"`
	IndexTrimmed       int `json:"index_trimmed"`
	CompletionTokens   int `json:"completion_tokens"`
	CompletionsTrimmed int `json:"completions_trimmed"`
	ImpressionTokens   int `json:"impression_tokens"`
	ImpressionsTrimmed int `json:"impressions_trimmed"`
	RelevantTokens     int `json:"relevant_tokens"`
	RelevantTrimmed    int `json:"relevant_trimmed"`
	TotalTokens        int `json:"total_tokens"`
}
