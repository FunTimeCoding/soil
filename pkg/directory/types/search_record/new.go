package search_record

func New(
	distinguishedName string,
	attributes map[string][]string,
) *Record {
	return &Record{DistinguishedName: distinguishedName, Attributes: attributes}
}
