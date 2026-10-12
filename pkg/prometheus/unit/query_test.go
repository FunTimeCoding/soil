package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/prometheus/constant"
	"github.com/funtimecoding/soil/pkg/prometheus/query"
	"github.com/funtimecoding/soil/pkg/prometheus/query/filter"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"testing"
)

func TestQueryFilter(t *testing.T) {
	assert.String(
		t,
		`up{container="Charlie",namespace="Bravo",node="Alfa"}`,
		query.Filter(
			constant.Up,
			filter.New().Node(stringsConstant.UpperAlfa).Namespace(
				stringsConstant.UpperBravo,
			).Container(
				stringsConstant.UpperCharlie,
			),
		),
	)
}

func TestNullFilter(t *testing.T) {
	assert.String(t, "up", query.Filter(constant.Up, nil))
}

func TestSumBy(t *testing.T) {
	assert.String(
		t,
		`sum(kube_pod_container_status_restarts_total{namespace="Alfa"}) by (namespace, pod)`,
		query.SumBy(
			constant.Restart,
			filter.New().Namespace(stringsConstant.UpperAlfa),
			constant.NamespaceLabel,
			constant.PodLabel,
		),
	)
}

func TestSumByRate(t *testing.T) {
	assert.String(
		t,
		`sum(rate(kube_pod_container_status_restarts_total{namespace="Alfa"}[5m])) by (namespace, pod)`,
		query.SumByRate(
			constant.Restart,
			filter.New().Namespace(stringsConstant.UpperAlfa),
			5,
			constant.NamespaceLabel,
			constant.PodLabel,
		),
	)
}

func TestFilterLike(t *testing.T) {
	assert.String(
		t,
		`up{instance=~"10.*"}`,
		query.Filter(constant.Up, filter.New().InstanceLike("10.*")),
	)
}

func TestChanges(t *testing.T) {
	assert.String(
		t,
		`changes(kube_pod_container_status_restarts_total{pod="Alfa"}[5m]) > 0`,
		query.Changes(
			constant.Restart,
			filter.New().Pod(stringsConstant.UpperAlfa),
			5,
		),
	)
}

func TestSumByLabelReplace(t *testing.T) {
	assert.String(
		t,
		`sum by (owner) (label_replace(up{instance="Alfa"}, "owner", "$1", "namespace", "([^-]*)-.*"))`,
		query.SumByLabelReplace(
			query.Filter(
				constant.Up,
				filter.New().Instance(stringsConstant.UpperAlfa),
			),
			constant.NamespaceLabel,
			"owner",
			"([^-]*)-.*",
		),
	)
}
