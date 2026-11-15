package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"k8s.io/client-go/rest"
)

func refusingConfiguration() *rest.Config {
	p, l := system.ClaimPort()
	errors.PanicClose(l)

	return &rest.Config{
		Host: locator.New(webConstant.Localhost).Port(p).Insecure().String(),
	}
}
