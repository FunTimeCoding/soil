package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/prometheus/label_result"
	"github.com/funtimecoding/soil/pkg/prometheus/query_result"
	"github.com/funtimecoding/soil/pkg/prometheus/result/generic/float"
	"github.com/funtimecoding/soil/pkg/prometheus/result/generic/integer"
	"github.com/funtimecoding/soil/pkg/prometheus/rule/rule_list"
	"github.com/prometheus/client_golang/api/prometheus/v1"
	"time"
)

type Client interface {
	Address() string
	AllLabels() ([]string, error)
	AllMetrics() ([]string, error)
	Close()
	GraphLink(expression string) string
	IntegerExists(
		q string,
		t time.Time,
	) (bool, error)
	LabelNames(
		matches []string,
		since time.Time,
	) (*label_result.Result, error)
	LabelValues(
		label string,
		matches []string,
		since time.Time,
	) (*label_result.Result, error)
	Metadata(metric string) (map[string][]v1.Metadata, error)
	MustAllLabels() []string
	MustAllMetrics() []string
	MustLabelNames(
		matches []string,
		since time.Time,
	) *label_result.Result
	MustMetadata(metric string) map[string][]v1.Metadata
	MustQuery(
		q string,
		t time.Time,
	) *query_result.Result
	MustQueryFloat(
		q string,
		t time.Time,
	) float64
	MustQueryIntegers(
		q string,
		t time.Time,
	) map[string]int
	MustQueryRange(
		q string,
		r v1.Range,
	) *query_result.Result
	MustQueryVector(q string) float64
	MustRules() *rule_list.List
	MustSeries() v1.TSDBResult
	MustStatus() v1.RuntimeinfoResult
	MustTargets() v1.TargetsResult
	Query(
		q string,
		t time.Time,
	) (*query_result.Result, error)
	QueryExemplar(
		q string,
		start time.Time,
		end time.Time,
	) ([]v1.ExemplarQueryResult, error)
	QueryFloat(
		q string,
		t time.Time,
	) (float64, error)
	QueryFloats(
		q string,
		t time.Time,
	) (map[string]float64, error)
	QueryInteger(
		q string,
		t time.Time,
	) (int, error)
	QueryIntegers(
		q string,
		t time.Time,
	) (map[string]int, error)
	QueryRange(
		q string,
		r v1.Range,
	) (*query_result.Result, error)
	QueryStructFloat(
		q string,
		fallback float64,
		t time.Time,
	) (*float.Result, error)
	QueryStructInteger(
		q string,
		fallback int,
		t time.Time,
	) (*integer.Result, error)
	QueryVector(q string) (float64, error)
	Reachable(x context.Context) error
	Rules() (*rule_list.List, error)
	Series() (v1.TSDBResult, error)
	Status() (v1.RuntimeinfoResult, error)
	Targets() (v1.TargetsResult, error)
}
