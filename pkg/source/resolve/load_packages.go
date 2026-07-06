package resolve

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/source/build_tag"
	"go/token"
	"golang.org/x/tools/go/packages"
	"strings"
)

func LoadPackages(
	directory string,
	patterns ...string,
) ([]*packages.Package, *token.FileSet, error) {
	set := token.NewFileSet()
	c := &packages.Config{
		Mode:  packages.LoadSyntax | packages.NeedModule,
		Fset:  set,
		Dir:   directory,
		Tests: true,
	}
	tags := build_tag.Discover(directory)

	if len(tags) > 0 {
		c.BuildFlags = []string{
			fmt.Sprintf("-tags=%s", strings.Join(tags, ",")),
		}
	}

	result, e := packages.Load(c, patterns...)

	if e != nil {
		return nil, nil, fmt.Errorf("load packages: %w", e)
	}

	for _, p := range result {
		for _, f := range p.Errors {
			return nil, nil, fmt.Errorf("%s: %s", p.PkgPath, f)
		}
	}

	return result, set, nil
}
