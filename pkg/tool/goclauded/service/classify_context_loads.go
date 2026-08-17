package service

import (
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/tool_call"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/context_load"
	"strings"
)

func classifyContextLoads(
	sessionIdentifier string,
	calls []tool_call.Call,
) []context_load.Load {
	var result []context_load.Load
	modePending := false

	for _, c := range calls {
		if c.Name == constant.SkillTool {
			modePending = inputString(c.Input, constant.Skill) ==
				constant.ModeSkill

			continue
		}

		if c.Name == constant.ListMemoriesTool {
			tag := inputString(c.Input, constant.Tag)

			if modePending && tag != "" {
				result = append(result, *modeLoad(sessionIdentifier, &c, tag))
				modePending = false
			}

			continue
		}

		if c.Name == constant.ProfileTool {
			result = append(
				result,
				profileLoads(sessionIdentifier, &c)...,
			)

			continue
		}

		if c.Name == constant.SearchMemoryTool {
			result = append(
				result,
				searchLoads(sessionIdentifier, &c)...,
			)

			continue
		}

		if strings.HasPrefix(c.Name, constant.GetMemoryTool) {
			result = append(
				result,
				memoryLoads(sessionIdentifier, &c)...,
			)
		}
	}

	return result
}
