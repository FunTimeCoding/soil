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

	SearchFormClass = "search-form"
	CardGridClass   = "card-grid"
	PlaceCardClass  = "place-card"
	PlaceCountClass = "place-count"
	PlaceKindClass  = "place-kind"

	InlineStyle = `
.card-grid {
	display: grid;
	grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
	gap: 1rem;
	margin-top: 1rem;
}
.place-card {
	border: 1px solid var(--pico-muted-border-color);
	border-radius: var(--pico-border-radius);
	padding: 1.25rem;
	transition: border-color 0.2s;
	background: var(--pico-card-background-color);
}
.place-card:hover { border-color: var(--pico-primary); }
.place-card a { text-decoration: none; color: inherit; }
.place-card h4 { margin-bottom: 0.3rem; }
.place-count { font-size: 1.6rem; font-weight: 600; }
.place-kind { color: var(--pico-muted-color); font-size: 0.8rem; }
.search-form {
	display: flex;
	gap: 0.5rem;
	align-items: center;
	max-width: 32rem;
}
.search-form input { margin-bottom: 0; }
.search-form button { margin-bottom: 0; width: auto; }
`
)
