package unclosed_resource

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"go/types"
	"golang.org/x/tools/go/packages"
	"strings"
)

func inModule(
	p *packages.Package,
	t *types.Package,
) bool {
	if t == nil {
		return false
	}

	if p.Module == nil {
		return t.Path() == p.PkgPath
	}

	return t.Path() == p.Module.Path ||
		strings.HasPrefix(t.Path(), join.Empty(p.Module.Path, "/"))
}
