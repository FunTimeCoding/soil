package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store"
)

type slimMemoryWithHistory struct {
	convert.SlimMemory
	Related []store.Related `json:"related,omitempty"`
	History []store.Version `json:"history,omitempty"`
}
