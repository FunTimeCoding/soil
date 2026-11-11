package result

type CensusUnloaded struct {
	Package string `json:"package"`
	Cause   string `json:"cause"`
	Lost    string `json:"lost"`
}
