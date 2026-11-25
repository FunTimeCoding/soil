package watched_request

func New(
	instance string,
	project int64,
	request int64,
) *Key {
	return &Key{Instance: instance, Project: project, Request: request}
}
