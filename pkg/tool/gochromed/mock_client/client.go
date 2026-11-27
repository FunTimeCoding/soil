package mock_client

import (
	"github.com/funtimecoding/soil/pkg/browser/chromium/mock_client"
	"github.com/funtimecoding/soil/pkg/browser/face"
	"sync"
)

type Client struct {
	*mock_client.Client
	pages map[string]face.ChromiumProtocolClient
	mutex sync.Mutex
}
