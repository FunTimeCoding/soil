package reading

type Reading struct {
	Server          string
	Battery         string
	Attached        bool
	DriverConnected bool
	Variables       map[string]string
}
