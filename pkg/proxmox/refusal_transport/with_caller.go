package refusal_transport

import (
	"context"
	"github.com/funtimecoding/soil/pkg/proxmox/refusal_transport/caller_key"
)

func WithCaller(x context.Context) context.Context {
	return context.WithValue(x, caller_key.Key{}, true)
}
