package unreachable

func (e *UnreachableError) Error() string { return e.Message }
