package state_tree

func (t *Tree) addReference(
	owner string,
	path string,
) {
	static, pattern := referenceOf(path)

	if pattern != nil {
		t.filePattern[owner] = append(t.filePattern[owner], pattern)

		return
	}

	t.references[owner] = append(t.references[owner], static)
}
