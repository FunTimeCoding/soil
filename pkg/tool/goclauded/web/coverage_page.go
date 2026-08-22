package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
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
	content := []gomponents.Node{
		html.H3(gomponents.Text(constant.CoverageTitle)),
		coverageTable(servers),
	}

	for _, e := range servers {
		if len(e.Tools) == 0 {
			continue
		}

		content = append(content, coverageDetail(e))
	}

	s.view.RenderPage(
		w,
		constant.CoverageTitle,
		constant.CoveragePath,
		content...)
}
