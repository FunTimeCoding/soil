package preposition_name

import "github.com/funtimecoding/soil/pkg/lint/constant"

func placedWord(
	parts []string,
	start int,
) string {
	last := len(parts) - 1

	for i := start; i <= last; i++ {
		s := parts[i]

		if constant.PrepositionAnywhere[s] {
			return s
		}

		if constant.PrepositionNotFirst[s] &&
			(i > 0 || last == 0) &&
			parts[0] != constant.NewSegment {
			return s
		}

		if constant.PrepositionLast[s] && i == last {
			return s
		}
	}

	return ""
}
