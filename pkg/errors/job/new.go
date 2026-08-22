package job

func New(
	identifier any,
	kind string,
	wrapped error,
) *JobError {
	return &JobError{
		Identifier: identifier,
		Kind:       kind,
		Wrapped:    wrapped,
	}
}
