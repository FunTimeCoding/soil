package argument

type ApplyPattern struct {
	PackagePath string `json:"package_path"`
	Symbol      string `json:"symbol"`
	Receiver    string `json:"receiver"`
	Pattern     string `json:"pattern"`
	Replacement string `json:"replacement"`
	Partial     bool   `json:"partial"`
	DryRun      bool   `json:"dry_run"`
}
