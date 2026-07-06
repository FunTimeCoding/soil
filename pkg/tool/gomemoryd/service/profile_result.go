package service

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/store"

type ProfileResult struct {
	Always      []store.Memory       `json:"always"`
	Relevant    []store.SearchResult `json:"relevant,omitempty"`
	Index       []store.MemorySummary `json:"index"`
	Impressions []store.Impression   `json:"impressions,omitempty"`
	Completions []CompletionEntry    `json:"completions,omitempty"`
}
