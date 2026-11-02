package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/constant"
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/types/report"
	"os"
	"path/filepath"
	"time"
)

func (s *Service) Generate(
	files []string,
	date time.Time,
) (*report.Report, error) {
	if len(files) == 0 {
		return nil, validation.New(constant.NoFiles)
	}

	s.logger.Structured("generate_request", "files", files)
	workdir, e := os.MkdirTemp("", "goraidparsed-")

	if e != nil {
		return nil, fmt.Errorf("create temporary directory: %w", e)
	}

	defer s.removeWorkdir(workdir)

	if f := linkFiles(files, workdir); f != nil {
		return nil, f
	}

	if f := s.parser.Parse(workdir, date); f != nil {
		return nil, fmt.Errorf("parse: %w", f)
	}

	tidCount := countTiddlers(workdir)

	if f := s.parser.Import(s.templatePath, workdir); f != nil {
		return nil, fmt.Errorf("generate tiddlers: %w", f)
	}

	source := s.findReport(workdir)

	if source == "" {
		return nil, fmt.Errorf("find report: no report in %s", workdir)
	}

	path := filepath.Join(
		s.outputPath,
		fmt.Sprintf(constant.ReportName, date.Format(constant.ReportTime)),
	)
	system.Move(source, path)

	return report.New(path, tidCount), nil
}
