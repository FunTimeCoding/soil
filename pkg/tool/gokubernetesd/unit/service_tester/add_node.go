package service_tester

import (
	core "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (t *Tester) AddNode(
	name string,
	cpuCores int64,
	memoryGiB int64,
) {
	_, e := t.Clientset.CoreV1().Nodes().Create(
		t.context(),
		&core.Node{
			ObjectMeta: v1.ObjectMeta{Name: name},
			Status: core.NodeStatus{
				Allocatable: core.ResourceList{
					core.ResourceCPU: *resource.NewMilliQuantity(
						cpuCores*1000,
						resource.DecimalSI,
					),
					core.ResourceMemory: *resource.NewQuantity(
						memoryGiB*1024*1024*1024,
						resource.BinarySI,
					),
				},
			},
		},
		v1.CreateOptions{},
	)

	if e != nil {
		panic(e)
	}
}
