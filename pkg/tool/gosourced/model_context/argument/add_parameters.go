package argument

type AddParameters struct {
	Functions []*AdditionFunction `json:"functions"`
	Imports   []string            `json:"imports"`
	DryRun    bool                `json:"dry_run"`
}
