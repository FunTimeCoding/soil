package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/store/record"
	"time"
)

func state(r *record.Record) string {
	if r.Revoked != nil {
		return "revoked"
	}

	if r.End.Before(time.Now()) {
		return "expired"
	}

	return "live"
}
