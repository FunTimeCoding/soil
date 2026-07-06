package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/certificate_detail"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/resource"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (s *Service) Certificate(
	x context.Context,
	clusterName string,
	name string,
	namespace string,
	unfiltered bool,
) (*certificate_detail.Detail, error) {
	c, e := s.ClusterByName(clusterName)

	if e != nil {
		return nil, e
	}

	cert, f := c.Dynamic().Resource(certificateGVR).Namespace(
		namespace,
	).Get(x, name, v1.GetOptions{})

	if f != nil {
		return nil, f
	}

	requests, g := c.Dynamic().Resource(certificateRequestGVR).Namespace(
		namespace,
	).List(x, v1.ListOptions{})
	var latestRequest map[string]any

	if g == nil {
		for _, q := range requests.Items {
			if resource.ExtractOwnerCertificate(q.Object) == name {
				latestRequest = q.Object
			}
		}
	}

	filtered := resource.FilterObject(cert.Object, unfiltered)
	var requestObject map[string]any
	allFiltered := filtered.Filtered

	if latestRequest != nil {
		requestFiltered := resource.FilterObject(latestRequest, unfiltered)
		requestObject = requestFiltered.Object

		if len(requestFiltered.Filtered) > 0 {
			allFiltered = append(allFiltered, requestFiltered.Filtered...)
		}
	}

	return certificate_detail.New(
		filtered.Object,
		requestObject,
		allFiltered,
	), nil
}
