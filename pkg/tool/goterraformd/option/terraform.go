package option

import (
	"github.com/funtimecoding/soil/pkg/kubernetes/face"
	provisionFace "github.com/funtimecoding/soil/pkg/provision/face"
	"github.com/funtimecoding/soil/pkg/web/option"
)

type Terraform struct {
	option.Web
	Metric            option.Web
	Repository        string
	ClonePath         string
	TerraformPath     string
	Executable        string
	Hostname          string
	Namespace         string
	PostgresLocator   string
	LitePath          string
	StateNamespace    string
	StateLease        string
	Downstream        []string
	DownstreamToken   string
	Cluster           func() face.KubernetesClient
	DownstreamClients func() []provisionFace.Client
}
