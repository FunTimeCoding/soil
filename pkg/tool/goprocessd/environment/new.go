package environment

func New(base []string) *Environment {
	return &Environment{base: base, overlay: make(map[string]string)}
}
