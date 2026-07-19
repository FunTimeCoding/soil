package environment

func (e *Environment) Set(key string, value string) {
	e.overlay[key] = value
}
