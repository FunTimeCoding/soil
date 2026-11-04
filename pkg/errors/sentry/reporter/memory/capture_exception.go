package memory

import "context"

func (m *Memory) CaptureException(e error) string {
	return m.Capture(context.Background(), e)
}
