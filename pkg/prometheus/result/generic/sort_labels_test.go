package generic

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/prometheus/constant"
	"testing"
)

func TestSortLabels(t *testing.T) {
	assert.Strings(
		t,
		[]string{"namespace", "pod", "container"},
		sortLabels(
			[]string{
				constant.ContainerLabel,
				constant.PodLabel,
				constant.NamespaceLabel,
			},
		),
	)
}
