package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/generated/server"
	"net/http"
	"slices"
	"time"
)

func ExtendedWrite(
	f server.StrictHandlerFunc,
	operation string,
) server.StrictHandlerFunc {
	if !slices.Contains(constant.BackupOperations, operation) {
		return f
	}

	return func(
		x context.Context,
		w http.ResponseWriter,
		r *http.Request,
		request any,
	) (any, error) {
		if e := http.NewResponseController(w).SetWriteDeadline(
			time.Now().Add(constant.BackupWriteDeadline),
		); e != nil {
			return nil, e
		}

		return f(x, w, r, request)
	}
}
