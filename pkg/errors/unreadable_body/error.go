package unreadable_body

import "fmt"

func (e *UnreadableBodyError) Error() string {
	return fmt.Sprintf("%s (body unreadable: %v)", e.Message, e.Wrapped)
}
