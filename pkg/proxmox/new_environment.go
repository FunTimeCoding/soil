package proxmox

import (
	"github.com/funtimecoding/soil/pkg/proxmox/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"time"
)

func NewEnvironment() *Client {
	var option []Option
	user := environment.Optional(constant.UserEnvironment)
	password := environment.Optional(constant.PasswordEnvironment)

	if user != "" && password != "" {
		option = append(option, WithUser(user, password))
	}

	token := environment.Optional(constant.TokenEnvironment)
	secret := environment.Optional(constant.SecretEnvironment)

	if token != "" && secret != "" {
		option = append(option, WithToken(token, secret))
	}

	if environment.Exists(constant.InsecureEnvironment) {
		option = append(option, WithInsecure())
	}

	if environment.Exists(constant.LogEnvironment) {
		option = append(option, WithLog())
	}

	if environment.Exists(constant.VerboseEnvironment) {
		option = append(option, WithVerbose())
	}

	if raw := environment.Optional(constant.TimeoutEnvironment); raw != "" {
		if d, e := time.ParseDuration(raw); e == nil {
			option = append(option, WithTimeout(d))
		}
	}

	return New(environment.Required(constant.HostEnvironment), option...)
}
