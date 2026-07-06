package service

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/format"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/resource"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/response"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sort"
	"strings"
)

var certificateGVR = schema.GroupVersionResource{
	Group:    "cert-manager.io",
	Version:  "v1",
	Resource: "certificates",
}

var certificateRequestGVR = schema.GroupVersionResource{
	Group:    "cert-manager.io",
	Version:  "v1",
	Resource: "certificaterequests",
}

func (s *Service) Certificates(
	x context.Context,
	clusterName string,
) ([]response.CertificateSummary, error) {
	c, e := s.ClusterByName(clusterName)

	if e != nil {
		return nil, e
	}

	certs, f := c.Dynamic().Resource(certificateGVR).Namespace(
		"",
	).List(x, v1.ListOptions{})

	if f != nil {
		if strings.Contains(
			f.Error(),
			"could not find the requested resource",
		) {
			return nil, fmt.Errorf("cert-manager not installed - Certificate CRD not found")
		}

		return nil, f
	}

	requests, g := c.Dynamic().Resource(certificateRequestGVR).Namespace(
		"",
	).List(x, v1.ListOptions{})
	var requestStatus map[string]string

	if g == nil {
		requestStatus = resource.BuildRequestStatus(requests)
	}

	result := []response.CertificateSummary{}

	for _, cert := range certs.Items {
		dnsNames := resource.ExtractDnsNames(cert.Object)
		notAfter := resource.ExtractNestedString(
			cert.Object,
			"status",
			"notAfter",
		)
		ready := resource.ExtractConditionStatus(cert.Object)
		summary := response.CertificateSummary{
			Name:      cert.GetName(),
			Namespace: cert.GetNamespace(),
			DnsNames:  join.CommaSpace(dnsNames),
			Ready:     ready,
			ExpiresAt: notAfter,
			Expires:   format.Expiry(notAfter),
		}
		key := fmt.Sprintf("%s/%s", cert.GetNamespace(), cert.GetName())

		if status, okay := requestStatus[key]; okay {
			summary.Renewal = status
		}

		result = append(result, summary)
	}

	sort.Slice(
		result,
		func(i, j int) bool {
			return result[i].ExpiresAt < result[j].ExpiresAt
		},
	)

	return result, nil
}
