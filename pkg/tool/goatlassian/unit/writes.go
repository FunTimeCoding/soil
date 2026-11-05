package unit

func writes() [][]string {
	return [][]string{
		{"delete-comment", "ALFA-1", "12"},
		{"delete-link", "12"},
		{"delete-page", "12"},
		{"link-issues", "ALFA-1", "ALFA-2"},
		{"update-comment", "ALFA-1", "12", "bravo"},
	}
}
