package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/cluster"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/port_forward"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/request"
)

func (s *Service) PortForward(
	x context.Context,
	clusterName string,
	q request.PortForward,
) (*port_forward.Forward, error) {
	if q.ResourceType == "" {
		return nil, validation.New("resourceType is required")
	}

	if q.Name == "" {
		return nil, validation.New("name is required")
	}

	if q.LocalPort <= 0 {
		return nil, validation.New("localPort is required")
	}

	if q.TargetPort <= 0 {
		return nil, validation.New("targetPort is required")
	}

	plural, okay := constant.PortForwardResource[q.ResourceType]

	if !okay {
		return nil, validation.New(
			"unsupported resource type for port forward: %s (use pod or service)",
			q.ResourceType,
		)
	}

	return guarded(
		s,
		x,
		clusterName,
		func(c *cluster.Cluster) (*port_forward.Forward, error) {
			return s.portForwardOn(c, q, plural)
		},
	)
}
