package caller_left

func New(cause error) *CallerLeft {
	return &CallerLeft{cause: cause}
}
