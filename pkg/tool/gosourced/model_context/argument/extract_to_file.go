package argument

type ExtractToFile struct {
	File     string `json:"file"`
	Function string `json:"function"`
	DryRun   bool   `json:"dry_run"`
}
