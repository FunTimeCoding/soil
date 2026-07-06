package model_context

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/model_context/argument"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/port_forward_state"
	"github.com/mark3labs/mcp-go/mcp"
	"k8s.io/client-go/tools/portforward"
	"k8s.io/client-go/transport/spdy"
	"net/http"
	"net/url"
	"os"
)

func (s *Server) PortForward(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.PortForward,
) (*mcp.CallToolResult, error) {
	if a.ResourceType == "" {
		return response.Fail("resourceType is required")
	}

	if a.Name == "" {
		return response.Fail("name is required")
	}

	if a.LocalPort <= 0 {
		return response.Fail("localPort is required")
	}

	if a.TargetPort <= 0 {
		return response.Fail("targetPort is required")
	}

	c, e := s.activeCluster(x)

	if e != nil {
		return response.Fail(e.Error())
	}

	plural, okay := portForwardResource[a.ResourceType]

	if !okay {
		return response.Fail(
			"unsupported resource type for port forward: %s (use pod or service)",
			a.ResourceType,
		)
	}

	namespace := a.Namespace

	if namespace == "" {
		namespace = "default"
	}

	path := fmt.Sprintf(
		"/api/v1/namespaces/%s/%s/%s/portforward",
		namespace,
		plural,
		a.Name,
	)
	hostURL, f := url.Parse(c.Configuration().Host)

	if f != nil {
		return s.captureFail(f, "parse host URL")
	}

	hostURL.Path = path
	transport, upgrader, g := spdy.RoundTripperFor(c.Configuration())

	if g != nil {
		return s.captureFail(g, "create SPDY transport")
	}

	dialer := spdy.NewDialer(
		upgrader,
		&http.Client{Transport: transport},
		"POST",
		hostURL,
	)
	stop := make(chan struct{})
	ready := make(chan struct{})
	ports := []string{fmt.Sprintf("%d:%d", a.LocalPort, a.TargetPort)}
	forwarder, h := portforward.New(
		dialer,
		ports,
		stop,
		ready,
		os.Stdout,
		os.Stderr,
	)

	if h != nil {
		return s.captureFail(h, "create port forwarder")
	}

	go func() {
		if i := forwarder.ForwardPorts(); i != nil {
			s.reporter.CaptureException(i)
		}
	}()
	<-ready
	identifier := fmt.Sprintf(
		"%s-%s-%d",
		a.ResourceType,
		a.Name,
		a.LocalPort,
	)
	s.service.StorePortForward(
		identifier,
		port_forward_state.New(
			identifier,
			fmt.Sprintf("%s/%s", a.ResourceType, a.Name),
			namespace,
			a.LocalPort,
			a.TargetPort,
			stop,
		),
	)

	return response.SuccessAny(
		map[string]any{
			"id":         identifier,
			"localPort":  a.LocalPort,
			"targetPort": a.TargetPort,
			"resource":   fmt.Sprintf("%s/%s", a.ResourceType, a.Name),
			"namespace":  namespace,
		},
	)
}
