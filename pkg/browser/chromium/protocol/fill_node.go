package protocol

func (p *Protocol) FillNode(
	backendNodeIdentifier int64,
	value string,
	direct bool,
) error {
	if direct {
		return p.fillNodeDirect(backendNodeIdentifier, value)
	}

	return p.fillNodeInsertText(backendNodeIdentifier, value)
}
