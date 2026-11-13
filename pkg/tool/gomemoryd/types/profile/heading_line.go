package profile

import (
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"strconv"
	"strings"
)

func headingLine(line string) *Heading {
	rest, found := strings.CutPrefix(line, constant.MemoryHeadingPrefix)

	if !found {
		return nil
	}

	if i := strings.LastIndex(rest, constant.RankSeparator); i >= 0 {
		_, e := strconv.ParseFloat(rest[i+len(constant.RankSeparator):], 64)

		if e == nil {
			rest = rest[:i]
		}
	}

	if strings.HasSuffix(rest, constant.TagClose) {
		if i := strings.LastIndex(rest, constant.TagOpen); i >= 0 {
			rest = rest[:i]
		}
	}

	rest, found = strings.CutSuffix(rest, constant.IdentifierClose)

	if !found {
		return nil
	}

	open := strings.LastIndex(rest, constant.IdentifierOpen)

	if open < 0 {
		return nil
	}

	identifier, e := strconv.ParseInt(
		rest[open+len(constant.IdentifierOpen):],
		10,
		64,
	)

	if e != nil {
		return nil
	}

	return NewHeading(identifier, rest[:open])
}
