package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/kubernetes/client/client"
	"github.com/funtimecoding/soil/pkg/kubernetes/client/metrics"
	"k8s.io/client-go/rest"
)

func TryInCluster(cluster string) (*Client, error) {
	configuration, e := rest.InClusterConfig()

	if e != nil {
		return nil, e
	}

	d, f := client.NewDynamic(configuration)

	if f != nil {
		return nil, f
	}

	result := Stub()
	result.context = context.Background()
	result.configuration = configuration
	result.client = client.New(configuration)
	result.metric = metrics.New(configuration)
	result.dynamic = d
	result.cluster = cluster

	return result, nil
}
