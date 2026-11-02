package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

type TagJournalClient interface {
	AddDeviceJournalEntry(
		x context.Context,
		device string,
		kind string,
		comments string,
	) (client.JournalEntry, error)
	AddDeviceTag(
		x context.Context,
		device string,
		tag string,
	) (client.Device, error)
	CreateTag(
		x context.Context,
		name string,
	) (client.Tag, error)
	DeleteJournalEntry(
		x context.Context,
		identifier int32,
	) error
	ListDeviceJournalEntries(
		x context.Context,
		device string,
		limit int32,
		offset int32,
	) ([]*client.JournalEntry, error)
	ListDeviceLabels(
		x context.Context,
		device string,
	) ([]*client.Label, error)
	ListDeviceTags(
		x context.Context,
		device string,
	) ([]string, error)
	ListTags(
		x context.Context,
	) ([]*client.Tag, error)
	RemoveDeviceLabel(
		x context.Context,
		device string,
		key string,
	) error
	RemoveDeviceTag(
		x context.Context,
		device string,
		tag string,
	) (client.Device, error)
	SetDeviceLabel(
		x context.Context,
		device string,
		key string,
		value string,
	) (client.Label, error)
	UpdateJournalEntry(
		x context.Context,
		identifier int32,
		kind string,
		comments string,
	) (client.JournalEntry, error)
}
