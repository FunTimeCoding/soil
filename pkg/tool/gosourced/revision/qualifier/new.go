package qualifier

func New(
	imports map[string]string,
	own map[string]bool,
	path string,
) *Qualifier {
	return &Qualifier{Imports: imports, Own: own, Path: path}
}
