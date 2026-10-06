package argument

type RemoveParameters struct {
	Functions []*ParameterFunction `json:"functions"`
	DryRun    bool                 `json:"dry_run"`
}
