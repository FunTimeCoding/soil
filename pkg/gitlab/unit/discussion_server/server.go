package discussion_server

import (
	"github.com/funtimecoding/soil/pkg/gitlab"
	"sync"
	"testing"
)

type Server struct {
	t        *testing.T
	client   *gitlab.Client
	lock     sync.Mutex
	requests []string
}
