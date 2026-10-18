package unavailable

func (e *UnavailableError) Error() string { return e.Message }
