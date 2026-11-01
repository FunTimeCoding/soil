package face

import "context"

type Client interface {
	TagJournalClient
	CustomizationClient
	DeviceClient
	VirtualizationClient
	NetworkClient
	CoreClient
	RackClient
	ModuleClient
	CablingClient
	InterfaceClient
	AddressClient
	OrganizationClient
	Close()
	Forget()
	Reachable(x context.Context) error
	Address() string
}
