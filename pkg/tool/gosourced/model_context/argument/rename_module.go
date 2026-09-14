package argument

type RenameModule struct {
	ModulePath    string `json:"module_path"`
	NewModulePath string `json:"new_module_path"`
	Force         bool   `json:"force"`
	DryRun        bool   `json:"dry_run"`
}
