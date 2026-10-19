package classify

import (
	"github.com/funtimecoding/soil/pkg/errors/constant"
	"github.com/funtimecoding/soil/pkg/errors/types/answer"
)

func Answer(e error) *answer.Answer {
	return constant.Answers[Class(e)]
}
