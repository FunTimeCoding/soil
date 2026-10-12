package service

import (
	"github.com/funtimecoding/soil/pkg/source/constant"
	"github.com/funtimecoding/soil/pkg/source/imports"
	"github.com/funtimecoding/soil/pkg/source/types/import_entry"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/relocation"
	"strings"
)

func arrivalName(
	q *relocation.FileQualification,
	sourcePackagePath string,
	targetPackagePath string,
	targetPackageName string,
) *relocation.ImportName {
	v, _ := imports.ViewOf(q.Owner, q.File)

	for _, e := range v.Entries {
		if e.Path == targetPackagePath {
			return relocation.NewImportName(e.Local, true)
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
		return v.LocalTaken(name) || heldByOther(v.Entries, arriving, name)
	}
	want := ""

	for _, d := range imports.Decide(v) {
		if d.Entry == arriving &&
			(d.Kind == constant.DecisionKeep || d.Kind == constant.DecisionRename) {
			want = d.Want
		}
	}

	if want == "" || taken(want) {
		want = targetPackageName

		if taken(want) {
			want = imports.ChooseAlias(
				targetPackagePath,
				modulePath,
				targetPackageName,
				taken,
			)
		}
	}

	if want == "" {
		return nil
	}

	result := relocation.NewImportName(want, false)

	if want != targetPackageName {
		result.Alias = want
	}

	return result
}
