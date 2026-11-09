package argument

type RestoreBackup struct {
	Backup string `json:"backup"`
	Target string `json:"target"`
}
