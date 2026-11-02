package log

import "bytes"

func (l *Logger) Write(p []byte) (int, error) {
	l.bufferMutex.Lock()
	defer l.bufferMutex.Unlock()
	buffer := bytes.NewBuffer(p)

	for {
		line, e := buffer.ReadBytes('\n')

		if len(line) > 0 {
			if line[len(line)-1] == '\n' {
				if len(line) != 1 || len(l.buffers) > 0 {
					l.writeBuffers(line)
				}

				l.stopFlush()
			} else {
				l.buffers = append(l.buffers, bytes.Clone(line))
				l.scheduleFlush()
			}
		}

		if e != nil {
			break
		}
	}

	return len(p), nil
}
