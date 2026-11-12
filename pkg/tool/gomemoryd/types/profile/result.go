package profile

import (
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/types/profile/completion"
)

type Result struct {
	Always      []record.Memory          `json:"always"`
	Relevant    []record.SearchResult    `json:"relevant,omitempty"`
	Index       []record.MemorySummary   `json:"index"`
	Impressions []record.Impression      `json:"impressions,omitempty"`
	Completions []*completion.Completion `json:"completions,omitempty"`
	Notice      string                   `json:"notice,omitempty"`
	Text        string                   `json:"-"`
}
