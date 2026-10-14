package unclosed_resource

func allKnown(
	known map[string]bool,
	callees []string,
) bool {
	for _, c := range callees {
		if !known[c] {
			return false
		}
	}

	return true
}
