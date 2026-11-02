package load_failure

func (e *LoadFailureError) Unwrap() error { return e.Cause }
