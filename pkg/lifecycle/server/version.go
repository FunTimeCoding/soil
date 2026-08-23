package server

type Version struct {
	Name      string `json:"name"`
	Version   string `json:"version"`
	GitHash   string `json:"git_hash"`
	BuildDate string `json:"build_date"`
}
