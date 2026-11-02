package face

import "context"

type Client interface {
	VirtualizationClient
	TagJournalClient
	CustomizationClient
	NetworkClient
	AddressClient
	DeviceClient
	OrganizationClient
	Address() string
	Close()
	Reachable(x context.Context) error
}
