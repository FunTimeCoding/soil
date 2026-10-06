package chunk_preview

func New(
	model string,
	window int,
	allowance int,
) *Preview {
	return &Preview{Model: model, Window: window, Allowance: allowance}
}
