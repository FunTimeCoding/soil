package store

type RelationOverview struct {
	SourceIdentifier int64
	SourceName       string
	SourceScope      string
	TargetIdentifier int64
	TargetName       string
	TargetScope      string
	Type             string
	CreatedAt        string
}
