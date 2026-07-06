package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func New(c *rest.Config) *kubernetes.Clientset {
	result, e := kubernetes.NewForConfig(c)
	errors.PanicOnError(e)

	return result
}
