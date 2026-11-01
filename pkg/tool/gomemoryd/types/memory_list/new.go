package memory_list

func New(
	memoryType string,
	tag string,
	scope string,
	page int,
	limit int,
) *List {
	return &List{
		Type:  memoryType,
		Tag:   tag,
		Scope: scope,
		Page:  page,
		Limit: limit,
	}
}
