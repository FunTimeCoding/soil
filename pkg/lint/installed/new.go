package installed

func New(
	name string,
	path string,
) *Binary {
	return &Binary{Name: name, Path: path}
}
