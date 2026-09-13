package web

import (
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/constant"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/funtimecoding/soil/pkg/web/palette"
	"github.com/funtimecoding/soil/pkg/web/route"
)

func (s *Server) Mount(g *guard.Mux) {
	g.WithSession(s.require)
	g.Open(route.Get(webConstant.SignInPath), s.signIn)
	g.Open(route.Get(webConstant.CallbackPath), s.callback)
	g.Open(route.Get(webConstant.SignOutPath), s.signOut)
	g.Session(route.Get(webConstant.PalettePath), palette.NewServe(s.registry))
	g.Session(route.Get(webConstant.RootPattern), s.user)
	g.Session(route.Post(constant.CreateUserPath), s.createUserSubmit)
	g.Session(route.Post(constant.DeleteUserPath), s.deleteUserSubmit)
	g.Session(route.Post(constant.SetPasswordPath), s.setPasswordSubmit)
	g.Session(route.Get(constant.GroupPath), s.group)
	g.Session(route.Post(constant.CreateGroupPath), s.createGroupSubmit)
	g.Session(route.Post(constant.DeleteGroupPath), s.deleteGroupSubmit)
	g.Session(route.Post(constant.AddMemberPath), s.addMemberSubmit)
	g.Session(route.Post(constant.RemoveMemberPath), s.removeMemberSubmit)
	g.Open(route.Get(webConstant.FaviconPath), s.favicon)
}
