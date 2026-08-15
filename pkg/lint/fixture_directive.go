package lint

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/file_report"
	"io"
	"strings"
)

func FixtureDirective(
	path string,
	r io.Reader,
) *file_report.Report {
	s := file_report.New(path, r)
	var pending bool
	var pendingLine string
	var pendingNumber int

	for s.Scan() {
		line, number := s.Text()

		if pending {
			if !strings.HasPrefix(line, "const ") && line != "const (" {
				s.AddConcern(
					constant.FixtureKey,
					constant.FixtureMisplacedText,
					path,
					pendingNumber,
					pendingLine,
					false,
				)
			}

			pending = false
		}

		if rule, okay := parseFixtureRule(line); okay {
			if !testdataPath(path) {
				s.AddConcern(
					constant.FixtureKey,
					constant.FixtureOutsideTestdataText,
					path,
					number,
					line,
					false,
				)
			} else if rule != constant.StrayConstantRule {
				s.AddConcern(
					constant.FixtureKey,
					constant.FixtureUnknownRuleText,
					path,
					number,
					line,
					false,
				)
			} else {
				pending = true
				pendingLine = line
				pendingNumber = number
			}
		}

		s.PassLine(line)
	}

	if pending {
		s.AddConcern(
			constant.FixtureKey,
			constant.FixtureMisplacedText,
			path,
			pendingNumber,
			pendingLine,
			false,
		)
	}

	return s.Finalize()
}
