package module_graph

func reaches(
	imports []string,
	reach map[string]bool,
) bool {
	for _, i := range imports {
		if reach[i] {
			return true
		}
	}

	return false
}
