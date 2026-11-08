package argument

type RenameReceivers struct {
	PackagePath string            `json:"package_path"`
	Receivers   []*ReceiverRename `json:"receivers"`
	DryRun      bool              `json:"dry_run"`
}
