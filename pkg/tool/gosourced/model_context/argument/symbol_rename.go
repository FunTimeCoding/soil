package argument

type SymbolRename struct {
	OldName  string `json:"old_name"`
	NewName  string `json:"new_name"`
	Receiver string `json:"receiver"`
}
