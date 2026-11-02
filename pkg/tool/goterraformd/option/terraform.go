package option

import (
	"github.com/funtimecoding/soil/pkg/kubernetes/face"
	provisionFace "github.com/funtimecoding/soil/pkg/provision/face"
)

type Terraform struct {
	Address           string
	ServiceTokens     []string
	MetricAddress     string
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
	Cluster           func() face.Client
	DownstreamClients func() []provisionFace.Client
}
