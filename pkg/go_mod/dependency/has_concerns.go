package dependency

func (d *Dependency) HasConcerns() bool {
	return len(d.concern) > 0
}
