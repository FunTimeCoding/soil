package service

import (
	"context"
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"time"
)

func (s *Service) RolloutRestart(
	x context.Context,
	clusterName string,
	resourceType string,
	name string,
	namespace string,
) error {
	c, e := s.ClusterByName(clusterName)

	if e != nil {
		return e
	}

	if namespace == "" {
		namespace = "default"
	}

	patch := map[string]any{
		"spec": map[string]any{
			"template": map[string]any{
				"metadata": map[string]any{
					"annotations": map[string]any{
						"kubectl.kubernetes.io/restartedAt": time.Now().Format(
							time.RFC3339,
						),
					},
				},
			},
		},
	}
	b, f := json.Marshal(patch)

	if f != nil {
		return f
	}

	apps := c.Clientset().AppsV1()

	switch resourceType {
	case "deployment", "deployments":
		_, g := apps.Deployments(namespace).Patch(
			x,
			name,
			types.StrategicMergePatchType,
			b,
			v1.PatchOptions{},
		)

		return g
	case "daemonset", "daemonsets":
		_, g := apps.DaemonSets(namespace).Patch(
			x,
			name,
			types.StrategicMergePatchType,
			b,
			v1.PatchOptions{},
		)

		return g
	case "statefulset", "statefulsets":
		_, g := apps.StatefulSets(namespace).Patch(
			x,
			name,
			types.StrategicMergePatchType,
			b,
			v1.PatchOptions{},
		)

		return g
	default:
		return validation.New(
			"unsupported resource type for rollout restart: %s (use deployment, daemonset, or statefulset)",
			resourceType,
		)
	}
}
