package notify

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/tool/common"
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
