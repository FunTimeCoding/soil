package argument

type ChangeVisibility struct {
	Symbol      string `json:"symbol"`
	PackagePath string `json:"package_path"`
	Receiver    string `json:"receiver"`
	DryRun      bool   `json:"dry_run"`
}
