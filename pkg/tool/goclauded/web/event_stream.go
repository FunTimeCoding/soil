package web

import (
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"net/http"
	"strconv"
)

func (s *Server) eventStream() http.HandlerFunc {
	return layout.HandleServerSideEventWithRequest(
		s.notifier,
		func(
			w http.ResponseWriter,
			f http.Flusher,
			r *http.Request,
		) {
			name := r.URL.Query().Get(constant.Subscriber)

			if name == "" {
				return
			}

			events, e := s.service.NextEvents(
				name,
				kindList(r),
				streamOverride(r),
				constant.StreamBatch,
			)

			if e != nil {
				return
			}

			for _, v := range events {
				layout.PushPayload(
					w,
					constant.EventStreamName,
					strconv.FormatUint(uint64(v.Identifier), 10),
					string(notation.Marshal(newStreamEvent(v))),
				)
			}

			f.Flush()
		},
	)
}
