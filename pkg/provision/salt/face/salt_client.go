package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/provision/salt/basic/response"
	"github.com/funtimecoding/soil/pkg/provision/salt/basic/response/local_return"
)

type SaltClient interface {
	AcceptKey(minion string) ([]string, error)
	DeleteKey(minion string) ([]string, error)
	Highstate(target string) (map[string]local_return.LocalReturn, error)
	Job(identifier string) (*response.Job, error)
	Jobs() ([]response.Job, error)
	Keys() (*response.KeysReturn, error)
	Local(
		glob string,
		function string,
		a []string,
	) (map[string]local_return.LocalReturn, error)
	LocalAsync(
		target string,
		function string,
		a []string,
	) (string, error)
	Login(x context.Context) error
	Minions() ([]response.Minion, error)
	Reachable(x context.Context) error
	Address() string
	Close()
}
