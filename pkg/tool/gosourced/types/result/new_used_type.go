package result

func NewUsedType(
	typeName string,
	methods []string,
) *UsedType {
	return &UsedType{Type: typeName, Methods: methods}
}
