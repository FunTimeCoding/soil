package service

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/store"

func NewProfileResult(
	always []store.Memory,
	index []store.MemorySummary,
	relevant []store.SearchResult,
	impressions []store.Impression,
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
