package oversize

func New(
	model string,
	window int,
	allowance int,
) *Report {
	return &Report{Model: model, Window: window, Allowance: allowance}
}
