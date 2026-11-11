package stored_document

type Document struct {
	SessionIdentifier string
	Path              string
	Body              string
	Metadata          map[string][]string
}
