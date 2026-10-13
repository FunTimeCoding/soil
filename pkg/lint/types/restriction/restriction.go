package restriction

type Restriction struct {
	Package   string
	Symbol    string
	AllowedIn []string
	Message   string
}
