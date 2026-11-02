package argument

type RetargetReferences struct {
	Paths   []*PathMove   `json:"paths"`
	Symbols []*SymbolMove `json:"symbols"`
	DryRun  bool          `json:"dry_run"`
}
