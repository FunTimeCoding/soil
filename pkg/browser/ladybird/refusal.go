package ladybird

import (
	"github.com/funtimecoding/soil/pkg/browser/constant"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/errors/validation"
)

func refusal(
	code string,
	text string,
) error {
	switch code {
	case constant.NoSuchNode, constant.NoSuchFrame:
		return not_found.Format("%s: %s", code, text)
	case constant.InvalidArgument:
		return validation.New("%s: %s", code, text)
	}

	return unexpected.Format("browser %s: %s", code, text)
}
