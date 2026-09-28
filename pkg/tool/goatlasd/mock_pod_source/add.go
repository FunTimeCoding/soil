package mock_pod_source

import (
	"github.com/funtimecoding/soil/pkg/kubernetes/types/native/pod"
	"k8s.io/api/core/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Client) Add(
	name string,
	namespace string,
	node string,
	labels map[string]string,
) {
	c.pods = append(
		c.pods,
		pod.New(
			&v1.Pod{
				ObjectMeta: meta.ObjectMeta{
					Name:      name,
					Namespace: namespace,
					Labels:    labels,
				},
				Spec: v1.PodSpec{
					NodeName:   node,
					Containers: []v1.Container{{Name: name}},
				},
			},
			"",
		),
	)
}
