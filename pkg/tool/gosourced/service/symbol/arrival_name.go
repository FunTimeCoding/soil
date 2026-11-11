package symbol

import (
	"fmt"
	sourceConstant "github.com/funtimecoding/soil/pkg/source/constant"
	"github.com/funtimecoding/soil/pkg/source/imports"
	"github.com/funtimecoding/soil/pkg/source/types/import_entry"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/relocation"
	"go/ast"
	"path/filepath"
	"strings"
)

func arrivalName(
	q *relocation.FileQualification,
	sourcePackagePath string,
	targetPackagePath string,
	targetPackageName string,
) (*relocation.ImportName, map[*ast.Ident]string, string) {
	v, _ := imports.ViewOf(q.Owner, q.File)

	for _, e := range v.Entries {
		if e.Path == targetPackagePath {
			return relocation.NewImportName(e.Local, true), nil, ""
		}
	}

	if !q.SamePackage && !sourceStays(q, sourcePackagePath) {
		for _, e := range v.Entries {
			if e.Path == sourcePackagePath {
				e.Stays = false
			}
		}
	}

	modulePath := ""

	if m := q.Owner.Module; m != nil &&
		strings.HasPrefix(targetPackagePath, m.Path) {
		modulePath = m.Path
	}

	arriving := import_entry.New(
		targetPackagePath,
		targetPackageName,
		targetPackageName,
		modulePath,
		false,
	)
	v.Entries = append(v.Entries, arriving)
	taken := func(name string) bool {
		return v.LocalTaken(name) || otherHeld(v.Entries, arriving, name)
	}
	decided := ""

	for _, d := range imports.Decide(v) {
		if d.Entry == arriving &&
			(d.Kind == sourceConstant.DecisionKeep || d.Kind == sourceConstant.DecisionRename) {
			decided = d.Want
		}
	}

	want := ""
	var locals map[*ast.Ident]string
	refusal := ""

	for _, candidate := range []string{decided, targetPackageName} {
		if candidate == "" || want != "" {
			continue
		}

		if !taken(candidate) {
			want = candidate

			continue
		}

		if otherHeld(v.Entries, arriving, candidate) {
			continue
		}

		shadowing, okay := shadowingLocals(q, candidate)

		if !okay {
			continue
		}

		held := func(name string) bool {
			return name == candidate || otherHeld(v.Entries, nil, name)
		}
		renames, blocker := localRenames(q, shadowing, held)

		if blocker != nil {
			p := q.Owner.Fset.Position(blocker.Pos())
			refusal = fmt.Sprintf(
				"local %s at %s:%d would shadow the import %s and no letter of its name is free",
				blocker.Name(),
				filepath.Base(p.Filename),
				p.Line,
				candidate,
			)

			continue
		}

		want = candidate
		locals = renames
	}

	if want == "" && refusal != "" {
		return nil, nil, refusal
	}

	if want == "" {
		want = imports.ChooseAlias(
			targetPackagePath,
			modulePath,
			targetPackageName,
			taken,
		)
	}

	if want == "" {
		return nil, nil, ""
	}

	result := relocation.NewImportName(want, false)

	if want != targetPackageName {
		result.Alias = want
	}

	return result, locals, ""
}
