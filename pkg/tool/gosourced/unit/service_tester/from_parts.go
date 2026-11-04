package service_tester

import (
	soilFace "github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/source/index"
	"github.com/funtimecoding/soil/pkg/source/inventory"
	gomemorydConstant "github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/face"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/commit"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/diff"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/file"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/module"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/packages"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/pattern"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/reference"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/symbol"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/workspace"
	"testing"
)

func fromParts(
	t *testing.T,
	i *inventory.Inventory,
	m face.Client,
	r soilFace.Reporter,
) *Services {
	t.Helper()
	l := logger.New(t.Context())
	g := peer.New(gomemorydConstant.Identity.Name(), m.Address(), m, l, r)
	t.Cleanup(g.Close)
	w := workspace.New(index.DefaultDirectory())
	c := commit.New(i, w, m, g, r)
	u := module.New(i, c)

	return &Services{
		Workspace: w,
		Commit:    c,
		Diff:      diff.New(i, u),
		File:      file.New(c),
		Module:    u,
		Packages:  packages.New(w, c),
		Pattern:   pattern.New(w, c),
		Reference: reference.New(w, c),
		Symbol:    symbol.New(w, c),
	}
}
