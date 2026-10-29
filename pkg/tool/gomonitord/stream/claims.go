package stream

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/convert"
	monitor "github.com/funtimecoding/soil/pkg/tool/gomonitord/face"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"net/http"
)

func Claims(
	v monitor.ClaimStream,
	r face.Reporter,
) http.HandlerFunc {
	return layout.HandleServerSideEvent(
		v,
		func(
			w http.ResponseWriter,
			f http.Flusher,
		) {
			claims, e := v.Claims()

			if e != nil {
				r.CaptureException(e)

				return
			}

			if _, g := fmt.Fprintf(
				w,
				constant.EventFormat,
				notation.Marshal(convert.Claims(claims)),
			); g != nil {
				return
			}

			f.Flush()
		},
	)
}
