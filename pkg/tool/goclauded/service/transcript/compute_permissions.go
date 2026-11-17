package transcript

import (
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/permission/allowlist"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/permission/policy"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/permission/scratch"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/permission"
	"strings"
	"time"
)

func (s *Service) computePermissions(days int) (*permission.Report, error) {
	cut := s.clock().Add(-time.Duration(days) * 24 * time.Hour)
	rules := map[string][]*allowlist.Rule{}
	roots := map[string][]string{}
	var prompts []*permission.Prompt
	sessions := 0

	for _, e := range s.Sessions() {
		active := false
		var prefixes []string
		looked := false

		for _, c := range s.ToolCalls(e.Identifier) {
			when, f := time.Parse(time.RFC3339Nano, c.Timestamp)

			if f != nil || when.Before(cut) {
				continue
			}

			active = true

			if !strings.HasPrefix(
				c.PermissionSource,
				generativeConstant.UserDecisionPrefix,
			) || c.Name == generativeConstant.QuestionTool {
				continue
			}

			p := permission.NewPrompt(
				e.Identifier,
				when,
				c.Detail,
				c.PermissionReason,
			)
			p.Declined = c.PermissionDecision == generativeConstant.RejectDecision
			prompts = append(prompts, p)

			if c.Name != generativeConstant.BashTool {
				p.Blockers = []string{c.Name}

				continue
			}

			project, g := projectRules(rules, e.WorkDirectory)

			if g != nil {
				return nil, g
			}

			scratchRoots, g := projectScratch(roots, e.WorkDirectory)

			if g != nil {
				return nil, g
			}

			if !looked {
				looked = true
				callsign, h := s.store.SessionCallsign(e.Identifier)

				if h != nil {
					return nil, h
				}

				if callsign != "" {
					prefixes = []string{scratch.Prefix(callsign)}
				}
			}

			directory := c.Directory

			if directory == "" {
				directory = e.WorkDirectory
			}

			w := policy.New(project, scratchRoots, directory).
				WithScratchPrefixes(prefixes...)
			p.Cleared = w.Allows(c.Detail)
			p.Blockers = w.Blockers()
		}

		if active {
			sessions++
		}
	}

	return permission.Compute(prompts, sessions, days), nil
}
