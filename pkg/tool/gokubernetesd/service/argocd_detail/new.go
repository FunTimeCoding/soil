package argocd_detail

func New(
	application map[string]any,
	filtered []string,
) *Detail {
	return &Detail{Application: application, Filtered: filtered}
}
