package firefox

import (
	"github.com/funtimecoding/soil/pkg/browser/constant"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"strings"
)

func refusal(message string) error {
	switch {
	case strings.HasPrefix(message, constant.FirefoxMatchedNothing),
		strings.HasPrefix(message, constant.FirefoxNotFoundInTime):

		return not_found.Format("%s", message)
	case message == constant.FirefoxSelectorRequired,
		message == constant.FirefoxExpressionRequired,
		message == constant.FirefoxLocatorRequired:

		return validation.New("%s", message)
	}

	return unexpected.Format("extension: %s", message)
}
