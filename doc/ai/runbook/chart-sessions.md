# Chart sessions

Reference for the `/chart-sessions` skill - naming and
description conventions.

## Name conventions

- Kebab-case, 2-4 words
- Feature or tool name as anchor
  (`goprocessd-build`, `dns-migration-and-hetzner-server`)

## Description conventions

- One line, semicolons separating compound arcs
- What was built/fixed, scope, key details

## Tools

```
goclaude session list              # find unnamed
goclaude session list --detail     # calibrate tone
goclaude session peek <id>         # overview with samples
goclaude session print <id>        # full transcript
goclaude session edit <id> ...     # save name + description
goclaude session delete <id>       # remove empty sessions
```
