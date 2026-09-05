package browser_tester

func (b *Browser) SnapshotText() string {
	b.T.Helper()

	return formatNodes(b.Snapshot(), 0)
}
