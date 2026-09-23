package constant

const (
	FixtureFrontendNewSample   = "package web\n\nimport \"github.com/funtimecoding/soil/pkg/web/layout\"\n\nfunc New() *Server {\n\treturn &Server{view: view.New(layout.New(constant.Identity).WithTheme(theme.Straw).WithStyle(constant.Style).WithCommandPalette(\"/palette\").WithItems(a, b))}\n}\n"
	FixtureFrontendMountSample = "package web\n\nfunc (s *Server) Mount(m *http.ServeMux) {\n\tm.HandleFunc(\"GET /palette\", p)\n\tm.HandleFunc(\"GET /favicon.ico\", s.favicon)\n}\n"
)
