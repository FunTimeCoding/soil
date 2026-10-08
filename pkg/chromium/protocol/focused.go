package protocol

func (p *Protocol) Focused() (map[string]any, error) {
	var result map[string]any
	e := p.Evaluate(
		`(function() {
                try {
                    const focused = document.activeElement;
                    if (focused) {
                        return {
                            tagName: focused.tagName,
                            id: focused.id || '',
                            className: focused.className || '',
                            selector: focused.id ? '#' + focused.id : focused.tagName.toLowerCase(),
                            type: focused.type || '',
                            contentEditable: focused.contentEditable || 'false'
                        };
                    }
                    return null;
                } catch (e) {
                    return { error: e.message };
                }
            })()`,
		&result,
	)

	if e != nil {
		return nil, e
	}

	return result, nil
}
