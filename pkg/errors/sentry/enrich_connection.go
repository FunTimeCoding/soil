package sentry

import (
	"github.com/funtimecoding/soil/pkg/errors/connection"
	"github.com/funtimecoding/soil/pkg/errors/constant"
	"github.com/getsentry/sentry-go"
)

func enrichConnection(
	e *sentry.Event,
	h *sentry.EventHint,
) *sentry.Event {
	if h == nil {
		return e
	}

	f := connection.Classify(original(h))

	if f == nil {
		return e
	}

	e.Fingerprint = f.Fingerprint()

	if e.Tags == nil {
		e.Tags = map[string]string{}
	}

	e.Tags[constant.Kind] = f.Kind
	e.Tags[constant.Host] = f.Host

	if f.Path != "" {
		e.Contexts[constant.Connection] = sentry.Context{constant.Path: f.Path}
	}

	for i := range e.Exception {
		e.Exception[i].Value = constant.QueryPattern.ReplaceAllString(
			e.Exception[i].Value,
			"$1",
		)
	}

	if last := len(e.Exception) - 1; last >= 0 {
		e.Exception[last].Type = f.Kind
		e.Exception[last].Value = f.Error()
	}

	return e
}
