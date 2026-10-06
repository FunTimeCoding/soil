package open_api

type Response struct {
	Reference   string             `yaml:"$ref"`
	Description string             `yaml:"description"`
	Content     map[string]Content `yaml:"content"`
}
