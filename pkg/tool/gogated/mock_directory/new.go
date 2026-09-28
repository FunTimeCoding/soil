package mock_directory

func New() *Directory {
	return &Directory{
		passwords: map[string]string{},
		members:   map[string]bool{},
	}
}
