package face

import (
	"github.com/funtimecoding/soil/pkg/opnsense/alias"
	"github.com/funtimecoding/soil/pkg/opnsense/blocklist"
	"github.com/funtimecoding/soil/pkg/opnsense/forward"
	"github.com/funtimecoding/soil/pkg/opnsense/host"
	"github.com/funtimecoding/soil/pkg/opnsense/lease"
	"github.com/funtimecoding/soil/pkg/opnsense/log_entry"
	"github.com/funtimecoding/soil/pkg/opnsense/network_interface"
	"github.com/funtimecoding/soil/pkg/opnsense/pool"
	"github.com/funtimecoding/soil/pkg/opnsense/request"
	"github.com/funtimecoding/soil/pkg/opnsense/rule"
	"github.com/funtimecoding/soil/pkg/opnsense/source_nat"
	"github.com/funtimecoding/soil/pkg/opnsense/state"
)

type OpnsenseSource interface {
	Leases(phrase string) ([]*lease.Lease, error)
	Hosts(phrase string) ([]*host.Host, error)
	Pools(phrase string) ([]*pool.Pool, error)
	Rules(phrase string) ([]*rule.Rule, error)
	Aliases(phrase string) ([]*alias.Alias, error)
	SourceNatRules(phrase string) ([]*source_nat.Rule, error)
	Forwards(phrase string) ([]*forward.Forward, error)
	Blocklists(phrase string) ([]*blocklist.Blocklist, error)
	Interfaces() ([]*network_interface.Interface, error)
	Log(limit int) ([]*log_entry.Entry, error)
	States(phrase string) ([]*state.State, error)
	AddHost(h *request.Host) (string, error)
	SetHost(identifier string, h *request.Host) error
	DeleteHost(identifier string) error
	ReconfigureDnsmasq() error
}
