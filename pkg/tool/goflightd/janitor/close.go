package janitor

func (j *Janitor) Close() {
	close(j.stop)
}
