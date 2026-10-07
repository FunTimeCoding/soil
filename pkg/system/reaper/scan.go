package reaper

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/system/types/zombie_detail"
	"os"
	"strconv"
	"strings"
)

func (r *Reaper) scan() map[int]zombie_detail.Detail {
	result := map[int]zombie_detail.Detail{}
	entries, e := os.ReadDir("/proc")

	if e != nil {
		return result
	}

	for _, entry := range entries {
		pid, e := strconv.Atoi(entry.Name())

		if e != nil {
			continue
		}

		raw, e := os.ReadFile(fmt.Sprintf("/proc/%d/status", pid))

		if e != nil {
			continue
		}

		content := string(raw)

		if !strings.Contains(content, "State:\tZ") {
			continue
		}

		detail := zombie_detail.Detail{}

		for _, line := range strings.Split(content, "\n") {
			if strings.HasPrefix(line, "Name:\t") {
				detail.Comm = strings.TrimPrefix(line, "Name:\t")
			}

			if strings.HasPrefix(line, "PPid:\t") {
				v, f := strconv.Atoi(strings.TrimPrefix(line, "PPid:\t"))

				if f == nil {
					detail.Ppid = v
				}
			}
		}

		result[pid] = detail
	}

	return result
}
