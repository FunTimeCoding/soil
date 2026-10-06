package service

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"

func NewProfileResult(
	always []record.Memory,
	index []record.MemorySummary,
	relevant []record.SearchResult,
	impressions []record.Impression,
	completions []CompletionEntry,
) *ProfileResult {
	result := &ProfileResult{
		Always:      always,
		Index:       index,
		Relevant:    relevant,
		Impressions: impressions,
		Completions: completions,
	}
	result.Text = ProfileText(result)

	return result
}
