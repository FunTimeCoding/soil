package chunk_preview

func New(
	model string,
	window int,
	allowance int,
	embedAllowance int,
) *Preview {
	return &Preview{
		Model:          model,
		Window:         window,
		Allowance:      allowance,
		EmbedAllowance: embedAllowance,
	}
}
