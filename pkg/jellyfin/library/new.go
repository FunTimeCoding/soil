package library

func New(
	identifier string,
	name string,
	collectionType string,
) *Library {
	return &Library{
		Identifier:     identifier,
		Name:           name,
		CollectionType: collectionType,
	}
}
