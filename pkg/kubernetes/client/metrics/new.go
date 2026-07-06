package metrics

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"k8s.io/client-go/rest"
	"k8s.io/metrics/pkg/client/clientset/versioned"
)

func New(c *rest.Config) *versioned.Clientset {
	result, e := versioned.NewForConfig(c)
	errors.PanicOnError(e)

	return result
}
