package example

import "github.com/funtimecoding/soil/pkg/tool/common"

func Post() {
	m := common.Mattermost()
	m.MustPostDefault("Hello friend")
}
