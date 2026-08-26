package testing

func SplitArchiveBytes(payload []byte) [][]byte {
	var segments [][]byte
	offset := 0

	for offset < len(payload) {
		nextBoundary := findNextGzipBoundary(payload, offset)

		if nextBoundary == -1 {
			segments = append(segments, payload[offset:])

			break
		}

		segments = append(segments, payload[offset:nextBoundary])
		offset = nextBoundary
	}

	return segments
}
