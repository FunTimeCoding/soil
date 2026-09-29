package model_context

func paginate[T any](
	items []T,
	limit int,
	offset int,
) []T {
	if offset > 0 {
		if offset >= len(items) {
			return nil
		}

		items = items[offset:]
	}

	if limit > 0 && limit < len(items) {
		items = items[:limit]
	}

	return items
}
