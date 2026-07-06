package service

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/format"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/node_capacity"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/response"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sort"
	"strings"
)

var metricsNodeGVR = schema.GroupVersionResource{
	Group:    "metrics.k8s.io",
	Version:  "v1beta1",
	Resource: "nodes",
}

func (s *Service) TopNodes(
	x context.Context,
	clusterName string,
) ([]response.NodeMetrics, error) {
	c, e := s.ClusterByName(clusterName)

	if e != nil {
		return nil, e
	}

	metrics, f := c.Dynamic().Resource(metricsNodeGVR).List(
		x,
		v1.ListOptions{},
	)

	if f != nil {
		if strings.Contains(
			f.Error(),
			"could not find the requested resource",
		) {
			return nil, fmt.Errorf("metrics API not available - install metrics-server")
		}

		return nil, f
	}

	nodes, g := c.Clientset().CoreV1().Nodes().List(x, v1.ListOptions{})

	if g != nil {
		return nil, g
	}

	allocatable := make(map[string]*node_capacity.NodeCapacity)

	for _, n := range nodes.Items {
		cpu := n.Status.Allocatable.Cpu()
		memory := n.Status.Allocatable.Memory()
		allocatable[n.Name] = node_capacity.New(
			cpu.MilliValue(),
			memory.Value(),
		)
	}

	result := []response.NodeMetrics{}

	for _, item := range metrics.Items {
		name := item.GetName()
		usage, _ := item.Object["usage"].(map[string]any)

		if usage == nil {
			continue
		}

		cpuRaw, _ := usage["cpu"].(string)
		memRaw, _ := usage["memory"].(string)
		cpuMillis := format.QuantityMillis(cpuRaw)
		memBytes := format.QuantityBytes(memRaw)
		entry := response.NodeMetrics{
			Name:   name,
			CPU:    format.MilliCores(cpuMillis),
			Memory: format.MiB(memBytes),
		}

		if cap, okay := allocatable[name]; okay {
			if cap.CpuMillis > 0 {
				entry.CPUPercent = fmt.Sprintf(
					"%d%%",
					cpuMillis*100/cap.CpuMillis,
				)
			}

			if cap.MemoryBytes > 0 {
				entry.MemoryPercent = fmt.Sprintf(
					"%d%%",
					memBytes*100/cap.MemoryBytes,
				)
			}
		}

		result = append(result, entry)
	}

	sort.Slice(
		result,
		func(i, j int) bool {
			return result[i].Name < result[j].Name
		},
	)

	return result, nil
}
