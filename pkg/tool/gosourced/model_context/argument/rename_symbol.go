package argument

type RenameSymbol struct {
	PackagePath string `json:"package_path"`
	OldName     string `json:"old_name"`
	NewName     string `json:"new_name"`
	Receiver    string `json:"receiver"`
	DryRun      bool   `json:"dry_run"`
}
