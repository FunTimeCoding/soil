package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/console/status/tag"
	"github.com/funtimecoding/soil/pkg/tool/common"
)

func Silence() {
	f := option.Color.Copy().Tag(tag.State)

	for _, a := range common.Alertmanager().MustSilences(true) {
		fmt.Println(a.Format(f))
	}
}
