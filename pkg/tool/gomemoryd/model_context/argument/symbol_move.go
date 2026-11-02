package argument

type SymbolMove struct {
	Directory    string `json:"directory"`
	Name         string `json:"name"`
	NewDirectory string `json:"new_directory"`
	NewName      string `json:"new_name"`
}
