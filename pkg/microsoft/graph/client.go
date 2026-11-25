package graph

import (
	"github.com/funtimecoding/soil/pkg/microsoft/face"
	"github.com/funtimecoding/soil/pkg/web/requester"
	"github.com/funtimecoding/soil/pkg/web/stall_client/pool"
	"golang.org/x/oauth2"
	"sync"
)

type Client struct {
	requester     *requester.Requester
	pool          *pool.Pool
	configuration *oauth2.Config
	store         face.TokenStore
	lock          sync.Mutex
	token         *oauth2.Token
	stored        string
}
