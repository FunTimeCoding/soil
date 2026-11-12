package turn

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/turn/session_key"
)

func WithSession(
	x context.Context,
	session string,
) context.Context {
	return context.WithValue(x, session_key.Key{}, session)
}
