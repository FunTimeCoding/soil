package base

import (
	"fmt"
	braveConstant "github.com/funtimecoding/soil/pkg/brave/constant"
	"github.com/funtimecoding/soil/pkg/system/run"
)

func (s *Stack) Start() {
	s.T.Helper()
	s.process = run.New().Open(
		braveConstant.BravePath,
		"--headless=new",
		fmt.Sprintf("--remote-debugging-port=%d", s.Port),
		fmt.Sprintf("--user-data-dir=%s", s.profile),
		"--no-first-run",
		"--no-default-browser-check",
		"--disable-gpu",
		"about:blank",
	)
	waitReady(s.T, s.pool.Client(), s.Port)
}
