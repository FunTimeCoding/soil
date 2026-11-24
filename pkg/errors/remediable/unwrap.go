package remediable

func (e *RemediableError) Unwrap() error { return e.cause }
