package layer_change

func New(
	path string,
	destination string,
	original []byte,
	current []byte,
) *Change {
	return &Change{
		Path:        path,
		Destination: destination,
		Original:    original,
		Current:     current,
	}
}
