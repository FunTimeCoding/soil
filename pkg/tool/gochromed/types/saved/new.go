package saved

func New(
	path string,
	title string,
	size int,
) *Saved {
	return &Saved{Path: path, Title: title, Size: size}
}
