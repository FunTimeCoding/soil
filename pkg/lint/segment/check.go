package segment

import "github.com/funtimecoding/soil/pkg/lint/constant"

func Check(
	name string,
	isVariable bool,
	isField bool,
) *Result {
	parts := Segments(name)
	multiSegment := len(parts) > 1
	allowSingleLetter := isVariable && !isField && !multiSegment

	for _, s := range parts {
		if constant.NoSuggestion[s] {
			return &Result{Segment: s, Banned: true}
		}

		entry, hasSuggestion := constant.Suggestions[s]

		if !hasSuggestion {
			continue
		}

		var applicable []string

		if allowSingleLetter {
			applicable = append(applicable, entry.Letters...)
		}

		applicable = append(applicable, entry.Words...)

		if len(applicable) == 0 {
			return &Result{Segment: s, Banned: true}
		}

		for _, a := range applicable {
			if ContainsSegment(name, a) {
				return nil
			}
		}

		return &Result{Segment: s, Applicable: applicable}
	}

	return nil
}
