package impression_outbox

import "time"

func New(
	source string,
	content string,
	capturedAt time.Time,
) *Impression {
	return &Impression{Source: source, Content: content, CapturedAt: capturedAt}
}
