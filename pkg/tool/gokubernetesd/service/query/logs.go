package query

type Logs struct {
	Name       string
	Namespace  string
	Container  string
	Tail       int
	Since      string
	Previous   bool
	Timestamps bool
}
