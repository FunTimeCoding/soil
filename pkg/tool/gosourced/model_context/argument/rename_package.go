package argument

type RenamePackage struct {
	PackagePath string `json:"package_path"`
	NewName     string `json:"new_name"`
	DryRun      bool   `json:"dry_run"`
}
