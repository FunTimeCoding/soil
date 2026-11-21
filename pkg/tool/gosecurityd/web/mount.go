package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/funtimecoding/soil/pkg/web/route"
)

func (s *Server) Mount(g *guard.Mux) {
	s.view.Mount(g)
	g.Open(route.Get(webConstant.RootPattern), s.flags)
	g.Open(route.Get(constant.LookupPath), s.lookup)
	g.Open(route.Get(constant.SystemsPath), s.systems)
	g.Open(route.Get(constant.SystemPath), s.systemPage)
	g.Open(route.Get(constant.UpgradesPath), s.upgrades)
	g.Open(route.Get(constant.VulnerabilitiesPath), s.vulnerabilities)
	g.Open(route.Get(constant.RemediationPath), s.remediationList)
	g.Open(route.Get(constant.UpgradeOpenPath), s.upgradeOpen)
	g.Open(route.Get(constant.UpgradeShutPath), s.upgradeShut)
	g.Open(route.Post(constant.FlagPath), s.flagSubmit)
	g.Open(route.Post(constant.MitigatePath), s.mitigate)
	g.Open(route.Get(constant.FlagOpenPath), s.flagOpen)
	g.Open(route.Get(constant.FlagShutPath), s.flagShut)
}
