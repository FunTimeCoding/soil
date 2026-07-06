package reaper

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func (r *Reaper) scan() map[int]zombieDetail {
	result := map[int]zombieDetail{}
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

		detail := zombieDetail{}

		for _, line := range strings.Split(content, "\n") {
			if strings.HasPrefix(line, "Name:\t") {
				detail.comm = strings.TrimPrefix(line, "Name:\t")
			}

			if strings.HasPrefix(line, "PPid:\t") {
				v, f := strconv.Atoi(
					strings.TrimPrefix(line, "PPid:\t"),
				)

				if f == nil {
					detail.ppid = v
				}
			}
		}

		result[pid] = detail
	}

	return result
}
