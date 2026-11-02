package decoration

import (
	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
	"golang.org/x/tools/go/packages"
)

type Set struct {
	Decorators   map[*packages.Package]*decorator.Decorator
	Files        map[string]*dst.File
	PackagePaths map[*dst.File]string
	Aliases      map[*dst.File]map[string]string
}
