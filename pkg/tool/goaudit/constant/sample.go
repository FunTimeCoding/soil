package constant

const (
	FixtureFrontendNewSample   = "package web\n\nimport \"github.com/funtimecoding/soil/pkg/web/layout\"\n\nfunc New() *Server {\n\treturn &Server{view: view.New(layout.New(constant.Identity, theme.Straw, favicon).WithStyle(constant.Style).WithItems(a, b))}\n}\n"
	FixtureFrontendPlainSample = "package web\n\nimport \"github.com/funtimecoding/soil/pkg/web/layout\"\n\nfunc New() *Server {\n\treturn &Server{view: view.New(layout.New(constant.Identity, \"\", nil))}\n}\n"
	FixtureFrontendSignOnMount = "package web\n\nfunc (s *Server) Mount(g *guard.Mux) {\n\ts.authorization.Mount(g)\n}\n"
	FixtureFrontendMountSample = "package web\n\nfunc (s *Server) Mount(g *guard.Mux) {\n\ts.authorization.Mount(g)\n\ts.view.Mount(g)\n}\n"
	FixtureClosedRun           = "package gotestd\n\nfunc Run() {\n\td := daemon.New()\n\tdefer d.Close()\n\tlifecycle.New()\n}\n"
	FixtureUnclosedRun         = "package gotestd\n\nfunc Run() {\n\td := daemon.New()\n\tlifecycle.New()\n}\n"
	FixtureClockRun            = "package gotestd\n\nfunc Run() {\n\td := daemon.New()\n\tdefer d.Close()\n\tlifecycle.New().WithWorker(worker.New())\n}\n"
	FixtureClientRun           = "package gotestd\n\nfunc Run() {\n\tc := alfa.NewClient()\n\td := daemon.New(c)\n\tdefer d.Close()\n\tlifecycle.New()\n}\n"
	FixtureMapperService       = "package book\n\nimport \"gorm.io/gorm\"\n\ntype Service struct{ mapper *gorm.DB }\n"
)
