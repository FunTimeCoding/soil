package pointer_cache

import (
	gitConstant "github.com/funtimecoding/soil/pkg/git/constant"
	goModConstant "github.com/funtimecoding/soil/pkg/go_mod/constant"
	"path"
)

func globalChange(p string) bool {
	name := path.Base(p)

	return name == gitConstant.IgnoreFile || name == goModConstant.ModFile
}
