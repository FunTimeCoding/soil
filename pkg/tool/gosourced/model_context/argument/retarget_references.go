package argument

type RetargetReferences struct {
	Paths  []*PathMove `json:"paths"`
	DryRun bool        `json:"dry_run"`
}
