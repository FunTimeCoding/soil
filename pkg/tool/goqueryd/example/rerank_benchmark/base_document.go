package rerank_benchmark

import "strings"

func baseDocument() string {
	return strings.Repeat(
		"Schema migrations use ALTER TABLE statements with version tracking to evolve the database safely. ",
		20,
	)
}
