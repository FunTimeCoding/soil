package machine_change

type Change struct {
	Name           *string
	PrimaryAddress *string
	Platform       *string
	Tenant         *string
	Cores          *float64
	Memory         *int32
	Status         *string
}
