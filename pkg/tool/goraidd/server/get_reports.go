package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/time"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/generated/server"
	"strings"
)

func (s *Server) GetReports(
	_ context.Context,
	_ server.GetReportsRequestObject,
) (server.GetReportsResponseObject, error) {
	var result []server.ReportResponse

	for _, e := range system.ReadDirectory(s.outputPath) {
		if e.IsDir() || !strings.HasSuffix(
			e.Name(),
			constant.HypertextExtension,
		) {
			continue
		}

		i, f := e.Info()

		if f != nil {
			return server.GetReports500JSONResponse(
				*s.captureFail(f, constant.UnexpectedError),
			), nil
		}

		result = append(
			result,
			server.ReportResponse{
				FileName: e.Name(),
				Time:     i.ModTime().Format(time.DateSecond),
				Size:     i.Size(),
			},
		)
	}

	return server.GetReports200JSONResponse(result), nil
}
