package argument

type RenameSymbols struct {
	PackagePath string          `json:"package_path"`
	Renames     []*SymbolRename `json:"renames"`
	DryRun      bool            `json:"dry_run"`
}
