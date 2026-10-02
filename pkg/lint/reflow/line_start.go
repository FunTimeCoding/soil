package reflow

import "bytes"

func lineStart(
	source []byte,
	at int,
) int {
	if before := bytes.LastIndexByte(source[:at], '\n'); before >= 0 {
		return before + 1
	}

	return 0
}
