package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/argocd_detail"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/resource"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (s *Service) ArgocdApplication(
	x context.Context,
	clusterName string,
	name string,
	unfiltered bool,
) (*argocd_detail.Detail, error) {
	c, e := s.ClusterByName(clusterName)

	if e != nil {
		return nil, e
	}

	app, f := c.Dynamic().Resource(argocdApplicationGVR).Namespace(
		constant.Argocd,
	).Get(x, name, v1.GetOptions{})

	if f != nil {
		return nil, f
	}

	filtered := resource.FilterObject(app.Object, unfiltered)
	syncStatus := resource.ExtractNestedString(
		app.Object,
		"status",
		"sync",
		"status",
	)

	if syncStatus == "Synced" && !unfiltered {
		resource.StripArgocdNoise(filtered.Object, &filtered.Filtered)
	}

	return argocd_detail.New(filtered.Object, filtered.Filtered), nil
}
