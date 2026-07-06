package authenticator

import "github.com/funtimecoding/soil/pkg/web/session_store"

type Authenticator struct {
	store        *session_store.Store
	loginAddress string
}
