package device_change

type Change struct {
	Name           *string
	PrimaryAddress *string
	Location       *string
	Platform       *string
	Type           *string
	Role           *string
	Site           *string
	Tenant         *string
	Status         *string
	Serial         *string
	Description    *string
	CustomFields   map[string]any
	LocalContext   *map[string]any
}
