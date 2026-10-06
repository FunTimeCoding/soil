package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"
)

type slimMemoryWithHistory struct {
	convert.SlimMemory
	Related []record.Related `json:"related,omitempty"`
	History []record.Version `json:"history,omitempty"`
}
