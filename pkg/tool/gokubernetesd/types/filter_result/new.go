package filter_result

func New(
	object map[string]any,
	filtered []string,
) *FilterResult {
	return &FilterResult{Object: object, Filtered: filtered}
}
