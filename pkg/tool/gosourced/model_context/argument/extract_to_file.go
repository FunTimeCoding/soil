package argument

type ExtractToFile struct {
	File   string `json:"file"`
	Symbol string `json:"symbol"`
	DryRun bool   `json:"dry_run"`
}
