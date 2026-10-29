package mock_client

import "github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/client"

type Client struct {
	LeaseList     []client.Lease
	HostList      []client.Host
	PoolList      []client.Pool
	RuleList      []client.Rule
	AliasList     []client.Alias
	SourceNatList []client.SourceNatRule
	ForwardList   []client.Forward
	BlocklistList []client.Blocklist
	InterfaceList []client.NetworkInterface
	LogList       []client.LogEntry
	StateList     []client.State
	Failure       error
	Reconfigured  int
	Closed        int
}
