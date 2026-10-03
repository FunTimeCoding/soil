package protocol

import (
	"context"
	"github.com/funtimecoding/soil/pkg/chromium"
	"time"
)

type Protocol struct {
	client  *chromium.Client
	context context.Context
	timeout time.Duration
}
