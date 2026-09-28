package constant

import (
	"github.com/funtimecoding/soil/pkg/identity"
	"time"
)

var Identity = identity.New(
	"goatlasd",
	"Architecture documentation by attribution",
	"goatlasd",
)

const (
	SourceKubernetes = "kubernetes"
	SourceNetbox     = "netbox"
	SourceProcess    = "goprocessd"
	SourceLease      = "gopnsensed"
	SourceOutpost    = "gooutpostd"

	OutpostHostsEnvironment = "GOOUTPOST_HOSTS"

	PlaceEnvironment = "PROCESS_PLACE"

	KindService = "service"

	ServiceLabelPrefix = "service."

	ApplicationLabel     = "app"
	ApplicationNameLabel = "app.kubernetes.io/name"

	SourceColumn          = "source"
	KindColumn            = "kind"
	ScopeColumn           = "scope"
	NameColumn            = "name"
	PackageColumn         = "package"
	VersionColumn         = "version"
	PlaceKindColumn       = "place_kind"
	PlaceIdentifierColumn = "place_identifier"
	PlaceNameColumn       = "place_name"
	SeenColumn            = "seen_at"

	HardwareAddressColumn = "hardware_address"
	AddressColumn         = "address"
	HostnameColumn        = "hostname"
	ReservedColumn        = "reserved"

	PlacementIndex = "create unique index if not exists placement_attribution on placement (source, kind, scope, name, place_kind, place_identifier)"
	SightingIndex  = "create unique index if not exists sighting_host on sighting (source, hardware_address)"

	PlacementOrder = "place_name, scope, name"
	SightingOrder  = "place_name, hostname, address"

	DeviceStatusFormat         = "device list status: %d"
	VirtualMachineStatusFormat = "virtual machine list status: %d"
	ProcessStatusFormat        = "process list status: %d"
	OutpostHostStatusFormat    = "outpost host status: %d"
	OutpostServiceStatusFormat = "outpost service list status: %d"
	OutpostTargetFailedFormat  = "outpost %s failed: %v\n"
	OutpostEveryTargetFailed   = "every outpost failed: %s"
	LeaseStatusFormat          = "lease list status: %d"
	PhysicalStatusFormat       = "physical address list status: %d"
	PlaceUnknownFormat         = "place not in the inventory: %s"

	StartMessage = "start collect worker"
	StopMessage  = "stop collect worker"

	MetricSourceLabel = "source"

	IntervalUsage = "Collection interval"

	RetentionArgument = "retention"
	RetentionUsage    = "Age at which an unseen placement is swept"

	SweepCondition = "source = ? and seen_at < ?"

	PlaceCondition     = "place_name = ?"
	PlaceKindCondition = "place_kind = ?"
	NameCondition      = "name like ?"
	PackageCondition   = "package = ?"
	NamePattern        = "%%%s%%"
	SourceCondition    = "source = ?"
	UnclaimedCondition = "place_kind = ''"

	PlaceSelect = "place_name as name, place_kind as kind, place_identifier as identifier, count(*) as count"
	PlaceGroup  = "place_name, place_kind, place_identifier"
	PlaceOrder  = "count desc, place_name"

	FixtureDeviceNode         = "golf"
	FixtureVirtualMachineNode = "echo"
	FixtureService            = "relay"
	FixtureScope              = "relay"
	FixtureApplianceService   = "lighting"
	FixturePeerService        = "mirror"
	FixturePlaceIdentifier    = 1
	FixtureHardwareAddress    = "02:aa:bb:cc:dd:07"
	FixtureAddress            = "198.51.100.7"
	FixtureHostname           = "foxtrot"

	FixtureInterval  = time.Minute
	FixtureRetention = time.Hour
)

var FixtureOutpostNames = []string{"alfa", "bravo", "charlie"}
