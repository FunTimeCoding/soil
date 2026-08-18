package argument

type RemoveImport struct {
	File       string `json:"file"`
	ImportPath string `json:"import_path"`
	DryRun     bool   `json:"dry_run"`
}
