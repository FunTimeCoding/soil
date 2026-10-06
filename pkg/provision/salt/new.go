package salt

import "github.com/funtimecoding/soil/pkg/provision/salt/basic"

func New(
	host string,
	port int,
	user string,
	password string,
	o ...Option,
) *Client {
	result := &Client{}

	for _, f := range o {
		f(result)
	}

	eauth := result.authentication

	if eauth == "" {
		eauth = "pam"
	}

	result.basic = basic.New(
		newLocator(host, port, result.insecure),
		user,
		password,
		eauth,
	)

	return result
}
