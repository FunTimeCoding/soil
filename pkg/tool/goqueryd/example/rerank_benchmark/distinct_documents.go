package rerank_benchmark

import "fmt"

func distinctDocuments(count int) []string {
	base := baseDocument()
	documents := make([]string, count)

	for i := range documents {
		documents[i] = fmt.Sprintf("Variant %d: %s", i, base)
	}

	return documents
}
