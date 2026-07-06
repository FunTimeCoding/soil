package constant

import (
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/console/status/tag"
	"github.com/prometheus/common/model"
)

const Graph = "/graph"
const Graph0Expression = "g0.expr"
const (
	HostEnvironment     = "PROMETHEUS_HOST"
	PortEnvironment     = "PROMETHEUS_PORT"
	UserEnvironment     = "PROMETHEUS_USER"
	PasswordEnvironment = "PROMETHEUS_PASSWORD"
	InsecureEnvironment = "PROMETHEUS_INSECURE"
)

// Metric
const (
	Up      = "up"
	Restart = "kube_pod_container_status_restarts_total"
	Load1   = "node_load1"
	Load5   = "node_load5"
	Load15  = "node_load15"
)

// Query result type
const (
	Matrix = "matrix"
	Vector = "vector"
	Scalar = "scalar"
	String = "string"
)

// Label
const (
	Name          = model.MetricNameLabel
	InstanceLabel = model.InstanceLabel
	TargetLabel   = "target"
	JobLabel      = model.JobLabel

	ContainerLabel   = "container"
	DaemonSetLabel   = "daemonset"
	DeploymentLabel  = "deployment"
	EndpointLabel    = "endpoint"
	NamespaceLabel   = "namespace"
	NodeLabel        = "node"
	PodLabel         = "pod"
	ScopeLabel       = "scope"
	ServiceLabel     = "service"
	StateLabel       = "state"
	StatefulSetLabel = "statefulset"
)

var Format = option.Color.Copy().Tag(
	tag.Category,
	tag.Emoji,
	tag.Instance,
)
