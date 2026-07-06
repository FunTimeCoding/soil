package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	kubernetes "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Client) DeleteJob(
	namespace string,
	name string,
) {
	e := c.client.BatchV1().Jobs(namespace).Delete(
		c.context,
		name,
		v1.DeleteOptions{
			PropagationPolicy: new(v1.DeletePropagationForeground),
		},
	)

	if e != nil {
		if kubernetes.IsNotFound(e) {
			return
		}
	}

	errors.PanicOnError(e)
}
