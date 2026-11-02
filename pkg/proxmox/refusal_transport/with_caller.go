package refusal_transport

import "context"

func WithCaller(x context.Context) context.Context {
	return context.WithValue(x, callerKey{}, true)
}
