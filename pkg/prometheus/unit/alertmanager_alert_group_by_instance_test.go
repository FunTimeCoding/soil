package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert"
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/prometheus/alertmanager/api/v2/models"
	"github.com/prometheus/common/model"
	"testing"
)

func TestGroupByInstance(t *testing.T) {
	assert.Any(
		t,
		map[string][]*alert.Alert{
			"instance1": {
				{
					Name:        "Alfa",
					Fingerprint: "fingerprint1",
					Labels:      models.LabelSet{"instance": "instance1"},
				},
				{
					Name:        "Bravo",
					Fingerprint: "fingerprint2",
					Labels:      models.LabelSet{"instance": "instance1"},
				},
			},
			"instance2": {
				{
					Name:        "Charlie",
					Fingerprint: "fingerprint3",
					Labels:      models.LabelSet{"instance": "instance2"},
				},
			},
		},
		alert.GroupByInstance(
			[]*alert.Alert{
				{
					Name:        constant.UpperAlfa,
					Fingerprint: "fingerprint1",
					Labels: models.LabelSet{
						model.InstanceLabel: "instance1",
					},
				},
				{
					Name:        constant.UpperBravo,
					Fingerprint: "fingerprint2",
					Labels: models.LabelSet{
						model.InstanceLabel: "instance1",
					},
				},
				{
					Name:        constant.UpperCharlie,
					Fingerprint: "fingerprint3",
					Labels: models.LabelSet{
						model.InstanceLabel: "instance2",
					},
				},
			},
		),
	)
}
