package fact

func NewInterface(
	path string,
	name string,
	methods map[string]string,
) *Interface {
	return &Interface{Package: path, Name: name, Methods: methods}
}
