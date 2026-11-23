package mock_token_store

import (
	"golang.org/x/oauth2"
	"sync"
)

type Store struct {
	Failure     error
	SaveFailure error
	lock        sync.Mutex
	token   *oauth2.Token
	saved   []*oauth2.Token
}
