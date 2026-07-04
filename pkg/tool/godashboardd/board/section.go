package board

type Section struct {
	Name    string   `yaml:"name"`
	Entries []*Entry `yaml:"entries"`
}
