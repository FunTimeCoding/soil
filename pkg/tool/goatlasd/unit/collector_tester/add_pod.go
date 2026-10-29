package collector_tester

import (
	"github.com/funtimecoding/soil/pkg/kubernetes/client/mock_client"
	"github.com/funtimecoding/soil/pkg/kubernetes/types/native/pod"
	core "k8s.io/api/core/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func AddPod(
	f *mock_client.Client,
	name string,
	namespace string,
	node string,
	labels map[string]string,
) {
	f.PodList = append(
		f.PodList,
		pod.New(
			&core.Pod{
				ObjectMeta: meta.ObjectMeta{
					Name:      name,
					Namespace: namespace,
					Labels:    labels,
				},
				Spec: core.PodSpec{
					NodeName:   node,
					Containers: []core.Container{{Name: name}},
				},
			},
			"",
		),
	)
}
