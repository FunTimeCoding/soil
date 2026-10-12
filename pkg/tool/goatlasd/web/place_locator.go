package web

import (
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
)

func placeLocator(
	kind string,
	name string,
) string {
	return join.Empty(
		constant.PlacesPath,
		stringsConstant.Slash,
		placeSegment(kind),
		stringsConstant.Slash,
		name,
	)
}
