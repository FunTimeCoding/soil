package client_configuration

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func New() *rest.Config {
	result, e := clientcmd.BuildConfigFromFlags("", path())
	errors.PanicOnError(e)

	return result
}
