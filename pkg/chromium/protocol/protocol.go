package protocol

import (
	"context"
	"github.com/funtimecoding/soil/pkg/chromium"
	"time"
)

type Protocol struct {
	client     *chromium.Client
	identifier string
	context    context.Context
	timeout    time.Duration
}
