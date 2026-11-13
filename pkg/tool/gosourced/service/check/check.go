package check

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/source/inventory/module"
	"github.com/funtimecoding/soil/pkg/tool/gocheck"
	"github.com/funtimecoding/soil/pkg/tool/gocheck/option"
	"github.com/funtimecoding/soil/pkg/tool/gocheck/types/check_result"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/turn"
	"slices"
)

func (s *Service) Check(
	x context.Context,
	o *option.Check,
) (*check_result.Result, error) {
	if !slices.ContainsFunc(
		s.inventory.Modules,
		func(m module.Module) bool {
			return m.Directory == o.Root
		},
	) {
		return nil, not_found.New(constant.RootKind, o.Root)
	}
	g, e := s.turns.Take(x, turn.Session(x))

	if e != nil {
		return nil, e
	}

	defer g.Release()
	s.resolved.Begin(s.trees.Refresh())
	s.documents.Rotate()
	o.Index = s.indexDirectory
	o.Inventory = s.inventory
	o.Trees = s.trees
	o.Documents = s.documents
	o.Resolved = s.resolved
	o.Workspace = s.workspace.Workspace

	if o.Fix {
		c, f := s.commit.Turn(x)

		if f != nil {
			return nil, f
		}

		defer c.Release()
		o.CommitLock = s.commit.CommitLock()
	}

	r, f := gocheck.Run(o)

	if f != nil {
		return nil, f
	}

	r.Stale = s.stale()

	return r, nil
}
