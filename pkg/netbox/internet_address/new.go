package internet_address

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/netbox-community/go-netbox/v4"
	"net"
)

func New(a *netbox.IPAddress) *Address {
	address, network, e := net.ParseCIDR(a.Display)
	errors.FatalOnError(e)
	var objectType string

	if t := a.AssignedObjectType.Get(); t != nil {
		objectType = *t

		if objectType != "" {
			validateObjectType(objectType)
		} else {
			errors.Warning(
				"empty object type for internet address %s\n",
				a.Display,
			)
		}
	}

	var object int64

	if a.AssignedObjectId.IsSet() {
		object = a.GetAssignedObjectId()
	}

	face, device, machine := holder(a.AssignedObject)
	var role string

	if r := a.GetRole(); r.Value != nil {
		role = string(*r.Value)
	}

	var status string

	if s := a.GetStatus(); s.Value != nil {
		status = string(*s.Value)
	}

	return &Address{
		Identifier:       a.GetId(),
		Name:             a.GetDisplay(),
		Address:          address,
		Network:          network,
		ObjectType:       objectType,
		ObjectIdentifier: object,
		Interface:        face,
		Device:           device,
		VirtualMachine:   machine,
		Role:             role,
		Status:           status,
		DomainName:       a.GetDnsName(),
		Description:      a.GetDescription(),
		Raw:              a,
	}
}
