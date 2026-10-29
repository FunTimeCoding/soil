package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/kubernetes/client/refusal"
	kubernetesConstant "github.com/funtimecoding/soil/pkg/kubernetes/constant"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/argocd_detail"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/cluster"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/resource"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

func argocdApplicationOn(
	x context.Context,
	c *cluster.Cluster,
	name string,
	unfiltered bool,
) (*argocd_detail.Detail, error) {
	app, f := c.Dynamic().Resource(constant.ArgocdApplicationGVR).Namespace(
		constant.Argocd,
	).Get(x, name, v1.GetOptions{})

	if f != nil {
		return nil, refusal.Classify(f)
	}

	filtered := resource.FilterObject(app.Object, unfiltered)
	syncStatus := resource.ExtractNestedString(
		app.Object,
		"status",
		"sync",
		"status",
	)

	if syncStatus == kubernetesConstant.ArgocdSyncSynced && !unfiltered {
		resource.StripArgocdNoise(filtered.Object, &filtered.Filtered)
	}

	return argocd_detail.New(filtered.Object, filtered.Filtered), nil
}
