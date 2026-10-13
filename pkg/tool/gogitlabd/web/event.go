package web

import (
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
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

			if subs.Has(constant.SummaryEvent) {
				layout.PushEvent(
					w,
					webConstant.LayoutSummaryStrip,
					layout.SummaryStripContent(summary(entries)),
				)
			}

			if subs.Has(constant.RequestEvent) {
				layout.PushEvent(
					w,
					constant.RequestEvent,
					s.requestSection(s.worker.Requests()),
				)
			}

			if subs.Has(constant.BoardEvent) {
				layout.PushEvent(w, constant.BoardEvent, s.boardTable(entries))
			}

			if subs.Has(constant.PipelineEvent) {
				project := queryInteger(r, argumentConstant.Project)
				pipeline := queryInteger(r, argumentConstant.Pipeline)
				instance, g := s.queryInstance(r)

				if g == nil && project != 0 && pipeline != 0 {
					jobs, e := s.service.PipelineJobs(
						instance,
						project,
						pipeline,
					)
					errors.PanicOnError(e)
					layout.PushEvent(
						w,
						constant.PipelineEvent,
						stageStrip(
							instance,
							project,
							pipeline,
							jobs,
							selectJob(jobs, 0),
						),
					)
				}
			}

			f.Flush()
		},
	)
}
