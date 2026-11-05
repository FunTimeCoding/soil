package mock_client

import (
	"github.com/funtimecoding/soil/pkg/chromium/face"
	"github.com/funtimecoding/soil/pkg/chromium/mock_client"
	"sync"
)

type Client struct {
	*mock_client.Client
	pages map[string]face.ProtocolClient
	mutex sync.Mutex
}
