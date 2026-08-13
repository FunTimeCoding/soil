package enrichment

func New(
	name string,
	entity string,
	category string,
) *Enrichment {
	return &Enrichment{Name: name, Entity: entity, Category: category}
}
