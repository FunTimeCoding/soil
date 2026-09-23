package base

import (
	"github.com/funtimecoding/soil/pkg/chromium/constant"
	"github.com/funtimecoding/soil/pkg/system/writer"
	"strings"
)

func heavyPage() string {
	var b strings.Builder
	b.WriteString(constant.FixtureHeavyPagePrefix)

	for i := 0; i < constant.FixtureHeavyRowCount; i++ {
		writer.Print(&b, constant.FixtureHeavyRowFormat, i)
	}

	b.WriteString(constant.FixtureHeavyPageSuffix)

	return b.String()
}
