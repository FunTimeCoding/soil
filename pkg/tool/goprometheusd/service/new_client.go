package service

import (
	"github.com/funtimecoding/soil/pkg/prometheus"
	"github.com/funtimecoding/soil/pkg/prometheus/face"
	"github.com/funtimecoding/soil/pkg/tool/goprometheusd/types/instance"
)

func newClient(i *instance.Instance) (face.Client, error) {
	return prometheus.New(i.Host, i.Port, i.Secure, i.User, i.Password), nil
}
