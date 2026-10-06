package chunk

import "strings"

func prose(
	text string,
	offset int,
) []Chunk {
	if strings.TrimSpace(text) == "" {
		return nil
	}

	return segment(text, offset)
}
