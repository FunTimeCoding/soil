package stored_document

func New(
	sessionIdentifier string,
	path string,
	body string,
	metadata map[string][]string,
) *Document {
	return &Document{
		SessionIdentifier: sessionIdentifier,
		Path:              path,
		Body:              body,
		Metadata:          metadata,
	}
}
