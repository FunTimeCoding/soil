package client

import (
	"context"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/metrics/pkg/client/clientset/versioned"
)

type Client struct {
	clients       map[string]*Client
	context       context.Context
	cluster       string
	configuration *rest.Config
	client        kubernetes.Interface
	metric        *versioned.Clientset
	dynamic       dynamic.Interface
	Verbose       bool
}
