package client

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/kubernetes/client/refusal"
	"github.com/funtimecoding/soil/pkg/strings/join/key_value"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
)

func (c *Client) DeleteJobWatch(
	namespace string,
	name string,
) error {
	_, e := c.Job(namespace, name)

	if not_found.Is(e) {
		return nil
	}

	if e != nil {
		return e
	}

	if f := c.DeleteJob(namespace, name); f != nil {
		return f
	}

	w, g := c.client.BatchV1().Jobs(namespace).Watch(
		c.context,
		v1.ListOptions{
			FieldSelector:  key_value.Equals("metadata.name", name),
			TimeoutSeconds: new(int64(120)),
		},
	)

	if g != nil {
		return refusal.Classify(g)
	}

	defer w.Stop()

	for v := range w.ResultChan() {
		if v.Type == watch.Deleted {
			return nil
		}
	}

	return nil
}
