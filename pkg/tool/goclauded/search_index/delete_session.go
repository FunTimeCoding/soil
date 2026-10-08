package search_index

func (x *Index) DeleteSession(session string) (int, error) {
	x.mutex.Lock()
	defer x.mutex.Unlock()

	return x.remove(session)
}
