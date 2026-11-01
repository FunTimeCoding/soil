package mock_client

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/types/notify_call"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/types/pulse_call"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/types/target"
	"sync"
)

type Client struct {
	mutex          sync.Mutex
	sessions       []*target.Target
	notes          []*notify_call.Call
	Pulses         []*pulse_call.Call
	Failure        error
	ReleaseFailure error
	Closed         int
}
