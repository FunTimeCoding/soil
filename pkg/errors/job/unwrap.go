package job

func (e *JobError) Unwrap() error {
	return e.Wrapped
}
