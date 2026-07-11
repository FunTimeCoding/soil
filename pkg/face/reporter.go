package face

type Reporter interface {
	CaptureException(e error) string
	CaptureWithContext(
		e error,
		key string,
		context map[string]any,
	) string
	Recover(v any) string
}
