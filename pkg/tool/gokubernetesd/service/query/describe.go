package query

type Describe struct {
	ResourceType string
	Name         string
	Namespace    string
	Unfiltered   bool
}
