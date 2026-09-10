---
name: commit-message
description: Draft a commit message for the current changes.
---

# Commit message

Draft a commit message for the current staged and unstaged
changes.

## Style

- Lowercase, no period, no conventional commits prefix
- Terse noun phrases, not sentences
- Comma-separated topics when the commit spans multiple
  scopes ("constants, fix lint task, update runbook")
- "update" alone is valid for catch-all dependency or
  minor changes
- Don't explain why - just name what changed

## Steps

1. **Gather changes** - run `git diff` and `git status`.
   Read enough to understand the scope.

2. **Draft** - one line. Match the voice above. If
   multiple unrelated things changed, comma-separate.

3. **Present** - show the draft. Wait for adjustment or
   approval. Don't stage or commit.
