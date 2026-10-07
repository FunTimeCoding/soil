package session

import (
	"github.com/funtimecoding/soil/pkg/provision/types/salt_login_request"
	"github.com/funtimecoding/soil/pkg/web/requester"
	"sync"
)

type Session struct {
	mutex   sync.Mutex
	login   *requester.Requester
	request salt_login_request.Request
	token   string
}
