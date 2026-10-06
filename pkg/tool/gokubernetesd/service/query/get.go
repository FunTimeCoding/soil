package query

type Get struct {
	ResourceType string
	Name         string
	Namespace    string
	Unfiltered   bool
}
