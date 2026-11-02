package caller_left

func (e *CallerLeft) Unwrap() error {
	return e.cause
}
