package oversize_file

func New(
	collection string,
	path string,
) *File {
	return &File{Collection: collection, Path: path}
}
