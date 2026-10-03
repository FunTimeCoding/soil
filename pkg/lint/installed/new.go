package installed

func New(
	name string,
	path string,
	mainPackage string,
) *Binary {
	return &Binary{
		Name:    name,
		Path:    path,
		Package: mainPackage,
		Modules: make(map[string]string),
	}
}
