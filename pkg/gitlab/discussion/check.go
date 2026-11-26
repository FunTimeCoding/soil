package discussion

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/gitlab/constant"
	"regexp"
)

func Check(identifier string) error {
	if !regexp.MustCompile(constant.DiscussionPattern).MatchString(identifier) {
		return validation.New(
			"discussion must be 40 lowercase hex characters: %q",
			identifier,
		)
	}

	return nil
}
