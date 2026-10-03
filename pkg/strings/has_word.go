package strings

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"regexp"
)

func HasWord(
	text string,
	word string,
) bool {
	result, e := regexp.MatchString(
		fmt.Sprintf(`\b%s\b`, regexp.QuoteMeta(word)),
		text,
	)
	errors.PanicOnError(e)

	return result
}
