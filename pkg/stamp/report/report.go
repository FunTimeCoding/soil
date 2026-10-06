package report

type Report struct {
	Name       string `json:"name"`
	Version    string `json:"version"`
	GitHash    string `json:"git_hash"`
	CommitDate string `json:"commit_date"`
	Module     string `json:"module"`
	Dirty      bool   `json:"dirty"`
}
