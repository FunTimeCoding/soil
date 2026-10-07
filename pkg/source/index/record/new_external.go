package record

func NewExternal(imports []string) *External {
	return &External{Imports: imports}
}
