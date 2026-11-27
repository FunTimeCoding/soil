package mock_client

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/browser/ladybird/types/call"
	"github.com/funtimecoding/soil/pkg/browser/ladybird/types/message"
	"github.com/funtimecoding/soil/pkg/browser/ladybird/types/response"
	"sync"
	"sync/atomic"
)

type Client struct {
	Failure     error
	mutex       sync.Mutex
	contexts    []*response.Context
	results     map[string]json.RawMessage
	refusals    map[string]error
	navigations map[string][]*message.Reply
	event       func(*message.Reply)
	calls       []*call.Call
	typed       []string
	image       []byte
	closed      atomic.Bool
}
