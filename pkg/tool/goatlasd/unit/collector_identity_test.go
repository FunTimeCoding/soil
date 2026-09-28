package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/collector/kubernetes"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/collector/label"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/unit/collector_tester"
	"testing"
)

func TestApplicationNamePrefersAppLabel(t *testing.T) {
	p := collector_tester.NewPod(
		"relay-7d9f",
		map[string]string{"app": "relay", "app.kubernetes.io/name": "other"},
		[]string{"container"},
	)
	assert.String(t, "relay", kubernetes.ApplicationName(p))
}

func TestApplicationNameFallsBackToChartLabel(t *testing.T) {
	p := collector_tester.NewPod(
		"loki-0",
		map[string]string{"app.kubernetes.io/name": "loki"},
		[]string{"container"},
	)
	assert.String(t, "loki", kubernetes.ApplicationName(p))
}

func TestApplicationNameFallsBackToContainer(t *testing.T) {
	p := collector_tester.NewPod("pod-abc", nil, []string{"sidecar"})
	assert.String(t, "sidecar", kubernetes.ApplicationName(p))
}

func TestApplicationNameFallsBackToPodName(t *testing.T) {
	p := collector_tester.NewPod("lonely", nil, nil)
	assert.String(t, "lonely", kubernetes.ApplicationName(p))
}

func TestApplicationNameIgnoresEmptyLabel(t *testing.T) {
	p := collector_tester.NewPod(
		"pod-abc",
		map[string]string{"app": ""},
		[]string{"sidecar"},
	)
	assert.String(t, "sidecar", kubernetes.ApplicationName(p))
}

func TestServiceNamesTakesPrefixedKeys(t *testing.T) {
	v := label.ServiceNames(
		collector_tester.Labels(
			"service.lighting",
			"mesh lighting",
			"service.climate",
			"thermostat",
		),
	)
	assert.Count(t, 2, v)
	assert.String(t, "lighting", v[0])
	assert.String(t, "climate", v[1])
}

func TestServiceNamesIgnoresOtherKeys(t *testing.T) {
	v := label.ServiceNames(
		collector_tester.Labels(
			"owner",
			"admin",
			"service.lighting",
			"mesh lighting",
		),
	)
	assert.Count(t, 1, v)
	assert.String(t, "lighting", v[0])
}

func TestServiceNamesIgnoresBarePrefix(t *testing.T) {
	assert.Count(
		t,
		0,
		label.ServiceNames(
			collector_tester.Labels(constant.ServiceLabelPrefix, "empty"),
		),
	)
}

func TestServiceNamesHandlesAbsentLabels(t *testing.T) {
	assert.Count(t, 0, label.ServiceNames(nil))
}
