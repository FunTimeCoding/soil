package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/prometheus/constant"
	"github.com/funtimecoding/soil/pkg/prometheus/result/generic"
	"testing"
)

func TestSortLabels(t *testing.T) {
	assert.Strings(
		t,
		[]string{"namespace", "pod", "container"},
		generic.SortLabels(
			[]string{
				constant.ContainerLabel,
				constant.PodLabel,
				constant.NamespaceLabel,
			},
		),
	)
}
