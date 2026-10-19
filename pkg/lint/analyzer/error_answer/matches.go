package error_answer

import "github.com/funtimecoding/soil/pkg/errors/types/answer"

func matches(
	want *answer.Answer,
	observed *answer.Answer,
) bool {
	if observed.Status == 0 {
		return want.Captured == observed.Captured
	}

	return *want == *observed
}
