package jenkins

import (
	"context"
	"github.com/bndr/gojenkins"
	"github.com/funtimecoding/soil/pkg/jenkins/basic"
)

type Client struct {
	basic   *basic.Client
	context context.Context
	client  *gojenkins.Jenkins
}
