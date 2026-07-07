---
name: lint
description: Run the lint pipeline and fix issues with judgment. Don't sweep errors under the rug.
---

# Lint

Read `${CLAUDE_PLUGIN_ROOT}/doc/ai/runbook/lint.md` for invocation, the tool
pipeline, spec pointers, and architecture docs. What
follows here is how to hold the fixes.

## Posture

Separate obvious fixes from judgment calls. Obvious
fixes (unused imports, spacing, naming drift) - just
fix them. Judgment calls - surface them. When the
convention doesn't cover the case, or the fix isn't
clear, present the finding and the options before
changing anything.

## Unresolved findings

Findings the linter can't resolve are decisions, not
facts. "Cannot de-alias", "cannot fix", informational
reports that survive a lint pass - each one is a
question, not a conclusion. Present them one at a time.
Don't batch them, don't declare them permanent, don't
categorize them as acceptable without input.

## What not to do

**Don't suppress errors with `_`.** Only byte counts
get the blank identifier, never error returns. If the
error genuinely can't happen, explain why.

**Don't add lint-ignore annotations.** No `//nolint`,
no `//lint:ignore`. If the linter is wrong, the fix is
in the code or in the analyzer - not in a comment.

**Don't invent conventions.** When a lint finding
reveals a gap - a pattern the conventions don't
address - flag it. That's a decision, not a fix.
