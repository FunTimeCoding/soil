package service

func OwnerOf(
	owner map[string]string,
	path string,
) string {
	if v, okay := owner[path]; okay {
		return v
	}

	return owner[aliasPath(path)]
}
