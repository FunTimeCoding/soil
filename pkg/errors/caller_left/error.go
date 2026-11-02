package caller_left

import "fmt"

func (e *CallerLeft) Error() string {
	return fmt.Sprintf("caller left before the answer: %s", e.cause)
}
