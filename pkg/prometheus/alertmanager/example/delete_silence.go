package example

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/tool/common"
)

func DeleteSilence() {
	a := argument.NewSimple("delete-silence")
	a.ParseSimple()
	c := common.Alertmanager()
	defer c.Close()
	c.MustDeleteSilence(a.RequiredPositional(0, "SILENCE_IDENTIFIER"))
}
