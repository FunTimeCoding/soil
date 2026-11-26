package browser_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/goladybirdd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goladybirdd/navigation"
	"github.com/funtimecoding/soil/pkg/tool/goladybirdd/option"
	"github.com/funtimecoding/soil/pkg/tool/goladybirdd/service"
)

func (b *Tester) Service() *service.Service {
	b.t.Helper()

	if b.service == nil {
		o := option.New()
		o.Profile = ""
		w := navigation.New(constant.NavigationWindow, constant.NavigationLimit)
		c := o.Client(w.Observe)
		b.t.Cleanup(c.Close)
		assert.FatalOnError(b.t, c.Subscribe(navigation.Events()...))
		g := peer.New(
			constant.LadybirdPeer,
			c.Address(),
			c,
			logger.New(b.t.Context()),
			memory.New(),
		)
		b.t.Cleanup(g.Close)
		b.service = service.New(c, g, w, b.t.TempDir())
	}

	return b.service
}
