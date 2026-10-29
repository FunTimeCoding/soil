package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_configured"
	"github.com/funtimecoding/soil/pkg/kubernetes/client/refusal"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/cluster"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/resource"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/response"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"sort"
	"strings"
)

func argocdApplicationsOn(
	x context.Context,
	c *cluster.Cluster,
) ([]response.ArgocdSummary, error) {
	apps, f := c.Dynamic().Resource(constant.ArgocdApplicationGVR).Namespace(
		constant.Argocd,
	).List(x, v1.ListOptions{})

	if f != nil {
		if strings.Contains(
			f.Error(),
			"could not find the requested resource",
		) {
			return nil, not_configured.Format(
				"ArgoCD not installed - Application CRD not found",
			)
		}

		return nil, refusal.Classify(f)
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
