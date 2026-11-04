package face

import "context"

type Reporter interface {
	Capture(
		x context.Context,
		e error,
	) string
	CaptureException(e error) string
	CaptureWithContext(
		e error,
		key string,
		context map[string]any,
	) string
	Recover(v any) string
}
