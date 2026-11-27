package base

import (
	"github.com/funtimecoding/soil/pkg/browser/constant"
	"github.com/funtimecoding/soil/pkg/system/writer"
	"strings"
)

func heavyPage() string {
	var b strings.Builder
	b.WriteString(constant.ChromiumFixtureHeavyPagePrefix)

	for i := 0; i < constant.ChromiumFixtureHeavyRowCount; i++ {
		writer.Print(&b, constant.ChromiumFixtureHeavyRowFormat, i)
	}

	b.WriteString(constant.ChromiumFixtureHeavyPageSuffix)

	return b.String()
}
