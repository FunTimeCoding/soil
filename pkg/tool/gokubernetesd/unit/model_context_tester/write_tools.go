package model_context_tester

import "github.com/funtimecoding/soil/pkg/tool/gokubernetesd/constant"

func WriteTools() []string {
	return []string{
		constant.ExecInPod,
		constant.Apply,
		constant.Patch,
		constant.Delete,
		constant.RolloutRestart,
		constant.PortForward,
		constant.StopPortForward,
	}
}
