package reporter

import "github.com/getsentry/sentry-go"

func (r *Reporter) CaptureWithContext(
	e error,
	key string,
	context map[string]any,
) string {
	if r.echo {
		r.print(e, key, context)
	}

	if r.hub == nil {
		return ""
	}

	var identifier string
	r.hub.WithScope(
		func(s *sentry.Scope) {
			s.SetContext(key, context)
			result := r.hub.CaptureException(e)

			if result != nil {
				identifier = string(*result)
			}
		},
	)

	return identifier
}
