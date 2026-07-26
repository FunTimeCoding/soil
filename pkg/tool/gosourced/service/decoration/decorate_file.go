package decoration

import (
	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
	"go/ast"
	"go/token"
	"golang.org/x/tools/go/packages"
)

func (s *Set) DecorateFile(
	set *token.FileSet,
	p *packages.Package,
	file *ast.File,
) (*dst.File, error) {
	filename := set.Position(file.Pos()).Filename

	if result, exists := s.Files[filename]; exists {
		return result, nil
	}

	dec, exists := s.Decorators[p]

	if !exists {
		dec = decorator.NewDecoratorFromPackage(p)
		s.Decorators[p] = dec
	}

	result, e := dec.DecorateFile(file)

	if e != nil {
		return nil, e
	}

	s.Files[filename] = result
	s.PackagePaths[result] = p.PkgPath

	return result, nil
}
