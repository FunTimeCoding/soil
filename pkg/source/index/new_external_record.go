package index

func NewExternalRecord(imports []string) *ExternalRecord {
	return &ExternalRecord{Imports: imports}
}
