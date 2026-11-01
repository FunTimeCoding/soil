package web

import (
	"github.com/funtimecoding/soil/pkg/raid/model/fight"
	"github.com/funtimecoding/soil/pkg/strings/join"
	timeConstant "github.com/funtimecoding/soil/pkg/time/constant"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/web/log_page"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
	"time"
)

func (s *Server) logs(
	w http.ResponseWriter,
	r *http.Request,
) {
	q := r.URL.Query()
	all := s.service.Fights()
	filtered := q.Has(constant.StartParameter) || q.Has(constant.EndParameter)
	start := q.Get(constant.StartParameter)
	end := q.Get(constant.EndParameter)
	var fights []fight.Fight
	offset := 0

	if filtered {
		fights = within(all, start, end)
	} else {
		today := time.Now().Format(timeConstant.DateYear)
		start = join.Empty(today, "T00:00")
		end = join.Empty(today, "T23:59")
		fights = all[:min(len(all), constant.PageSize)]
		offset = min(
			parseIntParameter(r, constant.OffsetParameter, 0),
			len(fights),
		)
	}

	total := len(fights)
	page := fights

	if !filtered {
		page = fights[offset:min(offset+constant.PageSize, total)]
	}

	listing := log_page.Listing(page, offset, total, start, end, !filtered)

	if s.view.IsExtendedRequest(r) {
		s.view.RenderFragment(w, listing)

		return
	}

	s.view.RenderPage(
		w,
		constant.LogsTitle,
		constant.LogsPath,
		html.H1(gomponents.Textf("Combat Logs (%d)", total)),
		log_page.Filter(start, end),
		log_page.Form(listing),
	)
}
