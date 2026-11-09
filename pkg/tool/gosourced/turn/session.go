package turn

import "context"

func Session(x context.Context) string {
	result, _ := x.Value(sessionKey{}).(string)

	return result
}
