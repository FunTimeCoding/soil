package client_configuration

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
)

func Interface() *api.Config {
	result, f := clientcmd.LoadFromFile(path())
	errors.PanicOnError(f)

	return result
}
