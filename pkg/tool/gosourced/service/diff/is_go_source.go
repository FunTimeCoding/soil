package diff

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	go_modConstant "github.com/funtimecoding/soil/pkg/go_mod/constant"
	"strings"
)

func isGoSource(name string) bool {
	return strings.HasSuffix(name, soil.GoExtension) ||
		name == go_modConstant.ModFile
}
