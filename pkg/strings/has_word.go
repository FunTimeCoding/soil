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
	// matches word with word boundaries
	// prevents partial matches like "test" matching "testing"
	result, e := regexp.MatchString(
		fmt.Sprintf(`\b%s\b`, regexp.QuoteMeta(word)),
		text,
	)
	errors.PanicOnError(e)

	return result
}
