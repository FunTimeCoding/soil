package watcher

func (w *Watcher) author(identifier string) string {
	result, e := w.client.User(identifier)

	if e != nil {
		return identifier
	}

	return result.Username
}
