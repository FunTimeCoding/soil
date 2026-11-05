package protocol

import (
	"context"
	"github.com/funtimecoding/soil/pkg/chromium/face"
	"time"
)

type Protocol struct {
	client     face.ChromiumClient
	identifier string
	context    context.Context
	timeout    time.Duration
}
