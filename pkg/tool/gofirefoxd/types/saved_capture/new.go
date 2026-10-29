package saved_capture

func New(
	path string,
	title string,
	size int,
) *Capture {
	return &Capture{Path: path, Title: title, Size: size}
}
