package constant

import "github.com/funtimecoding/soil/pkg/identity"

var Identity = identity.New(
	"gokubernetesd",
	"Kubernetes cluster operations",
	"gokubernetesd",
).WithInstructions(
	"Kubernetes cluster operations - get resources, describe objects, read logs, view events, exec into pods. Multi-cluster with session-scoped cluster selection. Call list_clusters and use_cluster before other operations.",
)

const (
	ListClusters    = "list_clusters"
	UseCluster      = "use_cluster"
	Get             = "get"
	Describe        = "describe"
	Logs            = "logs"
	Events          = "events"
	ExecInPod       = "exec_in_pod"
	Delete          = "delete"
	RolloutRestart  = "rollout_restart"
	PortForward     = "port_forward"
	StopPortForward = "stop_port_forward"
	Top             = "top"
	MuteEvent       = "mute_event"
	UnmuteEvent     = "unmute_event"
	ListMutedEvents = "list_muted_events"
	Argocd          = "argocd"
	Certificates    = "certificates"
	Patch           = "patch"
	Apply           = "apply"
)
