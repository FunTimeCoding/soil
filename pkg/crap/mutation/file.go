package mutation

type File struct {
	Name    string    `json:"file_name"`
	Mutants []*Mutant `json:"mutations"`
}
