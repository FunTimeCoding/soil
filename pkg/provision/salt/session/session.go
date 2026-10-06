package session

import (
	"github.com/funtimecoding/soil/pkg/web/requester"
	"sync"
)

type Session struct {
	mutex   sync.Mutex
	login   *requester.Requester
	request loginRequest
	token   string
}
