package search_index

func (x *Index) Append(
	session string,
	path string,
) {
	x.mutex.Lock()
	defer x.mutex.Unlock()
	x.appendFile(session, path)
}
