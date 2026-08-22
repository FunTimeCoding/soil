package job

type JobError struct {
	Identifier any
	Kind       string
	Detail     map[string]any
	Wrapped    error
}
