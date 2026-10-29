package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/kubernetes/client/refusal"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/apply_result"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/cluster"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/request"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func applyResourceOn(
	x context.Context,
	c *cluster.Cluster,
	q request.Apply,
	resource *unstructured.Unstructured,
) (*apply_result.Result, error) {
	kind := resource.GetKind()
	name := resource.GetName()
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
			return nil, conflict.Format(
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

		h = refusal.Classify(h)

		if not_found.Is(h) {
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
		return nil, refusal.Classify(h)
	}

	return apply_result.New(kind, name, namespace), nil
}
