package janitor

import "time"

func (j *Janitor) sweep() {
	cutoff := time.Now().Add(-j.retention)
	events := j.store.MustPruneEventsBefore(cutoff)
	snapshots := j.store.MustPruneSnapshotsBefore(cutoff)

	if events > 0 || snapshots > 0 {
		j.logger.Structured("pruned", "events", events, "snapshots", snapshots)
	}
}
