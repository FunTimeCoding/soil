package model_context

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goladybirdd/types/view"
)

func clicked(
	uid string,
	v *view.View,
) string {
	if v.Landed == "" {
		return join.Space("clicked", uid)
	}

	return join.Space("clicked", uid, "→", v.Landed)
}
