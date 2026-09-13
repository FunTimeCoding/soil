package directory

import (
	"github.com/funtimecoding/soil/pkg/directory/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func NewEnvironment() *Client {
	result := New(
		environment.Required(constant.HostEnvironment),
		environment.Required(constant.BaseEnvironment),
	).WithCredential(
		environment.Required(constant.BindEnvironment),
		environment.Required(constant.PasswordEnvironment),
	).WithFilter(
		environment.Required(constant.UserFilterEnvironment),
		environment.Required(constant.GroupFilterEnvironment),
	)

	if environment.Exists(constant.AuthorityEnvironment) {
		result = result.WithAuthority(
			environment.Required(constant.AuthorityEnvironment),
		)
	}

	if environment.Exists(constant.InsecureEnvironment) {
		result = result.WithInsecure()
	}

	if environment.Exists(constant.UntrustedEnvironment) {
		result = result.WithUntrusted()
	}

	if environment.Exists(constant.PortEnvironment) {
		result = result.WithPort(
			environment.RequiredInteger(constant.PortEnvironment),
		)
	}

	return result
}
