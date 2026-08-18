package concern

type Concern struct {
	Key      string
	Text     string
	Path     string
	Type     string
	Line     int
	LineText string
	Fixed    bool
	Planned  bool
}
