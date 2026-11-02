package decoration

import (
	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
	"golang.org/x/tools/go/packages"
)

func NewSet() *Set {
	return &Set{
		Decorators:   make(map[*packages.Package]*decorator.Decorator),
		Files:        make(map[string]*dst.File),
		PackagePaths: make(map[*dst.File]string),
		Aliases:      make(map[*dst.File]map[string]string),
	}
}
