package classify

import "fmt"

func Message(
	e error,
	message string,
) string {
	if message == "" {
		return e.Error()
	}

	return fmt.Sprintf("%s: %s", message, e)
}
