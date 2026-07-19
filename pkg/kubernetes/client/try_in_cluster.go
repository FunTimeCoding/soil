package client

import "k8s.io/client-go/rest"

func TryInCluster(cluster string) (*Client, error) {
	configuration, e := rest.InClusterConfig()

	if e != nil {
		return nil, e
	}

	return FromConfiguration(configuration, cluster)
}
