package web

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/constant"
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
			soil.HypertextExtension,
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
		constant.ReportsTitle,
		constant.ReportsPath,
		html.H1(gomponents.Text("Generated Reports")),
		reportsTable(reports),
	)
}
