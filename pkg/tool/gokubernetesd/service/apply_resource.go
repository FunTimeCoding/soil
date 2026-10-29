package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/apply_result"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/cluster"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/request"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/yaml"
)

func (s *Service) ApplyResource(
	x context.Context,
	clusterName string,
	q request.Apply,
) (*apply_result.Result, error) {
	var object map[string]any

	if e := yaml.Unmarshal([]byte(q.Manifest), &object); e != nil {
		return nil, validation.New("invalid YAML: %s", e)
	}

	resource := &unstructured.Unstructured{Object: object}

	if resource.GetKind() == "" {
		return nil, validation.New("manifest must include kind")
	}

	if resource.GetAPIVersion() == "" {
		return nil, validation.New("manifest must include apiVersion")
	}

	if resource.GetName() == "" {
		return nil, validation.New("manifest must include metadata.name")
	}

	return guarded(
		s,
		x,
		clusterName,
		func(c *cluster.Cluster) (*apply_result.Result, error) {
			return applyResourceOn(x, c, q, resource)
		},
	)
}
