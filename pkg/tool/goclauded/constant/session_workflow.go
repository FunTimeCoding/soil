package constant

const SessionWorkflowURI = "goclauded://guide/session-workflow"
const SessionWorkflow = `# Session Workflow

goclauded coordinates parallel sessions on a shared roster.
Each session is visible to the others - announces, updates,
completions, and summaries are shared.

## Lifecycle

    announce → [update]* → complete → [announce next | release]

**announce** - declare your name and topic. Binds your identity
to the MCP transport. Re-announce when scope changes; each call
replaces the previous topic. The name parameter is your assigned
pool name (shown in hook context).

**update** - record a milestone and update the roster topic
without completing. Logs to the completion table so other
sessions see progress in their hook context.

**complete** - mark the current topic done with a short summary
(commit-message density). Clears the topic. Calling complete
again on the same topic amends the existing record. If no
active topic, pass a topic parameter.

**summarize** - record a session summary covering the full arc
of the session. One summary per session; calling again amends.
Summaries are pushed to the search index for future sessions.

**moment** - capture a single line in the event stream. No
notifications, no indexing. Moments accumulate quietly.

**release** - leave the roster.

## Amending

Summarize and complete upsert by session identity - calling
them again updates the existing record and event in place.
One summary per session, one completion per topic.

The edit tool amends any event by ID. When editing a summarize
or complete event, the change cascades to the corresponding
summary or completion record automatically.

## Coordination

Use send to warn other sessions before disruptive operations
(linting, broad refactors, package renames). Broadcast to all
or address a specific session by name.

## Hook context

A UserPromptSubmit hook fires every turn and shows:
- Active sessions with topics
- Recent completions and scope changes (last 24 hours)
- Pending messages (direct or broadcast)
- Timeout notice if the session was timed out

The hook handles orientation automatically. Call history for
deeper lookback with limit, offset, since, before, and kind
filters.

## Timeouts

**Inactivity timeout (1 hour)** - session has a topic but
neither checked in nor wrote to its log. A long working turn
keeps the roster spot. Topic and timeout reason logged.

**Complete timeout (30 min)** - session completed, then went
quiet on both channels without re-announcing. Logged and shown
on next check.

## Session flows

Short: announce → work → complete → summarize → release.

Long: announce → re-announce as scope sharpens → update at
milestones → complete → announce next task → complete →
summarize the full arc → amend if needed.

Complete and summarize serve different audiences. Complete is
for the roster (immediate visibility). Summarize is for the
archive (future discoverability via search).
`
