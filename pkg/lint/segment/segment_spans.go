package segment

import "strings"

func segmentSpans(name string) []segmentSpan {
	var result []segmentSpan
	offset := 0

	for partIndex, part := range strings.Split(name, "_") {
		if partIndex > 0 {
			offset++
		}

		r := []rune(part)
		start := 0

		for i := 1; i < len(r); i++ {
			if !boundary(r, i) {
				continue
			}

			result = append(
				result,
				segmentSpan{
					start: offset + len(string(r[:start])),
					end:   offset + len(string(r[:i])),
					lower: strings.ToLower(string(r[start:i])),
				},
			)
			start = i
		}

		if start < len(r) {
			result = append(
				result,
				segmentSpan{
					start: offset + len(string(r[:start])),
					end:   offset + len(part),
					lower: strings.ToLower(string(r[start:])),
				},
			)
		}

		offset += len(part)
	}

	return result
}
