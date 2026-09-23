package protocol

import "github.com/funtimecoding/soil/pkg/errors"

func (p *Protocol) Focused() map[string]any {
	var result map[string]any
	errors.PanicOnError(
		p.Evaluate(
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
		),
	)

	return result
}
