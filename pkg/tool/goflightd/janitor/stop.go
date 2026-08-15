package janitor

func (j *Janitor) Stop() {
	close(j.stop)
}
