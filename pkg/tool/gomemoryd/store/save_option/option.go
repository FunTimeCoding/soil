package save_option

type Option struct {
	Name             string
	Content          string
	Description      string
	Type             string
	Scope            string
	Tags             []string
	Metadata         map[string]string
	Source           string
	ParentIdentifier *int64
	ProvenanceFile   string
	ProvenanceAnchor string
	ProvenanceHash   string
	Ordinal          int
}
