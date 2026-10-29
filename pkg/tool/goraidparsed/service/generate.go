package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/system"
	timeConstant "github.com/funtimecoding/soil/pkg/time/constant"
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/types/report"
	"os"
	"path/filepath"
	"time"
)

func (s *Service) Generate(
	files []string,
	date time.Time,
) (*report.Report, error) {
	s.logger.Structured("generate_request", "files", files)
	workdir, e := os.MkdirTemp("", "goraidparsed-")

	if e != nil {
		return nil, fmt.Errorf("create temporary directory: %w", e)
	}

	defer s.removeWorkdir(workdir)

	if f := linkFiles(files, workdir); f != nil {
		return nil, f
	}

	if f := s.python(
		filepath.Join(s.parserPath, "TW5_parse_top_stats_detailed.py"),
		workdir,
		"-d",
		date.Format("2006-01-02T15:04:05"),
	); f != nil {
		return nil, fmt.Errorf("parse: %w", f)
	}

	tidCount := countTiddlers(workdir)

	if f := s.python(
		filepath.Join(s.parserPath, "tiddler_import.py"),
		"--html",
		s.templatePath,
		"--logfolder",
		workdir,
		"--download",
		workdir,
	); f != nil {
		return nil, fmt.Errorf("generate tiddlers: %w", f)
	}

	source := s.findReport(workdir)

	if source == "" {
		return nil, fmt.Errorf("find report: no report in %s", workdir)
	}

	path := filepath.Join(
		s.outputPath,
		fmt.Sprintf("ONYX Log %s.html", date.Format(timeConstant.DateYear)),
	)
	system.Move(source, path)

	return report.New(path, tidCount), nil
}
