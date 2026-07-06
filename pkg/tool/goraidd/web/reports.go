package web

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/system"
	goraidd "github.com/funtimecoding/soil/pkg/tool/goraidd/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
	"os"
	"strings"
)

func (s *Server) reports(
	w http.ResponseWriter,
	r *http.Request,
) {
	var reports []os.DirEntry

	for _, n := range system.ReadDirectory(s.outputPath) {
		if !n.IsDir() && strings.HasSuffix(
			n.Name(),
			constant.HypertextExtension,
		) {
			reports = append(reports, n)
		}
	}

	if s.view.IsExtendedRequest(r) {
		s.view.RenderFragment(w, reportsTable(reports))

		return
	}

	s.view.RenderPage(
		w,
		goraidd.ReportsTitle,
		goraidd.ReportsPath,
		html.H1(gomponents.Text("Generated Reports")),
		reportsTable(reports),
	)
}
