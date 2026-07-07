---
name: wrap-up
description: Name and describe the current session by reviewing what happened, calibrating from existing sessions
---

# Session Wrap-up

Name and describe your own session before it ends.

## Steps

1. Run `goclaude session list` to see existing named sessions.
2. Run `goclaude session show` on a few named sessions to read their descriptions. Use these as calibration for the style, density, and format of both names and descriptions.
3. Investigate your own session — use `goclaude session peek` on your session ID, then `goclaude session print` with selective line ranges and `goclaude session tools` to understand the full arc of what happened.
4. Propose a name and description. Do not save yet.
   - **Name:** lowercase, dash-separated, 2-5 words. Captures the essence, not an exhaustive list. Avoid `-and-` connectors.
   - **Description:** 1-3 sentences at commit-message density. What was done, what was produced, what's the state at the end. Include both old and new names if things were renamed during the session.
5. Wait for adjustment or approval.
6. Once approved, run `goclaude session edit --name <name> --description <description>` or use the `edit_session` MCP tool to save both.
