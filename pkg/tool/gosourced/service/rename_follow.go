package service

import (
	"fmt"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/strings/camel"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/sink"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/packages"
	"path/filepath"
	"strings"
)

func renameFollow(
	set *token.FileSet,
	all []*packages.Package,
	declaration types.Object,
	newName string,
	out *sink.Sink,
	r *output.Results,
) {
	if !namesFile(declaration) {
		return
	}

	path := set.Position(declaration.Pos()).Filename
	_, file := findOwningFile(all, declaration.Pos())

	if file == nil || strings.HasSuffix(path, soil.TestSuffix) {
		return
	}

	if sole, okay := soleIdentity(file); !okay || sole != declaration.Name() {
		return
	}

	stem := strings.TrimSuffix(filepath.Base(path), soil.GoExtension)
	suffix := platformSuffix(stem)

	if strings.TrimSuffix(stem, suffix) != camel.ToSnake(declaration.Name()) {
		return
	}

	target := filepath.Join(
		filepath.Dir(path),
		join.Empty(camel.ToSnake(newName), suffix, soil.GoExtension),
	)

	if target == path {
		return
	}

	if system.FileExists(target) {
		r.AddConcern(
			concern.NewFile(
				constant.ConcernRenamed,
				fmt.Sprintf(
					"%s was not renamed: %s already exists",
					filepath.Base(path),
					filepath.Base(target),
				),
				path,
				true,
			),
		)

		return
	}

	out.Rename(path, target)
	r.AddConcern(
		concern.NewFile(
			constant.ConcernRenamed,
			fmt.Sprintf(
				"%s → %s",
				filepath.Base(path),
				filepath.Base(target),
			),
			path,
			true,
		),
	)
}
