package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store"
	"slices"
)

func reachText(m *store.Memory) string {
	if slices.Contains(m.Tags, constant.AlwaysTag) {
		return "always - in context every session"
	}

	if m.ParentIdentifier != nil {
		return "leaf - enters context when its parent is followed"
	}

	if slices.Contains(m.Tags, constant.NoIndexTag) {
		return "no-index - reached through search and relations only"
	}

	return "index - description every session, block when the topic matches"
}
