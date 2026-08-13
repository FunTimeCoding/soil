package service

import (
	"context"
	"fmt"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/yaml"
)

func (s *Service) ApplyResource(
	x context.Context,
	clusterName string,
	q ApplyQuery,
) (*ApplyResult, error) {
	c, e := s.ClusterByName(clusterName)

	if e != nil {
		return nil, e
	}

	var object map[string]any
	f := yaml.Unmarshal([]byte(q.Manifest), &object)

	if f != nil {
		return nil, fmt.Errorf("invalid YAML: %s", f)
	}

	resource := &unstructured.Unstructured{Object: object}
	kind := resource.GetKind()

	if kind == "" {
		return nil, fmt.Errorf("manifest must include kind")
	}

	if resource.GetAPIVersion() == "" {
		return nil, fmt.Errorf("manifest must include apiVersion")
	}

	name := resource.GetName()

	if name == "" {
		return nil, fmt.Errorf("manifest must include metadata.name")
	}

	gvr, namespaced, g := c.ResolveResource(kind)

	if g != nil {
		return nil, g
	}

	namespace := q.Namespace

	if namespace == "" {
		namespace = resource.GetNamespace()
	}

	if namespace == "" {
		namespace = "default"
	}

	options := v1.CreateOptions{}

	if q.DryRun {
		options.DryRun = []string{v1.DryRunAll}
	}

	if !q.Override {
		var existing *unstructured.Unstructured
		var h error

		if namespaced {
			existing, h = c.Dynamic().Resource(gvr).Namespace(namespace).Get(
				x,
				name,
				v1.GetOptions{},
			)
		} else {
			existing, h = c.Dynamic().Resource(gvr).Get(
				x,
				name,
				v1.GetOptions{},
			)
		}

		if h == nil && existing != nil {
			return nil, fmt.Errorf(
				"%s/%s already exists in %s - pass override: true to update",
				kind,
				name,
				namespace,
			)
		}
	}

	var h error

	if q.Override {
		updateOptions := v1.UpdateOptions{}

		if q.DryRun {
			updateOptions.DryRun = []string{v1.DryRunAll}
		}

		if namespaced {
			_, h = c.Dynamic().Resource(gvr).Namespace(namespace).Update(
				x,
				resource,
				updateOptions,
			)
		} else {
			_, h = c.Dynamic().Resource(gvr).Update(x, resource, updateOptions)
		}

		if h != nil {
			if namespaced {
				_, h = c.Dynamic().Resource(gvr).Namespace(namespace).Create(
					x,
					resource,
					options,
				)
			} else {
				_, h = c.Dynamic().Resource(gvr).Create(x, resource, options)
			}
		}
	} else {
		if namespaced {
			_, h = c.Dynamic().Resource(gvr).Namespace(namespace).Create(
				x,
				resource,
				options,
			)
		} else {
			_, h = c.Dynamic().Resource(gvr).Create(x, resource, options)
		}
	}

	if h != nil {
		return nil, h
	}

	return &ApplyResult{Kind: kind, Name: name, Namespace: namespace}, nil
}
