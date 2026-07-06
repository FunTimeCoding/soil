package site

import (
	"github.com/funtimecoding/soil/pkg/generative/anthropic/site/usage_result"
	"regexp"
	"strconv"
)

var resetPattern = regexp.MustCompile(
	`Resets?\s+(.+?)</span>`,
)

var valuePattern = regexp.MustCompile(
	`aria-valuenow="(\d+)"`,
)

func (s *Site) ReadUsage() *usage_result.Result {
	outer := s.protocol.Outer(
		"div:has(> div > div > div[role='progressbar'])",
	)

	if outer == "" {
		return nil
	}

	value := valuePattern.FindStringSubmatch(outer)

	if value == nil {
		return nil
	}

	percent, e := strconv.Atoi(value[1])

	if e != nil {
		return nil
	}

	reset := ""
	match := resetPattern.FindStringSubmatch(outer)

	if match != nil {
		reset = match[1]
	}

	return usage_result.New(
		percent,
		reset,
		0,
		"",
		0,
		0,
		"",
		"",
		"",
		0,
	)
}
