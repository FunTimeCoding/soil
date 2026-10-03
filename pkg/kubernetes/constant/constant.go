package constant

import (
	"github.com/funtimecoding/soil/pkg/console/constant"
	"regexp"
)

const (
	ContextEnvironment     = "KUBERNETES_CONTEXT"
	AutoCleanupEnvironment = "KUBERNETES_AUTO_CLEANUP"
)

const NodeAll = ""
const (
	Kubectl        = "kubectl"
	Configuration  = "config"
	CurrentContext = "current-context"
)

var (
	Format = constant.ExtendedColorFormat
	Dense  = constant.ColorFormat
)

const (
	PodsResource       = "pods"
	ExecuteSubResource = "exec"
)

const DNSConfigurationForming = "DNSConfigForming"

var IrrelevantEventReason = []string{DNSConfigurationForming}

const (
	TrivyNamespace = "trivy"
	TrivyCron      = "trivy"

	RenovateNamespace = "renovate"
	LabCron           = "lab"
	HubCron           = "hub"

	ManualJob    = "manual"
	ManualLabJob = "manual-lab"
	ManualHubJob = "manual-hub"
)

// Reference: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#dns-label-names
var NameExpression = regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`)

const (
	TrivyArgument = "trivy"
	LabArgument   = "lab"
	HubArgument   = "hub"
)

const InCluster = "in-cluster"
