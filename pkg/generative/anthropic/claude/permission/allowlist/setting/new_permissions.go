package setting

func NewPermissions(allow []string) *Permissions {
	return &Permissions{Allow: allow}
}
