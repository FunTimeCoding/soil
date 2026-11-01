package component

import (
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
)

func PlaceLocator(
	kind string,
	name string,
) string {
	return join.Empty(
		constant.PlacesPath,
		stringsConstant.Slash,
		segment(kind),
		stringsConstant.Slash,
		name,
	)
}
