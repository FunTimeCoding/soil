package unit

import "github.com/funtimecoding/soil/pkg/tool/gokubernetesd/constant"

func writeTools() []string {
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
