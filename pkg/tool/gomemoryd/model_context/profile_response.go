package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/service"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store"
)

type profileResponse struct {
	Always      []*convert.SlimMemory       `json:"always"`
	Relevant    []*convert.SlimSearchResult `json:"relevant,omitempty"`
	Index       []store.MemorySummary       `json:"index"`
	Impressions []store.Impression          `json:"impressions,omitempty"`
	Completions []service.CompletionEntry   `json:"completions,omitempty"`
}
