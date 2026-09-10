---
name: chart-sessions
description: Name and describe unnamed Claude Code sessions, one at a time with confirmation.
---

# Chart sessions

Name and describe unnamed sessions so the session list
reads like a log of what happened and when.

Read `${CLAUDE_PLUGIN_ROOT}/doc/ai/runbook/chart-sessions.md` for naming
and description conventions.

## Posture

One session at a time. Oldest unnamed first.

- Peek first to see the shape. Then use judgment -
  short sessions are fully visible from peek alone.
  Longer sessions have arcs that peek can't show.
  Read more with print or targeted greps until you
  can name what happened.
- Propose before peeking the next one. Don't stockpile
  reads.
- Don't count remaining sessions or summarize progress.
- Don't batch proposals across sessions.

## Flow

1. **Find** - `goclaude session list`. Look for sessions
   without a `*` marker. Active sessions (no date,
   0 lines) can be skipped. Work oldest first.

2. **Read** - `goclaude session peek <id>`. If the peek
   shows the full arc, propose. If not, read deeper
   with print or greps until you can.

3. **Confirm** - present the name and description. Wait
   for adjustment or approval before saving.

4. **Save** - `goclaude session edit <id> --name "..."
   --description "..."`. Both in one call. Then read
   the next one.

5. **Delete** - sessions with no conversation value
   (0 user messages, test-only). Propose deletion,
   confirm before acting.
