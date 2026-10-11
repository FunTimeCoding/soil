package key

func Find(
	v []*Key,
	name string,
) *Key {
	for _, k := range v {
		if k.Name == name {
			return k
		}
	}

	return nil
}
