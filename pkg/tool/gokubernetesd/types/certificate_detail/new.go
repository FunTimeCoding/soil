package certificate_detail

func New(
	certificate map[string]any,
	latestRequest map[string]any,
	filtered []string,
) *Detail {
	return &Detail{
		Certificate:   certificate,
		LatestRequest: latestRequest,
		Filtered:      filtered,
	}
}
