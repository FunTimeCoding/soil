package request

type Events struct {
	Namespace    string
	Kind         string
	Name         string
	Type         string
	Limit        int
	IncludeMuted bool
}
