package gosecret

type MigrateResult struct {
	Path    string
	Keys    []string
	Extra   []string
	Removed string
}
