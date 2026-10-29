package cluster

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"sync"
)

type Cluster struct {
	name          string
	clientset     kubernetes.Interface
	dynamic       dynamic.Interface
	configuration *rest.Config
	mutex         sync.Mutex
	resources     []*v1.APIResourceList
}
