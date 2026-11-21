package minimum

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
)

func Check(minimum float64) error {
	if minimum >= 0 && minimum <= constant.HighestScore {
		return nil
	}

	return validation.New("score outside 0 to 10: %v", minimum)
}
