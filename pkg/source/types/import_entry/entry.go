package import_entry

type Entry struct {
	Path     string
	Declared string
	Local    string
	Module   string
	Standard bool
	Stays    bool
}
