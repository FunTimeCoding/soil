package turn

import "context"

func SessionOf(x context.Context) string {
	result, _ := x.Value(sessionKey{}).(string)

	return result
}
