package argument

type MatchPattern struct {
	PackagePath string `json:"package_path"`
	Symbol      string `json:"symbol"`
	Receiver    string `json:"receiver"`
	Pattern     string `json:"pattern"`
}
