package chunk_preview

func New(
	model string,
	window int,
	tokenLimit int,
	embedTokenLimit int,
) *Preview {
	return &Preview{
		Model:           model,
		Window:          window,
		TokenLimit:      tokenLimit,
		EmbedTokenLimit: embedTokenLimit,
	}
}
