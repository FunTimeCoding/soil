package counting_mattermost

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost/mock_client"
	"sync/atomic"
)

type Client struct {
	*mock_client.Client
	me          atomic.Int32
	teamChannel atomic.Int32
}
