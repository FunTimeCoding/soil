package face

import (
	"github.com/funtimecoding/soil/pkg/netbox/internet_address"
	"github.com/funtimecoding/soil/pkg/netbox/internet_address_range"
	"github.com/funtimecoding/soil/pkg/netbox/prefix"
	"github.com/funtimecoding/soil/pkg/netbox/reservation"
	"github.com/funtimecoding/soil/pkg/netbox/service"
	"github.com/funtimecoding/soil/pkg/netbox/service_template"
	"github.com/funtimecoding/soil/pkg/netbox/site"
	"github.com/netbox-community/go-netbox/v4"
	"net"
)

type AddressClient interface {
	CreateAddress(
		interfaceIdentifier int32,
		address string,
		status string,
	) (*internet_address.Address, error)
	CreateInternetAddressRange(
		start string,
		end string,
		status string,
		description string,
	) (*internet_address_range.Range, error)
	CreatePrefix(
		cidr string,
		s *site.Site,
		description string,
	) (*prefix.Prefix, error)
	DeleteInternet(identifier int32) error
	DeviceAddresses(device string) ([]*internet_address.Address, error)
	InternetAddressRanges() ([]*internet_address_range.Range, error)
	Prefixes() ([]*prefix.Prefix, error)
	CreateInternet(
		objectType string,
		objectIdentifier int64,
		i net.IP,
		m net.IPMask,
	) (*netbox.IPAddress, error)
	FindUnusedAddress(sub *net.IPNet) (net.IP, error)
	InternetAddresses() ([]*internet_address.Address, error)
	MustCreateInternet(
		objectType string,
		objectIdentifier int64,
		i net.IP,
		m net.IPMask,
	) *netbox.IPAddress
	MustDeviceAddresses(device string) []*internet_address.Address
	MustFindUnusedAddress(sub *net.IPNet) net.IP
	MustInternetAddressRanges() []*internet_address_range.Range
	MustInternetAddresses() []*internet_address.Address
	MustPrefixByDescription(d string) *prefix.Prefix
	MustPrefixByName(n string) *prefix.Prefix
	MustPrefixes() []*prefix.Prefix
	MustServiceTemplates() []*service_template.Template
	MustServices() []*service.Service
	PrefixByDescription(d string) (*prefix.Prefix, error)
	PrefixByName(n string) (*prefix.Prefix, error)
	Reservations() ([]*reservation.Reservation, error)
	ServiceTemplates() ([]*service_template.Template, error)
	Services() ([]*service.Service, error)
}
