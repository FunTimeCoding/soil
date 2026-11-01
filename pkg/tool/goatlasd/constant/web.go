package constant

const (
	PlacesPath     = "/places"
	PlacementsPath = "/placements"
	SightingsPath  = "/sightings"

	DashboardTitle  = "Atlas"
	PlacementsTitle = "Placements"
	SightingsTitle  = "Sightings"

	SearchParameter    = "name"
	UnclaimedParameter = "unclaimed"
	UnclaimedValue     = "1"

	DeviceSegment  = "device"
	MachineSegment = "machine"

	KindParameter = "kind"
	NameParameter = "name"

	DeviceLabel  = "device"
	MachineLabel = "virtual machine"

	PlaceWord      = "place"
	PlacesWord     = "places"
	PlacementWord  = "placement"
	PlacementsWord = "placements"
	SightingWord   = "sighting"
	SightingsWord  = "sightings"

	UnclaimedSummary = "%d unclaimed"
	CardKindFormat   = "%s · %s"

	CarriesHeading   = "Runs here"
	NetworkHeading   = "Seen on the network"
	UnclaimedHeading = "Unclaimed"
	NothingCarried   = "Nothing attests this machine."
	NothingSeen      = "No lease carries this machine."
	PlaceUnknown     = "No such place in the atlas."
	NothingUnclaimed = "Every host on the network is in the inventory."
	NothingMatched   = "No placement carries that name."

	AllSightingsLink = "all sightings"
	UnclaimedOnly    = "unclaimed only"
	EveryHost        = "every host"
	SearchLabel      = "Name"
	SearchAction     = "Search"

	SourceHeader   = "Source"
	ScopeHeader    = "Scope"
	NameHeader     = "Name"
	SeenHeader     = "Last seen"
	AddressHeader  = "Address"
	HardwareHeader = "Hardware address"
	HostnameHeader = "Hostname"
	PlaceHeader    = "Place"

	PlaceCountClass = "place-count"

	InlineStyle = `
.place-count { font-size: 1.6rem; font-weight: 600; margin: 0; }
`
)
