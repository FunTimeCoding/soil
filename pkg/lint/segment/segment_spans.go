package segment

import (
	"strings"
	"unicode"
)

func segmentSpans(name string) []segmentSpan {
	var result []segmentSpan
	offset := 0

	for partIndex, part := range strings.Split(name, "_") {
		if partIndex > 0 {
			offset++
		}

		segmentStart := offset

		for i, r := range part {
			if i > 0 && unicode.IsUpper(r) {
				if segmentStart < offset+i {
					result = append(
						result,
						segmentSpan{
							start: segmentStart,
							end:   offset + i,
							lower: strings.ToLower(
								name[segmentStart : offset+i],
							),
						},
					)
				}

				segmentStart = offset + i
			}
		}

		if segmentStart < offset+len(part) {
			result = append(
				result,
				segmentSpan{
					start: segmentStart,
					end:   offset + len(part),
					lower: strings.ToLower(
						name[segmentStart : offset+len(part)],
					),
				},
			)
		}

		offset += len(part)
	}

	return result
}
