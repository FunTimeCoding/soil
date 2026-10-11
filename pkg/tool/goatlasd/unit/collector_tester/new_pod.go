package collector_tester

import (
	"github.com/funtimecoding/soil/pkg/kubernetes/types/native/pod"
	core "k8s.io/api/core/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func NewPod(
	name string,
	labels map[string]string,
	containers []string,
) *pod.Pod {
	var v []core.Container

	for _, c := range containers {
		v = append(v, core.Container{Name: c})
	}

	return pod.New(
		&core.Pod{
			ObjectMeta: meta.ObjectMeta{Name: name, Labels: labels},
			Spec:       core.PodSpec{Containers: v},
		},
		"",
	)
}
