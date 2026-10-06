package mock_reranker

import "strings"

func (r *Reranker) Count(text string) int {
	return len(strings.Fields(text))
}
