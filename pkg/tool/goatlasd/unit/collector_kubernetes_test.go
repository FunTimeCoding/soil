package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/kubernetes/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/unit/collector_tester"
	"testing"
)

func TestKubernetesPlacesOneRowPerNode(t *testing.T) {
	f := mock_client.New()
	collector_tester.AddPod(
		f,
		"alloy-a",
		"alloy",
		"delta",
		map[string]string{"app": "alloy"},
	)
	collector_tester.AddPod(
		f,
		"alloy-b",
		"alloy",
		constant.FixtureVirtualMachineNode,
		map[string]string{"app": "alloy"},
	)
	v := collector_tester.CollectPods(t, f)
	assert.Count(t, 2, v)
	assert.String(t, "alloy", v[0].Name)
	assert.String(t, "delta", v[0].PlaceName)
	assert.String(t, "echo", v[1].PlaceName)
}

func TestKubernetesCollapsesReplicasOnOneNode(t *testing.T) {
	f := mock_client.New()
	labels := map[string]string{"app": constant.FixtureService}
	collector_tester.AddPod(
		f,
		"relay-1",
		constant.FixtureScope,
		"delta",
		labels,
	)
	collector_tester.AddPod(
		f,
		"relay-2",
		constant.FixtureScope,
		"delta",
		labels,
	)
	v := collector_tester.CollectPods(t, f)
	assert.Count(t, 1, v)
	assert.String(t, "relay", v[0].Name)
}

func TestKubernetesSeparatesEqualNamesByNamespace(t *testing.T) {
	f := mock_client.New()
	collector_tester.AddPod(
		f,
		"redis-1",
		"storage",
		"delta",
		map[string]string{"app": "redis"},
	)
	collector_tester.AddPod(
		f,
		"redis-2",
		"media",
		"delta",
		map[string]string{"app": "redis"},
	)
	v := collector_tester.CollectPods(t, f)
	assert.Count(t, 2, v)
}

func TestKubernetesLeavesUnknownNodeUnplaced(t *testing.T) {
	f := mock_client.New()
	collector_tester.AddPod(
		f,
		"stray",
		"default",
		"hotel",
		map[string]string{"app": "stray"},
	)
	v := collector_tester.CollectPods(t, f)
	assert.Count(t, 1, v)
	assert.String(t, "", v[0].PlaceKind)
	assert.String(t, "", v[0].PlaceName)
}

func TestKubernetesCarriesSourceAndKind(t *testing.T) {
	f := mock_client.New()
	collector_tester.AddPod(
		f,
		"relay-1",
		constant.FixtureScope,
		"delta",
		map[string]string{"app": constant.FixtureService},
	)
	v := collector_tester.CollectPods(t, f)
	assert.String(t, "kubernetes", v[0].Source)
	assert.String(t, "service", v[0].Kind)
	assert.String(t, "relay", v[0].Scope)
}
