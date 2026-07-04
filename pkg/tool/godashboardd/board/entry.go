package board

type Entry struct {
	Label  string `yaml:"label"`
	Link   string `yaml:"link"`
	Icon   string `yaml:"icon"`
	Widget string `yaml:"widget"`
	Rows   []*Row `yaml:"rows"`
}
