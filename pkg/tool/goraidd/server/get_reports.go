package server

import (
	"context"
	soil "github.com/funtimecoding/soil/pkg/constant"
	timeConstant "github.com/funtimecoding/soil/pkg/time/constant"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/generated/server"
)

func (s *Server) GetReports(
	_ context.Context,
	_ server.GetReportsRequestObject,
) (server.GetReportsResponseObject, error) {
	reports, e := s.service.Reports()

	if e != nil {
		return server.GetReports500JSONResponse(
			*s.captureFail(e, soil.UnexpectedError),
		), nil
	}

	var result []server.ReportResponse

	for _, r := range reports {
		result = append(
			result,
			server.ReportResponse{
				FileName: r.Name,
				Time:     r.Modified.Format(timeConstant.DateSecond),
				Size:     r.Size,
			},
		)
	}

	return server.GetReports200JSONResponse(result), nil
}
