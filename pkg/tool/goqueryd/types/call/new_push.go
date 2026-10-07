package call

func NewPush(
	collection string,
	name string,
	body string,
	metadata map[string][]string,
) *Push {
	return &Push{
		Collection: collection,
		Name:       name,
		Body:       body,
		Metadata:   metadata,
	}
}
