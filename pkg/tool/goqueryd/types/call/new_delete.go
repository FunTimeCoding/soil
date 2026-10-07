package call

func NewDelete(
	collection string,
	path string,
) *Delete {
	return &Delete{Collection: collection, Path: path}
}
