package profile

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"

func NewResult(
	always []record.Memory,
	index []record.MemorySummary,
	relevant []record.SearchResult,
	impressions []record.Impression,
	completions []*Completion,
) *Result {
	result := &Result{
		Always:      always,
		Index:       index,
		Relevant:    relevant,
		Impressions: impressions,
		Completions: completions,
	}
	result.Text = Text(result)

	return result
}
