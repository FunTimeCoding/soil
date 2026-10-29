package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/client"
)

type Client interface {
	face.Reachable
	Address() string
	Leases(
		x context.Context,
		phrase string,
	) ([]client.Lease, error)
	Hosts(
		x context.Context,
		phrase string,
	) ([]client.Host, error)
	Pools(
		x context.Context,
		phrase string,
	) ([]client.Pool, error)
	Rules(
		x context.Context,
		phrase string,
	) ([]client.Rule, error)
	Aliases(
		x context.Context,
		phrase string,
	) ([]client.Alias, error)
	SourceNatRules(
		x context.Context,
		phrase string,
	) ([]client.SourceNatRule, error)
	Forwards(
		x context.Context,
		phrase string,
	) ([]client.Forward, error)
	Blocklists(
		x context.Context,
		phrase string,
	) ([]client.Blocklist, error)
	Interfaces(x context.Context) ([]client.NetworkInterface, error)
	Log(
		x context.Context,
		limit int,
	) ([]client.LogEntry, error)
	States(
		x context.Context,
		phrase string,
	) ([]client.State, error)
	AddHost(
		x context.Context,
		h client.HostRequest,
		apply bool,
	) (*client.Identifier, error)
	SetHost(
		x context.Context,
		identifier string,
		h client.HostRequest,
		apply bool,
	) error
	DeleteHost(
		x context.Context,
		identifier string,
		apply bool,
	) error
	ReconfigureDnsmasq(x context.Context) error
	Close()
}
