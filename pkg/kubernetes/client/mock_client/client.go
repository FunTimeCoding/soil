package mock_client

import (
	"github.com/funtimecoding/soil/pkg/failure"
	"github.com/funtimecoding/soil/pkg/kubernetes/types/custom/certificate"
	"github.com/funtimecoding/soil/pkg/kubernetes/types/native/lease"
	"github.com/funtimecoding/soil/pkg/kubernetes/types/native/pod"
	"k8s.io/api/networking/v1"
	"sync"
)

type Client struct {
	Failure        error
	PodList        []*pod.Pod
	lease          *lease.Lease
	pods           map[string]bool
	podError       error
	lookups        []string
	failures       *failure.Failure
	certificates   []*certificate.Certificate
	triggered      []string
	issueOnTrigger bool
	policies       map[string]*v1.NetworkPolicy
	mutex          sync.Mutex
}
