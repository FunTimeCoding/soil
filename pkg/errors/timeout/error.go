package timeout

func (e *TimeoutError) Error() string { return e.Message }
