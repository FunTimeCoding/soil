package constant

import (
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"regexp"
)

const (
	ContextEnvironment     = "KUBERNETES_CONTEXT"
	AutoCleanupEnvironment = "KUBERNETES_AUTO_CLEANUP"
)

// NodeAll analogous to NamespaceAll
const NodeAll = ""
const (
	Kubectl        = "kubectl"
	Configuration  = "config"
	CurrentContext = "current-context"
)

var (
	Format = option.ExtendedColor
	Dense  = option.Color
)

const (
	PodsResource       = "pods"
	ExecuteSubResource = "exec"
)

const DNSConfigurationForming = "DNSConfigForming" // Event reason

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

var NameExpression = regexp.MustCompile(
	`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`,
) // DNS-1123
