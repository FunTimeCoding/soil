package service

import (
	"bytes"
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/connection"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/cluster"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/exec_result"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/request"
	"k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/remotecommand"
	"k8s.io/client-go/transport/spdy"
	"time"
)

func execInPodOn(
	x context.Context,
	c *cluster.Cluster,
	q request.Exec,
) (*exec_result.Result, error) {
	timeout := q.Timeout

	if timeout <= 0 {
		timeout = 60 * time.Second
	}

	b, cancel := context.WithTimeout(x, timeout)
	defer cancel()
	r := c.Clientset().CoreV1().RESTClient().Post().Resource("pods").Namespace(
		q.Namespace,
	).Name(
		q.Name,
	).SubResource(
		"exec",
	)
	r.VersionedParams(
		&v1.PodExecOptions{
			Command:   q.Command,
			Container: q.Container,
			Stdout:    true,
			Stderr:    true,
		},
		scheme.ParameterCodec,
	)
	transport, upgrader, f := spdy.RoundTripperFor(c.Configuration())

	if f != nil {
		return nil, fmt.Errorf("create SPDY transport: %w", f)
	}

	executor, g := remotecommand.NewSPDYExecutorForTransports(
		transport,
		upgrader,
		"POST",
		r.URL(),
	)

	if g != nil {
		return nil, fmt.Errorf("create executor: %w", g)
	}

	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	h := executor.StreamWithContext(
		b,
		remotecommand.StreamOptions{Stdout: stdout, Stderr: stderr},
	)

	if i := connection.Classify(h); i != nil {
		return nil, i
	}

	return exec_result.New(stdout.String(), stderr.String(), h), nil
}
