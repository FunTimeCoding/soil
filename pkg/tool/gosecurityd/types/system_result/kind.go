package system_result

import "github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"

func (f *Finding) Kind() string {
	switch {
	case f.Reboot:
		return constant.KindReboot
	case f.Upgradable && f.Kernel:
		return constant.KindKernelUpgrade
	case f.Upgradable:
		return constant.KindUpgrade
	case f.Fix == constant.FindingDebian:
		return constant.KindPending
	case f.Fix == constant.FindingPaid:
		return constant.KindPaid
	}

	return constant.KindOpen
}
