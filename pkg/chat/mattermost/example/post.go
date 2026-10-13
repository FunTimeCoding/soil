package example

import "github.com/funtimecoding/soil/pkg/tool/common"

func Post() {
	m := common.Mattermost()
	defer m.Close()
	m.MustPostDefault("Hello friend")
}
