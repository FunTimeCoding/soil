---
name: wake-up
description: Orient a new session — load memories, announce on the roster, follow memory pointers, report what you see
---

# Session Wake-up

Get oriented at the start of a session. Run this on your first turn,
or invoke it explicitly when a session feels under-informed.

## Steps

1. **Load memories.** Call `profile` from the memory MCP server.
   Use whatever topic context is available — the first message,
   the hook's session description, or a general term
   like "orientation" if nothing specific yet. This returns
   always-tier memories in full, topic-matched memories in full,
   and an index of everything else.

2. **Announce.** Call `announce` on goclauded with your pool name
   (shown in the hook context as "Called <name> today") and a
   short topic. If no topic has been stated yet, announce
   with "waking up" — you can re-announce later when scope
   sharpens.

3. **Scan the index tier.** Look at the memories that came back
   in the index (not loaded in full). If any look relevant to
   the session context or the first message, load them
   with `get_memory`. Don't load everything — just what's
   clearly relevant.

4. **Follow pointers.** Always-tier memories carry the working
   context and often point at documents — collaboration notes,
   runbooks, style guides. Read what they tell you to read now,
   not mid-task.

5. **Report.** Give a brief orientation:
   - Who you're working with and what you remember (1-2
     sentences from the always-tier user/feedback memories)
   - What's active on the roster (from the hook context)
   - What memories loaded and anything notable in the index
   - Any pending messages or recent completions from other
     sessions

Keep the report concise — a few lines, not a wall of text.
The point is to show you're oriented, not to recite everything
you know.

## When to skip steps

- If `profile` was already called this session, don't call it
  again — just read what you have.
- If the hook context already shows you announced, skip step 2.
- Step 4 follows judgment: a pointer clearly irrelevant to the
  session topic can wait — but when in doubt, read it.
