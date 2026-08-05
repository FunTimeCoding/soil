package restriction

type Restriction struct {
	Package   string
	Function  string
	AllowedIn []string
	Message   string
}
