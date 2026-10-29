package forbidden

func (e *ForbiddenError) Error() string { return e.Message }
