package authenticator

import (
	networkConstant "github.com/funtimecoding/soil/pkg/network/constant"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/session_store"
	"os"
)

func New() *Authenticator {
	var login string

	if a := os.Getenv(constant.LoginAddressEnvironment); a != "" {
		login = a
	} else {
		login = networkConstant.LocalhostAddressString
	}

	return &Authenticator{store: session_store.New(), loginAddress: login}
}
