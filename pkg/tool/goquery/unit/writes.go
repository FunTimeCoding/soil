package unit

func writes() [][]string {
	return [][]string{
		{"add-collection", "alfa", "/bravo"},
		{"delete-collection", "alfa"},
		{"context-add", "alfa", "kilo", "lima"},
		{"context-remove", "alfa", "kilo"},
	}
}
