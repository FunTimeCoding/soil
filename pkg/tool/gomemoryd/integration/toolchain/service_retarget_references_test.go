package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/lint/types/move_set"
	"github.com/funtimecoding/soil/pkg/lint/types/path_move"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/unit/service_tester"
	"path/filepath"
	"testing"
)

func TestRetargetReferencesRewritesAfterTheMove(t *testing.T) {
	root := t.TempDir()
	testutil.WriteFile(t, root, "pkg/bravo/service/run.go", "package service\n")
	o := service_tester.NewReferenced(t, root)
	plain := memoryOption("plain", "Read `pkg/alfa/service/run.go` first.")
	based := memoryOption("based", "See `service/run.go`.")
	based.Base = new("pkg/alfa")
	p, e := o.Service.CreateMemory(plain)
	assert.FatalOnError(t, e)
	b, f := o.Service.CreateMemory(based)
	assert.FatalOnError(t, f)
	reports, g := o.Service.RetargetReferences(
		move_set.New(
			[]*path_move.Move{
				path_move.New(
					filepath.Join(root, "pkg/alfa"),
					filepath.Join(root, "pkg/bravo"),
				),
			},
			nil,
		),
		false,
	)
	assert.FatalOnError(t, g)
	assert.Count(t, 2, reports)
	movedPlain, h := o.Service.GetMemory(p.Identifier)
	assert.FatalOnError(t, h)
	assert.String(
		t,
		"Read `pkg/bravo/service/run.go` first.",
		movedPlain.Content,
	)
	movedBased, i := o.Service.GetMemory(b.Identifier)
	assert.FatalOnError(t, i)
	assert.String(t, "See `service/run.go`.", movedBased.Content)
	assert.String(t, "pkg/bravo", movedBased.Metadata[constant.BaseKey])
}

func TestRetargetReferencesDryRunWritesNothing(t *testing.T) {
	root := t.TempDir()
	testutil.WriteFile(t, root, "pkg/bravo/run.go", "package bravo\n")
	o := service_tester.NewReferenced(t, root)
	m, e := o.Service.CreateMemory(
		memoryOption("plain", "Read `pkg/alfa/run.go` first."),
	)
	assert.FatalOnError(t, e)
	reports, f := o.Service.RetargetReferences(
		move_set.New(
			[]*path_move.Move{path_move.New("pkg/alfa", "pkg/bravo")},
			nil,
		),
		true,
	)
	assert.FatalOnError(t, f)
	assert.Count(t, 1, reports)
	assert.String(t, "pkg/bravo/run.go", reports[0].Changes[0].After)
	unchanged, g := o.Service.GetMemory(m.Identifier)
	assert.FatalOnError(t, g)
	assert.String(t, "Read `pkg/alfa/run.go` first.", unchanged.Content)
}
