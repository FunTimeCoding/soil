package face

import (
	"github.com/funtimecoding/soil/pkg/kubernetes/filter"
	"github.com/funtimecoding/soil/pkg/kubernetes/types/native/pod"
)

type PodSource interface {
	Pods(f *filter.Filter) []*pod.Pod
}
