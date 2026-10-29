package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/cluster"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/port_forward"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/request"
	"k8s.io/client-go/tools/portforward"
	"k8s.io/client-go/transport/spdy"
	"net/http"
	"net/url"
	"os"
)

func (s *Service) portForwardOn(
	c *cluster.Cluster,
	q request.PortForward,
	plural string,
) (*port_forward.Forward, error) {
	namespace := q.Namespace

	if namespace == "" {
		namespace = "default"
	}

	hostLocator, f := url.Parse(c.Configuration().Host)

	if f != nil {
		return nil, fmt.Errorf("parse host URL: %w", f)
	}

	hostLocator.Path = fmt.Sprintf(
		"/api/v1/namespaces/%s/%s/%s/portforward",
		namespace,
		plural,
		q.Name,
	)
	transport, upgrader, g := spdy.RoundTripperFor(c.Configuration())

	if g != nil {
		return nil, fmt.Errorf("create SPDY transport: %w", g)
	}

	result := port_forward.New(
		fmt.Sprintf("%s-%s-%d", q.ResourceType, q.Name, q.LocalPort),
		fmt.Sprintf("%s/%s", q.ResourceType, q.Name),
		namespace,
		q.LocalPort,
		q.TargetPort,
	)
	ready := make(chan struct{})
	forwarder, h := portforward.New(
		spdy.NewDialer(
			upgrader,
			&http.Client{Transport: transport},
			"POST",
			hostLocator,
		),
		[]string{fmt.Sprintf("%d:%d", q.LocalPort, q.TargetPort)},
		result.Stopping(),
		ready,
		os.Stdout,
		os.Stderr,
	)

	if h != nil {
		return nil, fmt.Errorf("create port forwarder: %w", h)
	}

	failed := make(chan error, 1)
	go func() {
		i := forwarder.ForwardPorts()

		select {
		case <-ready:
			s.forwardEnded(result, i)
		default:
			failed <- i
		}
	}()

	if j := port_forward.AwaitReady(ready, failed); j != nil {
		return nil, fmt.Errorf("start port forward: %w", j)
	}

	s.StorePortForward(result)

	return result, nil
}
