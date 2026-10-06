package layout

type Section struct {
	Name    string   `yaml:"name"`
	Entries []*Entry `yaml:"entries"`
}
