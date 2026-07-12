package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/kubernetes/client/client"
	"github.com/funtimecoding/soil/pkg/kubernetes/client/metrics"
	"k8s.io/client-go/rest"
)

func fromConfiguration(
	configuration *rest.Config,
	cluster string,
) (*Client, error) {
	d, e := client.NewDynamic(configuration)

	if e != nil {
		return nil, e
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
