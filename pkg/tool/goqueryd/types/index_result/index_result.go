package index_result

type IndexResult struct {
	Collection string `json:"collection"`
	Indexed    int    `json:"indexed"`
	Updated    int    `json:"updated"`
	Unchanged  int    `json:"unchanged"`
	Removed    int    `json:"removed"`
}
