package argument

type NameParameters struct {
	PackagePath string           `json:"package_path"`
	Functions   []*NamedFunction `json:"functions"`
	Type        string           `json:"type"`
	Name        string           `json:"name"`
	Imports     []string         `json:"imports"`
	DryRun      bool             `json:"dry_run"`
}
