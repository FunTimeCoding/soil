package lifecycle

import "github.com/funtimecoding/soil/pkg/system"

func (l *Lifecycle) RunUntilSignal() {
	l.Run()
	system.KillSignalBlock()
	l.Stop()
}
