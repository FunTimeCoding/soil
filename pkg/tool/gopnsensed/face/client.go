package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/client"
)

type Client interface {
	face.Reachable
	Address() string
	ListLeases(
		x context.Context,
		phrase string,
	) ([]client.Lease, error)
	ListHosts(
		x context.Context,
		phrase string,
	) ([]client.Host, error)
	ListPools(
		x context.Context,
		phrase string,
	) ([]client.Pool, error)
	ListRules(
		x context.Context,
		phrase string,
	) ([]client.Rule, error)
	ListAliases(
		x context.Context,
		phrase string,
	) ([]client.Alias, error)
	ListSourceNat(
		x context.Context,
		phrase string,
	) ([]client.SourceNatRule, error)
	ListForwards(
		x context.Context,
		phrase string,
	) ([]client.Forward, error)
	ListBlocklists(
		x context.Context,
		phrase string,
	) ([]client.Blocklist, error)
	ListInterfaces(x context.Context) ([]client.NetworkInterface, error)
	FirewallLog(
		x context.Context,
		limit int,
	) ([]client.LogEntry, error)
	FirewallStates(
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
