package mock_pod_source

import "github.com/funtimecoding/soil/pkg/kubernetes/types/native/pod"

type Client struct {
	pods []*pod.Pod
}
