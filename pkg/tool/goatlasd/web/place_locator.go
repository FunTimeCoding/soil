package web

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	atlas "github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
)

func placeLocator(
	kind string,
	name string,
) string {
	return join.Empty(
		atlas.PlacesPath,
		constant.Slash,
		placeSegment(kind),
		constant.Slash,
		name,
	)
}
