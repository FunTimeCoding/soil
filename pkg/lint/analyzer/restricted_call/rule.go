package restricted_call

type Rule struct {
	Package   string
	Function  string
	AllowedIn []string
	Message   string
}
