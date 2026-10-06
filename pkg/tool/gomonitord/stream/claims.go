package stream

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/convert"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/store"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"net/http"
)

func Claims(
	s *store.Store,
	n face.EventNotifier,
	r face.Reporter,
) http.HandlerFunc {
	return layout.HandleServerSideEvent(
		n,
		func(
			w http.ResponseWriter,
			f http.Flusher,
		) {
			v, e := s.Claims()

			if e != nil {
				r.CaptureException(e)

				return
			}

			if _, g := fmt.Fprintf(
				w,
				constant.EventFormat,
				notation.Marshal(convert.Claims(v)),
			); g != nil {
				return
			}

			f.Flush()
		},
	)
}
