package mechanic

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/funtimecoding/soil/pkg/web/route"
)

func (s *Server) Mount(g *guard.Mux) {
	g.OpenMount(route.Get(constant.LivePath), s.event())
	g.Open(route.Get(constant.RootPath), s.swap)
	g.Open(route.Get(constant.TriggerPath), s.trigger)
	g.Open(route.Get(constant.NotifyPath), s.notification)
	g.Open(route.Get(constant.StreamPath), s.stream)
	g.Open(route.Get(constant.CounterPath), s.counter)
	g.Open(route.Post(constant.CounterPath), s.counterIncrement)
	g.Open(route.Post(constant.RowPath), s.rowReplace)
	g.Open(route.Get(constant.EchoPath), s.echo)
	g.Open(route.Post(constant.FailPath), s.fail)
	g.Open(route.Post(constant.OutOfBandPath), s.outOfBand)
	g.Open(route.Post(constant.SubscribePath), s.subscribe)
	g.Open(route.Get(constant.ExtraPath), s.extra)
	g.Open(route.Post(constant.RemovePath), s.remove)
	g.Open(route.Post(constant.AppendPath), s.appendItem)
	g.Open(route.Post(constant.QuietPath), s.quiet)
	g.Open(route.Get(constant.BranchPath), s.branch)
	g.Open(route.Get(constant.PollPath), s.poll)
	g.Open(route.Post(constant.FirePath), s.fire)
	g.Open(route.Post(constant.SlowPath), s.slow)
	g.Open(route.Post(constant.RedirectPath), s.redirect)
	g.Open(route.Post(constant.RefreshPath), s.refresh)
}
