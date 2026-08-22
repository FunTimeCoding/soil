package unexpected

func (e *UnexpectedError) Error() string { return e.Message }
