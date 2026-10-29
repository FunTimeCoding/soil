package service

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gogated/option"
	"github.com/funtimecoding/soil/pkg/web/authorization/client"
	"github.com/funtimecoding/soil/pkg/web/constant"
)

func authorizationClient(o *option.Option) *client.Client {
	return client.New(
		o.Issuer,
		o.AdminClientIdentifier,
		o.AdminClientSecret,
		constant.SignInPath,
		join.Empty(o.Issuer, constant.CallbackPath),
		client.DeriveKey(o.Secret),
	)
}
