package example

import (
	"context"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/location"
	"github.com/funtimecoding/soil/pkg/web/request_context"
	"log/slog"
	"net/http"
	"os"
)

func Log() {
	x := context.Background()
	l := slog.New(
		slog.NewJSONHandler(
			os.Stdout,
			&slog.HandlerOptions{Level: slog.LevelInfo},
		),
	)
	m := http.NewServeMux()
	m.HandleFunc(
		location.Root,
		func(
			w http.ResponseWriter,
			r *http.Request,
		) {
			c := request_context.New(w, r)
			c.LogStart(x, l)
			c.WriteSuccess()
		},
	)
	web.Serve(context.Background(), m, 8080, false)
}
