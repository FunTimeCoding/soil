package service

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/resource"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/response"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sort"
	"strings"
)

var argocdApplicationGVR = schema.GroupVersionResource{
	Group:    "argoproj.io",
	Version:  "v1alpha1",
	Resource: "applications",
}

func (s *Service) ArgocdApplications(
	x context.Context,
	clusterName string,
) ([]response.ArgocdSummary, error) {
	c, e := s.ClusterByName(clusterName)

	if e != nil {
		return nil, e
	}

	apps, f := c.Dynamic().Resource(argocdApplicationGVR).Namespace(
		constant.Argocd,
	).List(x, v1.ListOptions{})

	if f != nil {
		if strings.Contains(
			f.Error(),
			"could not find the requested resource",
		) {
			return nil, fmt.Errorf("ArgoCD not installed - Application CRD not found")
		}

		return nil, f
	}

	result := []response.ArgocdSummary{}

	for _, app := range apps.Items {
		result = append(
			result,
			response.ArgocdSummary{
				Name: app.GetName(),
				Namespace: resource.ExtractNestedString(
					app.Object,
					"spec",
					"destination",
					"namespace",
				),
				Sync: resource.ExtractNestedString(
					app.Object,
					"status",
					"sync",
					"status",
				),
				Health: resource.ExtractNestedString(
					app.Object,
					"status",
					"health",
					"status",
				),
				Operation: resource.ExtractNestedString(
					app.Object,
					"status",
					"operationState",
					"phase",
				),
			},
		)
	}

	sort.Slice(
		result,
		func(i, j int) bool {
			return result[i].Name < result[j].Name
		},
	)

	return result, nil
}
