package chunk

func splitAtBoundaries(
	content string,
	boundaries []int,
) []Chunk {
	var result []Chunk
	start := 0

	for _, boundary := range boundaries {
		if boundary <= start {
			continue
		}

		result = append(result, segment(content[start:boundary], start)...)
		start = boundary
	}

	if start < len(content) {
		result = append(result, segment(content[start:], start)...)
	}

	return result
}
