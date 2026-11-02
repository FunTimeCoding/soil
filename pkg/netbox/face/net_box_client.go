package face

import "context"

type NetBoxClient interface {
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
