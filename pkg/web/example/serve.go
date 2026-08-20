package example

import (
	"github.com/funtimecoding/soil/pkg/text/multi_line"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/request_context"
	"net/http"
)

func Serve() {
	m := http.NewServeMux()
	m.HandleFunc(
		constant.RootPath,
		func(
			w http.ResponseWriter,
			e *http.Request,
		) {
			c := request_context.New(w, e)
			c.SetLastLocation()
			l := multi_line.New()
			l.Add("Hello friend.")
			c.WriteOkay(l.Render())
		},
	)
	web.Listen(m)
}
