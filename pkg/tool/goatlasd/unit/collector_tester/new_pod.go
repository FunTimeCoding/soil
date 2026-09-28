package collector_tester

import (
	"github.com/funtimecoding/soil/pkg/kubernetes/types/native/pod"
	"k8s.io/api/core/v1"
	v11 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func NewPod(
	name string,
	labels map[string]string,
	containers []string,
) *pod.Pod {
	var v []v1.Container

	for _, c := range containers {
		v = append(v, v1.Container{Name: c})
	}

	return pod.New(
		&v1.Pod{
			ObjectMeta: v11.ObjectMeta{Name: name, Labels: labels},
			Spec:       v1.PodSpec{Containers: v},
		},
		"",
	)
}
