package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/prometheus"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/alert_enricher"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/alert_enricher/enrichment"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/alert_processor"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/field_changer/severity"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/rule_parser"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/mock_client"
	"github.com/funtimecoding/soil/pkg/prometheus/client"
	"github.com/funtimecoding/soil/pkg/prometheus/loki"
	"github.com/funtimecoding/soil/pkg/prometheus/loki/basic"
	"github.com/funtimecoding/soil/pkg/prometheus/push"
	"github.com/funtimecoding/soil/pkg/prometheus/query/filter"
	"github.com/funtimecoding/soil/pkg/prometheus/result/metric"
	"github.com/funtimecoding/soil/pkg/prometheus/round_tripper"
	"github.com/funtimecoding/soil/pkg/prometheus/rule/rule_list"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"testing"
)

func TestConstructors(t *testing.T) {
	assert.NotNil(t, enrichment.New(upper.Alfa, upper.Bravo, upper.Charlie))
	assert.NotNil(t, alert_enricher.New())
	assert.NotNil(t, alert_processor.New(nil, nil, nil, nil, nil, nil))
	assert.NotNil(t, severity.New(upper.Alfa, upper.Bravo, upper.Charlie))
	assert.NotNil(t, rule_parser.New(nil))
	assert.NotNil(t, mock_client.New())
	assert.NotNil(t, client.New("", nil))
	assert.NotNil(
		t,
		basic.New(upper.Alfa, upper.Bravo, upper.Charlie, false),
	)
	assert.NotNil(
		t,
		loki.New(upper.Alfa, upper.Bravo, upper.Charlie, false),
	)
	assert.NotNil(
		t,
		prometheus.New(
			constant.Localhost,
			0,
			false,
			"",
			"",
			"",
		),
	)
	assert.NotNil(t, push.New(upper.Alfa, 0, false, upper.Bravo))
	assert.NotNil(t, filter.New())
	assert.NotNil(t, metric.New(upper.Alfa))
	assert.NotNil(t, round_tripper.New("", ""))
	assert.NotNil(t, rule_list.New())
}
