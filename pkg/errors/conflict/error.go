package conflict

func (e *ConflictError) Error() string { return e.Message }
