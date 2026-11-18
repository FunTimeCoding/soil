package detached

func (e *DetachedError) Error() string { return e.Message }
