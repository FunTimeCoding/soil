package protocol

import (
	"context"
	"github.com/funtimecoding/soil/pkg/chromium"
)

type Protocol struct {
	client  *chromium.Client
	context context.Context
}
