package alert

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"github.com/prometheus/alertmanager/api/v2/models"
	"github.com/prometheus/common/model"
	"testing"
)

func TestGroupByInstance(t *testing.T) {
	assert.Any(
		t,
		map[string][]*Alert{
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
		GroupByInstance(
			[]*Alert{
				{
					Name:        upper.Alfa,
					Fingerprint: "fingerprint1",
					Labels: models.LabelSet{
						model.InstanceLabel: "instance1",
					},
				},
				{
					Name:        upper.Bravo,
					Fingerprint: "fingerprint2",
					Labels: models.LabelSet{
						model.InstanceLabel: "instance1",
					},
				},
				{
					Name:        upper.Charlie,
					Fingerprint: "fingerprint3",
					Labels: models.LabelSet{
						model.InstanceLabel: "instance2",
					},
				},
			},
		),
	)
}
