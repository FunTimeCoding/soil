package authenticator

import (
	"github.com/funtimecoding/soil/pkg/network/constant"
	"github.com/funtimecoding/soil/pkg/web/session_store"
	"os"
)

func New() *Authenticator {
	var login string

	if a := os.Getenv(LoginAddressEnvironment); a != "" {
		login = a
	} else {
		login = constant.LocalhostAddressString
	}

	return &Authenticator{store: session_store.New(), loginAddress: login}
}
