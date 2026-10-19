package error_answer

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/types/answer"
)

func describe(
	a *answer.Answer,
	tool bool,
) string {
	if tool && a.Captured {
		return "captureFail"
	}

	if tool {
		return "response.Fail"
	}

	if a.Captured {
		return fmt.Sprintf("%d ErrorResponse", a.Status)
	}

	return fmt.Sprintf("%d Error", a.Status)
}
