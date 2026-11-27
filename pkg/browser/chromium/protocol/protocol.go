package protocol

import (
	"context"
	"github.com/funtimecoding/soil/pkg/browser/face"
	"time"
)

type Protocol struct {
	client     face.ChromiumClient
	identifier string
	context    context.Context
	timeout    time.Duration
}
