package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/kubernetes/client/executor"
	"github.com/funtimecoding/soil/pkg/kubernetes/constant"
	"io"
	"k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/remotecommand"
)

func (c *Client) Execute(
	stdout io.Writer,
	stderr io.Writer,
	namespace string,
	pod string,
	container string,
	command ...string,
) {
	r := c.client.CoreV1().RESTClient().Post().Resource(
		constant.PodsResource,
	).Namespace(namespace).Name(pod).SubResource(
		constant.ExecuteSubResource,
	)
	r.VersionedParams(
		&v1.PodExecOptions{
			Command:   command,
			Container: container,
			Stdout:    true,
			Stderr:    true,
		},
		scheme.ParameterCodec,
	)
	errors.PanicOnError(
		executor.New(c.configuration, r).StreamWithContext(
			c.context,
			remotecommand.StreamOptions{Stdout: stdout, Stderr: stderr},
		),
	)
}
