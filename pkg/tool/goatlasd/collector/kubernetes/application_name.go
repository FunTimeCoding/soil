package kubernetes

import (
	"github.com/funtimecoding/soil/pkg/kubernetes/types/native/pod"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
)

func ApplicationName(p *pod.Pod) string {
	if v := p.Raw.Labels[constant.ApplicationLabel]; v != "" {
		return v
	}

	if v := p.Raw.Labels[constant.ApplicationNameLabel]; v != "" {
		return v
	}

	if len(p.Raw.Spec.Containers) > 0 {
		return p.Raw.Spec.Containers[0].Name
	}

	return p.Name
}
