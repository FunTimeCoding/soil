package system_page

import (
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

	return fact.New().
		Row("Release", v.Release).
		Row("Kernel", v.Kernel).
		RowNode("Coverage", CoverageBadge(v.Coverage)).
		Row("Clean", component.YesNo(v.Clean)).
		Row("Reported", time.FormatCompact(v.Updated)).
		Row("Stale report", component.YesNo(v.Stale)).
		Row("Tracker verified", verified).
		Row("Priority", strconv.Itoa(v.Priority)).
		Row("Labels", join.CommaSpace(labels)).
		Node()
}
