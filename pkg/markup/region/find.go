package region

func Find(
	content string,
	name string,
) (string, error) {
	inner, closing, e := locate(content, name)

	if e != nil {
		return "", e
	}

	return content[inner:closing], nil
}
