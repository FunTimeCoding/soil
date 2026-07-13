package rerank_benchmark

func identicalDocuments(count int) []string {
	base := baseDocument()
	documents := make([]string, count)

	for i := range documents {
		documents[i] = base
	}

	return documents
}
