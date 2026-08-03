package alert

import (
	"github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/strings/join/key_value"
)

func (a *Alert) formatName(f *option.Format) string {
	result := a.Name

	if f.UseColor {
		result = constant.Yellow("%s", result)
	}

	if f.HasTag(constant.TagEmoji) {
		if v := a.emoji(); len(v) > 0 {
			result = key_value.Space(join.Empty(v...), result)
		}
	}

	return result
}
