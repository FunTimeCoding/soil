package request

type Get struct {
	ResourceType string
	Name         string
	Namespace    string
	Unfiltered   bool
}
