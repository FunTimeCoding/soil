package web

import (
	"github.com/funtimecoding/soil/pkg/raid"
	"github.com/funtimecoding/soil/pkg/strings/join"
	timeLibrary "github.com/funtimecoding/soil/pkg/time"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
	"time"
)

func (s *Server) logs(
	w http.ResponseWriter,
	r *http.Request,
) {
	all := s.store.Fights()
	filtered := r.URL.Query().Has("start") || r.URL.Query().Has("end")
	startValue := r.URL.Query().Get("start")
	endValue := r.URL.Query().Get("end")

	if !filtered {
		startValue = join.Empty(
			time.Now().Format(timeLibrary.DateYear),
			"T00:00",
		)
		endValue = join.Empty(
			time.Now().Format(timeLibrary.DateYear),
			"T23:59",
		)
	}

	var fights []raid.Fight

	if filtered {
		for _, f := range all {
			if startValue != "" {
				start, e := time.Parse("2006-01-02T15:04", startValue)

				if e == nil && f.Timestamp.Before(start) {
					continue
				}
			}

			if endValue != "" {
				end, e := time.Parse("2006-01-02T15:04", endValue)

				if e == nil && f.Timestamp.After(end) {
					continue
				}
			}

			fights = append(fights, f)
		}
	} else {
		if len(all) > constant.PageSize {
			fights = all[:constant.PageSize]
		} else {
			fights = all
		}
	}

	total := len(fights)

	if filtered {
		if s.view.IsExtendedRequest(r) {
			s.view.RenderFragment(
				w,
				logsTable(fights, 0, total, startValue, endValue, true),
			)

			return
		}

		s.view.RenderPage(
			w,
			constant.LogsTitle,
			constant.LogsPath,
			html.H1(gomponents.Textf("Combat Logs (%d)", total)),
			dateFilter(startValue, endValue),
			html.Form(
				html.Class("generate-form"),
				html.Method("post"),
				logsTable(fights, 0, total, startValue, endValue, true),
				html.Button(
					html.Type("submit"),
					html.FormAction("/generate"),
					gomponents.Text("Generate Report"),
				),
				html.Button(
					html.Type("submit"),
					html.FormAction("/raids/create"),
					html.Class("secondary"),
					gomponents.Text("Create Raid"),
				),
			),
		)

		return
	}

	offset := parseIntParameter(r, "offset", 0)

	if offset > total {
		offset = total
	}

	end := offset + constant.PageSize

	if end > total {
		end = total
	}

	page := fights[offset:end]

	if s.view.IsExtendedRequest(r) {
		s.view.RenderFragment(
			w,
			logsTable(page, offset, total, startValue, endValue, false),
		)

		return
	}

	s.view.RenderPage(
		w,
		constant.LogsTitle,
		constant.LogsPath,
		html.H1(gomponents.Textf("Combat Logs (%d)", total)),
		dateFilter(startValue, endValue),
		html.Form(
			html.Class("generate-form"),
			html.Method("post"),
			logsTable(page, offset, total, startValue, endValue, false),
			html.Button(
				html.Type("submit"),
				html.FormAction("/generate"),
				gomponents.Text("Generate Report"),
			),
			html.Button(
				html.Type("submit"),
				html.FormAction("/raids/create"),
				html.Class("secondary"),
				gomponents.Text("Create Raid"),
			),
		),
	)
}
