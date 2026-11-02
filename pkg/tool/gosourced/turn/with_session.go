package turn

import "context"

func WithSession(
	x context.Context,
	session string,
) context.Context {
	return context.WithValue(x, sessionKey{}, session)
}
