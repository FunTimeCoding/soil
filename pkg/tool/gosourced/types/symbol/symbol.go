package symbol

type Symbol struct {
	Package   string
	Path      string
	Clause    string
	Receiver  string
	Name      string
	Kind      string
	Variant   string
	File      string
	Line      int
	Start     int
	End       int
	Text      string
	Signature string
	Body      string
	Shape     string
}
