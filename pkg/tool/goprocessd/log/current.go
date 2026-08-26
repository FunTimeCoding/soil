package log

func (l *Logger) Current() ([]string, int) {
	mutex.Lock()
	defer mutex.Unlock()
	result := make([]string, len(l.history)-l.generationStart)
	copy(result, l.history[l.generationStart:])

	return result, l.generationStart
}
