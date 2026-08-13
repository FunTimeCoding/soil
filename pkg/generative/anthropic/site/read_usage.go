package site

import (
	"github.com/funtimecoding/soil/pkg/generative/anthropic/site/usage_result"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"strconv"
)

func (s *Site) ReadUsage() *usage_result.Result {
	outer := s.protocol.Outer("div:has(> div > div > div[role='meter'])")

	if outer == "" {
		return nil
	}

	value := constant.ValuePattern.FindStringSubmatch(outer)

	if value == nil {
		return nil
	}

	percent, e := strconv.Atoi(value[1])

	if e != nil {
		return nil
	}

	reset := ""
	match := constant.ResetPattern.FindStringSubmatch(outer)

	if match != nil {
		reset = match[1]
	}

	return usage_result.New(percent, reset, 0, "", 0, 0, "", "", "", 0)
}
