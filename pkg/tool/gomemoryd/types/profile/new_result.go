package profile

import (
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/types/profile/completion"
)

func NewResult(
	always []record.Memory,
	index []record.MemorySummary,
	relevant []record.SearchResult,
	impressions []record.Impression,
	completions []*completion.Completion,
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
