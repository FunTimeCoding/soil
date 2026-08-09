package gomemory

func typeLabel(relationType *string) string {
	if relationType == nil || *relationType == "" {
		return "-"
	}

	return *relationType
}
