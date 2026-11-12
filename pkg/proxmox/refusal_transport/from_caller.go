package refusal_transport

import (
	"context"
	"github.com/funtimecoding/soil/pkg/proxmox/refusal_transport/caller_key"
)

func fromCaller(x context.Context) bool {
	result, _ := x.Value(caller_key.Key{}).(bool)

	return result
}
