package document_page

func shortHash(hash string) string {
	return hash[:min(len(hash), 16)]
}
