package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/web/coverage_page"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
	"sort"
)

func (s *Server) coveragePage(
	w http.ResponseWriter,
	_ *http.Request,
) {
	servers := s.service.Coverage()
	sort.SliceStable(
		servers,
		func(i, j int) bool {
			if servers[i].CallsRecent != servers[j].CallsRecent {
				return servers[i].CallsRecent < servers[j].CallsRecent
			}

			return servers[i].CallsTotal < servers[j].CallsTotal
		},
	)
	s.view.RenderPage(
		w,
		constant.CoverageTitle,
		constant.CoveragePath,
		html.H3(gomponents.Text(constant.CoverageTitle)),
		coverage_page.Region(servers),
	)
}
