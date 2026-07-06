package server

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/run"
	timeLibrary "github.com/funtimecoding/soil/pkg/time"
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/generated/server"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func (s *Server) PostGenerate(
	_ context.Context,
	r server.PostGenerateRequestObject,
) (server.PostGenerateResponseObject, error) {
	date := time.Now()

	if r.Body.Date != nil {
		date = *r.Body.Date
	}

	s.logger.Structured("generate_request", "files", r.Body.Files)
	workdir, e := os.MkdirTemp("", "goraidparsed-")

	if e != nil {
		return server.PostGenerate500JSONResponse(
			*s.captureFail(e, "failed to create temp directory"),
		), nil
	}

	defer func() {
		if f := os.RemoveAll(workdir); f != nil {
			s.reporter.CaptureException(f)
		}
	}()

	for _, file := range r.Body.Files {
		system.Stat(file)

		if f := os.Symlink(
			file,
			filepath.Join(workdir, filepath.Base(file)),
		); f != nil {
			return server.PostGenerate500JSONResponse(
				*s.captureFail(f, "failed to symlink file"),
			), nil
		}
	}

	dateArgument := date.Format("2006-01-02T15:04:05")
	parserScript := filepath.Join(
		s.parserPath,
		"TW5_parse_top_stats_detailed.py",
	)
	parser := run.New().NoPanic()
	parser.Directory = s.parserPath
	parser.Start("python", parserScript, workdir, "-d", dateArgument)

	if parser.Error != nil {
		return server.PostGenerate500JSONResponse{
			Error: "parser failed",
			EventIdentifier: s.reporter.CaptureWithContext(
				parser.Error,
				"parser",
				map[string]any{
					"output": parser.OutputString,
					"stderr": parser.ErrorString,
				},
			),
		}, nil
	}

	tidCount := 0
	entries := system.ReadDirectory(workdir)

	for _, entry := range entries {
		if strings.HasSuffix(entry.Name(), ".tid") {
			tidCount++
		}
	}

	tiddlerScript := filepath.Join(
		s.parserPath,
		"tiddler_import.py",
	)
	t := run.New().NoPanic()
	t.Directory = s.parserPath
	t.Start(
		"python",
		tiddlerScript,
		"--html",
		s.templatePath,
		"--logfolder",
		workdir,
		"--download",
		workdir,
	)

	if t.Error != nil {
		return server.PostGenerate500JSONResponse{
			Error: "tiddler generator failed",
			EventIdentifier: s.reporter.CaptureWithContext(
				t.Error,
				"tiddler_generator",
				map[string]any{
					"output": t.OutputString,
					"stderr": t.ErrorString,
				},
			),
		}, nil
	}

	reportName := fmt.Sprintf(
		"ONYX Log %s.html",
		date.Format(timeLibrary.DateYear),
	)
	var reportSource string

	for _, entry := range entries {
		if strings.HasSuffix(entry.Name(), constant.HypertextExtension) &&
			entry.Name() != filepath.Base(s.templatePath) {
			reportSource = filepath.Join(workdir, entry.Name())

			break
		}
	}

	// Re-scan after tiddler generation for the output HTML
	if reportSource == "" {
		updated := system.ReadDirectory(workdir)

		for _, entry := range updated {
			if strings.HasSuffix(entry.Name(), constant.HypertextExtension) &&
				entry.Name() != filepath.Base(s.templatePath) {
				reportSource = filepath.Join(workdir, entry.Name())

				break
			}
		}
	}

	reportPath := filepath.Join(s.outputPath, reportName)
	system.Move(reportSource, reportPath)

	return server.PostGenerate200JSONResponse{
		Path:     reportPath,
		TidCount: tidCount,
	}, nil
}
