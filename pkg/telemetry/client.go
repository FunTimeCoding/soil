package telemetry

import (
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/client"
	"sync"
)

type Client struct {
	client *client.Client
	group  sync.WaitGroup
}
