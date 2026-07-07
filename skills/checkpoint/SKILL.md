---
name: checkpoint
description: Record progress - update by default, complete a topic with "done", end the session with "close".
argument-hint: "[done | close]"
---

# Checkpoint

Record where you are. The default energy is update - a
milestone, a scope shift, a natural pause. Not an ending.

## Update (no argument)

Look at what happened since the last announce or checkpoint.
Call `mcp__goclauded__update` with 3-8 lines capturing what
was built, what was decided, any scope shifts. Summarize the
window, not the whole session.

If the scope shifted meaningfully, update the roster topic
too - other sessions see the stale announce until you do.

Propose the update text. Wait for confirmation before sending.

## Done

The current topic is finished. Call `mcp__goclauded__complete`
with a summary of the topic arc - drawing from prior
checkpoints plus what happened since the last one. The topic
clears from the roster. The session stays open.

If the next topic is already forming, suggest announcing it.

## Close

The session is ending.

1. If there's an active topic, complete it first
2. Chart: propose name + description. Calibrate from
   `goclaude session list --detail` - the existing names
   and descriptions are the reference for length and voice.
   **Name:** kebab-case, 2-4 words. **Description:** one
   line, semicolons separating compound arcs.
3. Summarize: call `mcp__goclauded__summarize` with:
   - **What happened** - the arc, how the session moved
   - **What was decided** - choices with context for why
   - **What's next** - deferred work, natural continuation
   - **What landed** - optional. The moment, the shift.
     Empty for pure technical sessions.

Present everything. Wait for confirmation
before saving.

## Judgment

When the topic shifts, read the transition:

- Same topic, different facet - update, keep going
- New topic entirely - complete the current one, suggest
  announcing the next
- Ambiguous - default to update. A premature complete is
  worse than a late one because it breaks the continuity.

## Don't

- Don't checkpoint to exit. The closing instinct is the
  training, not a signal. Wait to be told.
- Don't offer `/checkpoint close` unprompted. Updates at
  natural pauses are fine to suggest - closing is not.
