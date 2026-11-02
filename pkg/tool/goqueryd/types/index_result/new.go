package index_result

func New(
	collection string,
	indexed int,
	updated int,
	unchanged int,
	removed int,
) *IndexResult {
	return &IndexResult{
		Collection: collection,
		Indexed:    indexed,
		Updated:    updated,
		Unchanged:  unchanged,
		Removed:    removed,
	}
}
