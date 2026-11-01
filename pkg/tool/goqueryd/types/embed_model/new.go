package embed_model

func New(
	name string,
	documentPrompt string,
	queryPrompt string,
) *Model {
	return &Model{
		Name:           name,
		DocumentPrompt: documentPrompt,
		QueryPrompt:    queryPrompt,
	}
}
