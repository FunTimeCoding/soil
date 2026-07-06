package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/strings/join/key_value"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
)

func (c *Client) DeleteJobWatch(
	namespace string,
	name string,
) {
	if c.Job(namespace, name) == nil {
		return
	}

	c.DeleteJob(namespace, name)
	w, e := c.client.BatchV1().Jobs(namespace).Watch(
		c.context,
		v1.ListOptions{
			FieldSelector:  key_value.Equals("metadata.name", name),
			TimeoutSeconds: new(int64(120)),
		},
	)
	errors.PanicOnError(e)
	defer w.Stop()

	for v := range w.ResultChan() {
		if v.Type == watch.Deleted {
			return
		}
	}
}
