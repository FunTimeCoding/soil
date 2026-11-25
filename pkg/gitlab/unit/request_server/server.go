package request_server

import (
	"github.com/funtimecoding/soil/pkg/gitlab"
	"net/url"
	"sync"
	"testing"
)

type Server struct {
	t       *testing.T
	total   int
	client  *gitlab.Client
	lock    sync.Mutex
	queries []url.Values
}
