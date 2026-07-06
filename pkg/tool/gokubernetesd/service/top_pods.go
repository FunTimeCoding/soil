package service

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/format"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/response"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sort"
	"strings"
)

var metricsPodGVR = schema.GroupVersionResource{
	Group:    "metrics.k8s.io",
	Version:  "v1beta1",
	Resource: "pods",
}

func (s *Service) TopPods(
	x context.Context,
	clusterName string,
	q TopQuery,
) ([]response.PodMetrics, error) {
	c, e := s.ClusterByName(clusterName)

	if e != nil {
		return nil, e
	}

	namespace := q.Namespace

	if namespace == "" {
		namespace = "default"
	}

	var metrics *unstructured.UnstructuredList
	var f error

	if q.AllNamespaces {
		metrics, f = c.Dynamic().Resource(metricsPodGVR).List(
			x,
			v1.ListOptions{},
		)
	} else {
		metrics, f = c.Dynamic().Resource(metricsPodGVR).Namespace(
			namespace,
		).List(x, v1.ListOptions{})
	}

	if f != nil {
		if strings.Contains(
			f.Error(),
			"could not find the requested resource",
		) {
			return nil, fmt.Errorf("metrics API not available - install metrics-server")
		}

		return nil, f
	}

	result := []response.PodMetrics{}

	for _, item := range metrics.Items {
		containers, _ := item.Object["containers"].([]any)
		var totalCPU int64
		var totalMem int64

		for _, raw := range containers {
			c, okay := raw.(map[string]any)

			if !okay {
				continue
			}

			usage, _ := c["usage"].(map[string]any)

			if usage == nil {
				continue
			}

			cpuRaw, _ := usage["cpu"].(string)
			memRaw, _ := usage["memory"].(string)
			totalCPU += format.QuantityMillis(cpuRaw)
			totalMem += format.QuantityBytes(memRaw)
		}

		result = append(
			result,
			response.PodMetrics{
				Name:      item.GetName(),
				Namespace: item.GetNamespace(),
				CPU:       format.MilliCores(totalCPU),
				Memory:    format.MiB(totalMem),
				CpuMillis: totalCPU,
			},
		)
	}

	sort.Slice(
		result,
		func(i, j int) bool {
			return result[i].CpuMillis > result[j].CpuMillis
		},
	)

	return result, nil
}
