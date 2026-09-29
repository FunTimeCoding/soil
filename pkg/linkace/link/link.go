package link

type Link struct {
	Identifier      int
	Title           string
	Link            string
	Description     string
	Host            string
	Status          int
	Visibility      int
	ListIdentifiers []int
	TagIdentifiers  []int
	Lists           []Relation
	Tags            []Relation
}
