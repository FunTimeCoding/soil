package face

type ContextProvider interface {
	ErrorContext() (string, map[string]any)
}
