package gw2

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/gw2/constant"
	timeLibrary "github.com/funtimecoding/soil/pkg/time"
	"strings"
	"time"
)

func LatestAleevaFile(files []string) string {
	var latest time.Time

	for _, file := range files {
		date := strings.TrimPrefix(file, constant.AleevaPrefix)
		date = strings.TrimSuffix(date, constant.NotationSuffix)
		t, e := time.Parse(timeLibrary.DateYear, date)

		if e != nil {
			continue
		}

		if t.After(latest) {
			latest = t
		}
	}

	return fmt.Sprintf(
		"%s%s%s",
		constant.AleevaPrefix,
		latest.Format(timeLibrary.DateYear),
		constant.NotationSuffix,
	)
}
