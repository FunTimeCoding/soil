package ambiguous

func (e *AmbiguousError) Error() string { return e.Message }
