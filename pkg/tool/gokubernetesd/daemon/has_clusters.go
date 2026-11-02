package daemon

func (d *Daemon) HasClusters() bool {
	return len(d.clusters) > 0
}
