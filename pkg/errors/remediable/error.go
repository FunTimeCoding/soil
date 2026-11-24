package remediable

func (e *RemediableError) Error() string { return e.Message }
