package lint

import (
	gitConstant "github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/funtimecoding/soil/pkg/lint/option"
	"github.com/funtimecoding/soil/pkg/lint/pointer/resolver"
	"github.com/funtimecoding/soil/pkg/lint/repository"
	"github.com/funtimecoding/soil/pkg/system/writer"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

func pointerGlobal(
	repo *repository.Repository,
	o *option.Lint,
	base *resolver.Resolver,
) string {
	var b strings.Builder

	for _, part := range [][]string{
		base.Roots,
		repo.ImplicitBases,
		o.Registries,
		repo.Modules,
	} {
		writer.Print(&b, "%q\n", slices.Sorted(slices.Values(part)))
	}

	roots := []string{repo.Root}

	for _, s := range repo.Siblings {
		roots = append(roots, repo.Absolute(s))
	}

	for _, root := range roots {
		i, e := os.Stat(
			filepath.Join(root, filepath.FromSlash(gitConstant.ExcludeFile)),
		)

		if e == nil {
			writer.Print(&b, "%d %d\n", i.Size(), i.ModTime().UnixNano())
		}
	}

	return b.String()
}
