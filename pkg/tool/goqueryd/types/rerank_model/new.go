package rerank_model

func New(
	name string,
	sequenceLength int,
	maximumLength int,
) *Model {
	return &Model{
		Name:           name,
		SequenceLength: sequenceLength,
		MaximumLength:  maximumLength,
	}
}
