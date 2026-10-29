package cluster

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func NewWithResources(
	name string,
	clientset kubernetes.Interface,
	dynamic dynamic.Interface,
	configuration *rest.Config,
	resources []*v1.APIResourceList,
) *Cluster {
	return &Cluster{
		name:          name,
		clientset:     clientset,
		dynamic:       dynamic,
		configuration: configuration,
		resources:     resources,
	}
}
