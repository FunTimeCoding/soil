package constant

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/prometheus/common/model"
	"time"
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

var Format = constant.ColorFormat.Copy().Tag(
	constant.TagCategory,
	constant.TagEmoji,
	constant.TagInstance,
)

const (
	None        = "none"
	NoComment   = "no comment"
	UnknownRule = "unknown rule"

	AlertmanagerDefaultDuration = 10 * time.Minute

	AlertmanagerHostEnvironment     = "ALERTMANAGER_HOST"
	AlertmanagerUserEnvironment     = "ALERTMANAGER_USER"
	AlertmanagerPasswordEnvironment = "ALERTMANAGER_PASSWORD"
	AlertmanagerInsecureEnvironment = "ALERTMANAGER_INSECURE"

	AmtoolCommand = "amtool"

	AmtoolPath          = "amtool"
	AmtoolConfiguration = "config.yml"

	AmtoolConfigurationPrefix = "config-"

	KubernetesPrefix = "Kube"

	HighMemoryUsage = "HighMemoryUsage" // Test alert name

	Alerts = "/alerts"

	PermanentTag = "#permanent"
)

// Alert label
const (
	AlertnameLabel = model.AlertNameLabel

	SummaryLabel     = "summary"
	DescriptionLabel = "description"
	PrometheusLabel  = "prometheus"
	SeverityLabel    = "severity"
	MessageLabel     = "message"
)

// Alert state
const (
	ActiveState     = "active"
	SuppressedState = "suppressed"
)

// Severity
const (
	CriticalSeverity    = "critical"
	InformationSeverity = "info"
	NoneSeverity        = "none"
	UnknownSeverity     = "unknown"
	WarningSeverity     = "warning"
)

const ExpiredState = "expired"          // Silence state
const NodeNotReady = "KubeNodeNotReady" // Alert name

var (
	AlertmanagerFormat = constant.ColorFormat.Copy().Tag(constant.TagComment)

	AlertStates = []string{ActiveState, SuppressedState}
	Severities  = []string{
		CriticalSeverity,
		WarningSeverity,
		InformationSeverity,
		NoneSeverity,
		UnknownSeverity,
	}
	SevereSeverities = []string{CriticalSeverity, WarningSeverity}
	RedSeverities    = []string{CriticalSeverity}
	YellowSeverities = []string{WarningSeverity}
)

const (
	Silent = "silent"

	NoHost = "no host"
)

const PushgatewayPort = int(9091)

var ErrorNotFound = errors.New("not found")
var ExampleGroups = []string{
	"alertmanager",
	"events",
	"go",
	"kube",
	"net",
	"node",
	"pg",
	"postgres",
	"probe",
	"process",
	"prometheus",
	"promhttp",
	"pushgateway",
	"reloader",
}
