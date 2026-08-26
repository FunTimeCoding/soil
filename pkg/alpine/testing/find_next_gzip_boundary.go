package testing

import "github.com/funtimecoding/soil/pkg/alpine/constant"

func findNextGzipBoundary(
	payload []byte,
	startOffset int,
) int {
	searchStart := startOffset + constant.GzipMinimumSize

	for i := searchStart; i < len(payload)-constant.GzipHeaderSize; i++ {
		if payload[i] == 0x1f && payload[i+1] == 0x8b && payload[i+2] == 0x08 {
			return i
		}
	}

	return -1
}
