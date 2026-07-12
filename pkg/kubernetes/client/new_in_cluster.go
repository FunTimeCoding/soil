package client

import "github.com/funtimecoding/soil/pkg/errors"

func NewInCluster(cluster string) *Client {
	result, e := TryInCluster(cluster)
	errors.PanicOnError(e)

	return result
}
