package watcher

func (w *Watcher) Index(root string) {
	w.remember(root, root)
	p, e := w.client.FindPost(root)

	if e != nil {
		w.reporter.CaptureException(e)

		return
	}

	replies, f := w.client.Thread(p)

	if f != nil {
		w.reporter.CaptureException(f)

		return
	}

	for _, v := range replies {
		w.remember(v.Identifier, root)
	}
}
