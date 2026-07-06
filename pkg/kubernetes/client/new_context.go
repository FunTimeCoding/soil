package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/kubernetes/client/client"
	"github.com/funtimecoding/soil/pkg/kubernetes/client/client_configuration"
	"github.com/funtimecoding/soil/pkg/kubernetes/client/metrics"
	"github.com/funtimecoding/soil/pkg/kubernetes/client/operation"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

func NewContext(cluster string) (*Client, error) {
	c, configurationFail := client_configuration.NewContext(cluster)

	if configurationFail != nil {
		return nil, configurationFail
	}

	d, dynamicFail := client.NewDynamic(c)

	if dynamicFail != nil {
		return nil, dynamicFail
	}

	result := Stub()
	result.context = context.Background()
	result.configuration = c
	result.client = client.New(c)
	result.metric = metrics.New(c)
	result.cluster = cluster
	result.dynamic = d
	_, nodesFail := operation.Node(result.client).List(
		result.context,
		v1.ListOptions{},
	)

	if nodesFail != nil {
		return nil, nodesFail
	}

	return result, nil
}
