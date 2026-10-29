package project_match

func New(
	key string,
	name string,
	kind string,
) *Match {
	return &Match{Key: key, Name: name, Type: kind}
}
