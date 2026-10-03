package installed

type Binary struct {
	Name    string
	Path    string
	Module  string
	Package string
	Hash    string
	Version string
	Dirty   bool
	Modules map[string]string
}
