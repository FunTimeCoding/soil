package cluster

import (
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func New(
	name string,
	clientset kubernetes.Interface,
	dynamic dynamic.Interface,
	configuration *rest.Config,
) *Cluster {
	return &Cluster{
		name:          name,
		clientset:     clientset,
		dynamic:       dynamic,
		configuration: configuration,
	}
}
