package service

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"

type ProfileResult struct {
	Always      []record.Memory        `json:"always"`
	Relevant    []record.SearchResult  `json:"relevant,omitempty"`
	Index       []record.MemorySummary `json:"index"`
	Impressions []record.Impression    `json:"impressions,omitempty"`
	Completions []CompletionEntry      `json:"completions,omitempty"`
	Text        string                 `json:"-"`
}
