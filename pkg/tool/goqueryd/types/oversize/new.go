package oversize

func New(
	model string,
	window int,
	tokenLimit int,
	embedTokenLimit int,
) *Report {
	return &Report{
		Model:           model,
		Window:          window,
		TokenLimit:      tokenLimit,
		EmbedTokenLimit: embedTokenLimit,
	}
}
