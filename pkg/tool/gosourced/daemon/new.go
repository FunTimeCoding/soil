package daemon

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	gomemorydConstant "github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/option"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/census"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/check"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/commit"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/diff"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/file"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/module"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/packages"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/pattern"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/reference"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/symbol"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/workspace"
)

func New(
	o *option.Source,
	l *logger.Logger,
	r face.Reporter,
) *Daemon {
	m := o.Memory()
	g := peer.New(gomemorydConstant.Identity.Name(), m.Address(), m, l, r)
	w := workspace.New(o.IndexDirectory)
	c := commit.New(o.Inventory, w, m, g, r)
	u := module.New(o.Inventory, c)

	return &Daemon{
		memory:    m,
		guard:     g,
		census:    census.New(w),
		check:     check.New(o.Inventory, c, w, o.IndexDirectory, r),
		diff:      diff.New(o.Inventory, u),
		file:      file.New(c),
		module:    u,
		packages:  packages.New(w, c),
		pattern:   pattern.New(w, c),
		reference: reference.New(w, c),
		symbol:    symbol.New(w, c),
		workspace: w,
	}
}
