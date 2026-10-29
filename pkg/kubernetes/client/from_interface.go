package client

import (
	"context"
	"k8s.io/client-go/kubernetes"
)

func FromInterface(
	k kubernetes.Interface,
	cluster string,
) *Client {
	result := Stub()
	result.context = context.Background()
	result.client = k
	result.cluster = cluster

	return result
}
