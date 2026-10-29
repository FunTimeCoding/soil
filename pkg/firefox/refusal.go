package firefox

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/firefox/constant"
	"strings"
)

func refusal(message string) error {
	switch {
	case strings.HasPrefix(message, constant.MatchedNothing),
		strings.HasPrefix(message, constant.NotFoundInTime):

		return not_found.Format("%s", message)
	case message == constant.SelectorRequired,
		message == constant.ExpressionRequired,
		message == constant.LocatorRequired:

		return validation.New("%s", message)
	}

	return unexpected.Format("extension: %s", message)
}
