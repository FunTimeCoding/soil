package dropped

func (e *DroppedError) Error() string { return e.Message }
