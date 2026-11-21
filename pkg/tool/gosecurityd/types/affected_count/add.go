package affected_count

import "github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"

func (c *Count) Add(kind string) {
	c.Systems++

	switch kind {
	case constant.KindReboot:
		c.Reboot++
	case constant.KindUpgrade, constant.KindKernelUpgrade:
		c.Upgrade++
	case constant.KindPending:
		c.Pending++
	case constant.KindPaid:
		c.Paid++
	default:
		c.Open++
	}
}
