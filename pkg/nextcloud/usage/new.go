package usage

func New(
	files int64,
	shares int64,
) *Usage {
	return &Usage{Files: files, Shares: shares}
}
