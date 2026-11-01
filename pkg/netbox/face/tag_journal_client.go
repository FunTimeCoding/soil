package face

import (
	"github.com/funtimecoding/soil/pkg/netbox/device"
	"github.com/funtimecoding/soil/pkg/netbox/journal_entry"
	"github.com/funtimecoding/soil/pkg/netbox/tag"
	"github.com/funtimecoding/soil/pkg/netbox/virtual_machine"
)

type TagJournalClient interface {
	AddDeviceJournalEntry(
		device string,
		kind string,
		comments string,
	) (*journal_entry.Entry, error)
	AddTag(
		deviceName string,
		tag string,
	) (*device.Device, error)
	AddVirtualJournalEntry(
		machine string,
		kind string,
		comments string,
	) (*journal_entry.Entry, error)
	AddVirtualTag(
		name string,
		tag string,
	) (*virtual_machine.Machine, error)
	CreateTag(name string) (*tag.Tag, error)
	DeleteJournalEntry(identifier int32) error
	DeviceJournalEntries(
		device string,
		limit int32,
		offset int32,
	) ([]*journal_entry.Entry, error)
	DeviceTagNames(name string) ([]string, error)
	RemoveTag(
		deviceName string,
		tag string,
	) (*device.Device, error)
	RemoveVirtualTag(
		name string,
		tag string,
	) (*virtual_machine.Machine, error)
	Tags() ([]*tag.Tag, error)
	UpdateJournalEntry(
		identifier int32,
		kind string,
		comments string,
	) (*journal_entry.Entry, error)
	VirtualJournalEntries(
		machine string,
		limit int32,
		offset int32,
	) ([]*journal_entry.Entry, error)
	DevicesByTag(s string) ([]*device.Device, error)
	MustAddTag(
		deviceName string,
		tag string,
	) *device.Device
	MustDeviceTagNames(name string) []string
	MustDevicesByTag(s string) []*device.Device
	MustRemoveTag(
		deviceName string,
		tag string,
	) *device.Device
	MustTags() []*tag.Tag
	TagByName(n string) (*tag.Tag, error)
}
