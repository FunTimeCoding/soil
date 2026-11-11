package argument

type MoveMethods struct {
	PackagePath       string   `json:"package_path"`
	Type              string   `json:"type"`
	Methods           []string `json:"methods"`
	TargetPackagePath string   `json:"target_package_path"`
	TargetType        string   `json:"target_type"`
	TargetFile        string   `json:"target_file"`
	DryRun            bool     `json:"dry_run"`
}
