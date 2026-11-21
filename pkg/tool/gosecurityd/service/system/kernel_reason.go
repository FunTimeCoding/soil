package system

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/system"
	"strings"
	"time"
)

func kernelReason(
	v *system.System,
	running string,
	now time.Time,
) string {
	top := ""

	for _, o := range v.KernelOffers {
		base, _ := backports(o.Version)
		top = highest(top, base)
	}

	if v.KernelMetapackageInstalled != nil && !*v.KernelMetapackageInstalled {
		if top == "" {
			return constant.ReasonNoMetapackage
		}

		return fmt.Sprintf(constant.ReasonNoMetapackageOffer, top)
	}

	if _, number := backports(running); number > 0 {
		offered := false

		for _, o := range v.KernelOffers {
			for _, origin := range o.Origins {
				offered = offered ||
					strings.Contains(origin, constant.BackportsOrigin)
			}
		}

		if len(v.KernelOffers) > 0 && !offered {
			return constant.ReasonNoBackports
		}
	}

	if v.ListsUpdated != nil && now.Sub(*v.ListsUpdated) > constant.ListsOldAfter {
		return fmt.Sprintf(
			constant.ReasonListsOld,
			int(now.Sub(*v.ListsUpdated).Hours()/24),
		)
	}

	if len(v.KernelOffers) == 0 {
		return constant.ReasonUnknown
	}

	return fmt.Sprintf(constant.ReasonOffer, top)
}
