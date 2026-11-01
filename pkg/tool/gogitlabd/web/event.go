package web

import (
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/web/board_page"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/web/pipeline_page"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"github.com/funtimecoding/soil/pkg/web/subscription"
	"net/http"
)

func (s *Server) event() http.HandlerFunc {
	return layout.HandleServerSideEventWithRequest(
		s.worker.Notifier(),
		func(
			w http.ResponseWriter,
			f http.Flusher,
			r *http.Request,
		) {
			subs := subscription.Parse(r)
			entries := s.worker.Entries()
			multiple := s.multiple()

			if subs.Has(constant.SummaryEvent) {
				layout.PushEvent(
					w,
					webConstant.LayoutSummaryStrip,
					layout.SummaryStripContent(board_page.Summary(entries)),
				)
			}

			if subs.Has(constant.RequestEvent) {
				layout.PushEvent(
					w,
					constant.RequestEvent,
					board_page.Requests(s.worker.Requests(), multiple),
				)
			}

			if subs.Has(constant.BoardEvent) {
				layout.PushEvent(
					w,
					constant.BoardEvent,
					board_page.Table(entries, multiple),
				)
			}

			if subs.Has(constant.PipelineEvent) {
				project := queryInteger(r, argumentConstant.Project)
				pipeline := queryInteger(r, argumentConstant.Pipeline)
				instance, g := s.queryInstance(r)

				if g == nil && project != 0 && pipeline != 0 {
					v, e := s.service.JobView(instance, project, pipeline, 0)
					errors.PanicOnError(e)
					layout.PushEvent(
						w,
						constant.PipelineEvent,
						pipeline_page.New(v).StageStrip(),
					)
				}
			}

			f.Flush()
		},
	)
}
