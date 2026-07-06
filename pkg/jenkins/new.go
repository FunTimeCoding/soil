package jenkins

import (
	"context"
	"github.com/bndr/gojenkins"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/jenkins/basic"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
)

func New(
	host string,
	user string,
	password string,
) *Client {
	errors.FatalOnEmpty(host, "host")
	errors.FatalOnEmpty(user, "user")
	errors.FatalOnEmpty(password, "password")
	p := constant.ListenPort
	c := gojenkins.CreateJenkins(
		nil,
		locator.New(host).Port(p).Insecure().String(),
		user,
		password,
	)
	x := context.Background()
	_, e := c.Init(x)
	errors.PanicOnError(e)

	return &Client{
		basic:   basic.New(host, p, user, password),
		context: x,
		client:  c,
	}
}
