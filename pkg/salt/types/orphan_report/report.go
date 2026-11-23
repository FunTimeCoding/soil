package orphan_report

type Report struct {
	Summary          *Summary `json:"summary"`
	Unreached        []string `json:"unreached"`
	UnreachedFiles   []string `json:"unreached_files"`
	Possible         []string `json:"possible"`
	Missing          []string `json:"missing"`
	DeadEnvironments []string `json:"dead_environments"`
}
