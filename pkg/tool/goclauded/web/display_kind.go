package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/context_load"
)

func displayKind(entry *context_load.Load) string {
	if entry.Tier != "" {
		return constant.LoadKindProfile
	}

	return entry.Kind
}
