package system_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/time"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/system_result"
	"github.com/funtimecoding/soil/pkg/web/component"
	"github.com/funtimecoding/soil/pkg/web/component/fact"
	"maps"
	"maragu.dev/gomponents"
	"slices"
	"strconv"
)

func Facts(v *system_result.Summary) gomponents.Node {
	labels := make([]string, 0, len(v.Labels))

	for _, k := range slices.Sorted(maps.Keys(v.Labels)) {
		labels = append(labels, join.Equals([]string{k, v.Labels[k]}))
	}

	verified := "no tracker loaded"

	if !v.TrackerVerified.IsZero() {
		verified = time.FormatCompact(v.TrackerVerified)
	}

	lists := "not reported"

	if v.ListsUpdated != nil {
		lists = time.FormatCompact(*v.ListsUpdated)
	}

	offers := make([]string, 0, len(v.KernelOffers))

	for _, o := range v.KernelOffers {
		offers = append(
			offers,
			fmt.Sprintf(
				"%s (%d: %s)",
				o.Version,
				o.Priority,
				join.CommaSpace(o.Origins),
			),
		)
	}

	return fact.New().
		Row("Release", v.Release).
		RowNode("Coverage", CoverageBadge(v.Coverage)).
		Row("Clean", component.YesNo(v.Clean)).
		RowNode("Vulnerabilities at 9.0 or higher", CriticalBadge(v.Critical)).
		RowNode("Top score", ScoreBadge(v.TopScore)).
		Row("Package findings with a fix", strconv.Itoa(v.Findings-v.NoFix)).
		Row("Package findings without a fix yet", strconv.Itoa(v.NoFix)).
		RowNode("Kernel", KernelBadge(v.KernelState)).
		Row("Kernel detail", v.KernelReason).
		Row("Running kernel", join.Space(v.Kernel, v.RunningKernel)).
		Row("Newest installed kernel", v.NewestKernel).
		Row(
			"Kernel findings with a fix",
			fmt.Sprintf("%d of %d", v.KernelWithFix, v.KernelFindings),
		).
		Row("Kernel offered by", join.SemicolonSpace(offers)).
		Row("Package lists dated", lists).
		Row("Reported", time.FormatCompact(v.Updated)).
		Row("Stale report", component.YesNo(v.Stale)).
		Row("Tracker verified", verified).
		Row("Priority", strconv.Itoa(v.Priority)).
		Row("Labels", join.CommaSpace(labels)).
		Node()
}
