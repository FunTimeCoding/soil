package oversize

func New(
	model string,
	window int,
	allowance int,
	embedAllowance int,
) *Report {
	return &Report{
		Model:          model,
		Window:         window,
		Allowance:      allowance,
		EmbedAllowance: embedAllowance,
	}
}
