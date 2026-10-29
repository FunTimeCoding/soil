package service

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/service/report"
	"strings"
)

func (s *Service) Reports() ([]*report.Report, error) {
	var result []*report.Report

	for _, n := range system.ReadDirectory(s.outputPath) {
		if n.IsDir() || !strings.HasSuffix(n.Name(), soil.HypertextExtension) {
			continue
		}

		i, e := n.Info()

		if e != nil {
			return nil, e
		}

		result = append(result, report.New(n.Name(), i.Size(), i.ModTime()))
	}

	return result, nil
}
