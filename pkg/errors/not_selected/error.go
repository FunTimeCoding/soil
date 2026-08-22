package not_selected

func (e *NotSelectedError) Error() string { return e.Message }
