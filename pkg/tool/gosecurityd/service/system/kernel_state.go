package system

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/service/system/matched"
	"slices"
)

func kernelState(m *matched.Matched) string {
	if !m.KernelChecked {
		return constant.KernelUnknown
	}

	result := constant.KernelCurrent

	for _, f := range m.Findings {
		if !f.Kernel || f.Fix == constant.FindingNone {
			continue
		}

		state := constant.KernelPaid

		switch {
		case f.Reboot:
			state = constant.KernelReboot
		case f.Upgradable:
			state = constant.KernelUpgrade
		case f.Fix == constant.FindingDebian:
			state = constant.KernelStuck
		}

		if slices.Index(constant.KernelStateOrder, state) <
			slices.Index(constant.KernelStateOrder, result) {
			result = state
		}
	}

	return result
}
