package refusal_transport

import "context"

func fromCaller(x context.Context) bool {
	result, _ := x.Value(callerKey{}).(bool)

	return result
}
