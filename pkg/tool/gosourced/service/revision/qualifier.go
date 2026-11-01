package revision

type qualifier struct {
	imports map[string]string
	own     map[string]bool
	path    string
}
