package orphan_report

type Summary struct {
	Reached        int `json:"reached"`
	Possible       int `json:"possible"`
	Unreached      int `json:"unreached"`
	Missing        int `json:"missing"`
	UnreachedFiles int `json:"unreached_files"`
}
