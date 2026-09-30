package event

import "fmt"

func (e *EventError) Error() string {
	return fmt.Sprintf("event %s: %v", e.Type, e.Wrapped)
}
