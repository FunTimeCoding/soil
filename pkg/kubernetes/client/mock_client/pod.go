package mock_client

import (
	kubernetesConstant "github.com/funtimecoding/soil/pkg/kubernetes/constant"
	"github.com/funtimecoding/soil/pkg/kubernetes/types/native/pod"
	"github.com/funtimecoding/soil/pkg/strings/join"
	core "k8s.io/api/core/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Client) Pod(
	namespace string,
	name string,
) (*pod.Pod, error) {
	c.mutex.Lock()
	c.lookups = append(c.lookups, join.Slash([]string{namespace, name}))
	c.mutex.Unlock()

	if c.podError != nil {
		return nil, c.podError
	}

	if !c.pods[name] {
		return nil, nil
	}

	return pod.New(
		&core.Pod{
			ObjectMeta: meta.ObjectMeta{Namespace: namespace, Name: name},
		},
		kubernetesConstant.InCluster,
	), nil
}
