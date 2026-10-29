package mock_client

import (
	"github.com/funtimecoding/soil/pkg/kubernetes/argocd/application"
	"sync"
)

type Client struct {
	mutex        sync.Mutex
	applications []*application.Application
	Failure      error
	Closed       int
}
