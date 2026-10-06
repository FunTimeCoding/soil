package record

type SourceTypeTag struct {
	Collection string `json:"collection"`
	PathPrefix string `json:"path_prefix"`
	SourceType string `json:"source_type"`
}
