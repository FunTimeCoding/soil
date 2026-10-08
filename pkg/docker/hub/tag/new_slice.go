package tag

func NewSlice(v []Response) []*Tag {
	var result []*Tag

	for i := range v {
		result = append(result, New(&v[i]))
	}

	return result
}
