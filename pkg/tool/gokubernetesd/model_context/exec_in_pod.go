package model_context

import (
	"bytes"
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
	"k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/remotecommand"
	"k8s.io/client-go/transport/spdy"
	"time"
)

func (s *Server) ExecInPod(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.ExecInPod,
) (*mcp.CallToolResult, error) {
	if a.Name == "" {
		return response.Fail("name is required")
	}

	if a.Namespace == "" {
		return response.Fail("namespace is required")
	}

	if len(a.Command) == 0 {
		return response.Fail("command is required")
	}

	c, e := s.activeCluster(x)

	if e != nil {
		return response.Fail(e.Error())
	}

	timeout := time.Duration(a.Timeout) * time.Millisecond

	if timeout <= 0 {
		timeout = 60 * time.Second
	}

	b, cancel := context.WithTimeout(x, timeout)
	defer cancel()
	request := c.Clientset().CoreV1().RESTClient().Post().Resource(
		"pods",
	).Namespace(a.Namespace).Name(a.Name).SubResource("exec")
	request.VersionedParams(
		&v1.PodExecOptions{
			Command:   a.Command,
			Container: a.Container,
			Stdout:    true,
			Stderr:    true,
		},
		scheme.ParameterCodec,
	)
	transport, upgrader, f := spdy.RoundTripperFor(c.Configuration())

	if f != nil {
		return s.captureFail(f, "create SPDY transport")
	}

	executor, g := remotecommand.NewSPDYExecutorForTransports(
		transport,
		upgrader,
		"POST",
		request.URL(),
	)

	if g != nil {
		return s.captureFail(g, "create executor")
	}

	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	h := executor.StreamWithContext(
		b,
		remotecommand.StreamOptions{Stdout: stdout, Stderr: stderr},
	)

	if h != nil {
		if stderr.Len() > 0 {
			return response.Fail(
				"command failed: %s\nstderr: %s",
				h,
				stderr.String(),
			)
		}

		return s.captureFail(h, "exec in pod")
	}

	if stderr.Len() > 0 {
		return response.Success(
			fmt.Sprintf(
				"%s\n--- stderr ---\n%s",
				stdout.String(),
				stderr.String(),
			),
		)
	}

	return response.Success(stdout.String())
}
