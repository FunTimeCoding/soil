package notify

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/common"
)

func Notify() {
	c := common.Alertmanager()
	s := &State{}
	stop := make(chan struct{})
	go worker(stop, c, s)
	system.KillSignalBlock()
	close(stop)
	fmt.Println("Stop")
}
