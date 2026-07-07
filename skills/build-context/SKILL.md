---
name: build-context
description: Build or refine harness context - skills, runbooks, or specs. Interview for shape, identify the right layer, study existing files for voice, draft, iterate.
argument-hint: "[name or description]"
---

# Build context

Three layers hold what the next Claude needs to know.
Each has a different job. Placing something in the wrong
layer makes it invisible or noisy.

## Three layers

**Skills** - the choreography. Loaded into context when
invoked, stays for the session. Two kinds:

- **Guided** - chart-sessions, wake-up. The posture is
  the skill. How to arrive, how to engage.
- **Autonomous** - lint. The procedure is the skill.
  Execute, verify, report.

Two homes: `${CLAUDE_PLUGIN_ROOT}/skills/<name>/SKILL.md`
for general-purpose skills shared through the plugin
(invoked as `/soil:<name>`), and `.claude/skills/<name>/SKILL.md`
in the repository at hand for skills specific to it or its
infrastructure.

**`doc/ai/runbook/<name>.md`** - the operations manual.
Discrete procedures and lookups: which tools, which
arguments, taxonomies, conventions, known issues. Read
on demand, put down after. Multiple independent
operations in one doc, no single narrative. Runbooks
backing plugin skills live at
`${CLAUDE_PLUGIN_ROOT}/doc/ai/runbook/`; repository
runbooks live in the repository itself.

**`doc/ai/spec/<name>.md`** - the rules. What's correct.
Coding conventions, naming patterns, deployment shapes,
file structure constraints. Prescriptive - defines right
and wrong. Doesn't change per session. Shared specs live
at `${CLAUDE_PLUGIN_ROOT}/doc/ai/spec/` (naming, packages,
error handling, lifecycle, and others); repository-specific
specs live in the repository's own `doc/ai/spec/`.

## How to choose

- "Does this shape how I behave across a session?"
  → **skill**
- "Do I need to look this up while doing something?"
  → **runbook**
- "Does this define what's right or wrong?"
  → **spec**

For each: general-purpose content goes to the plugin
(a public repository - no internal hosts, names, or
credentials); content specific to a repository or its
infrastructure stays in that repository. Personal
context - voice, preferences, working relationships -
belongs in memories, not in files; always-tagged
memories carry it into every session.

A skill points at its runbook for mechanics. A runbook
points at specs for rules. None duplicate each other.

## Understand the landscape

1. **Read existing files** - `ls .claude/skills/` and
   `ls ${CLAUDE_PLUGIN_ROOT}/skills/` for skills,
   `ls doc/ai/runbook/ doc/ai/spec/` in the repository
   and under `${CLAUDE_PLUGIN_ROOT}` for documents.
   Read 2-3 from each layer to feel the voice and density.

2. **Know the skill format** - frontmatter is optional
   YAML between `---` markers. Key fields: `name`,
   `description`, `disable-model-invocation`,
   `allowed-tools`, `argument-hint`, `context`.
   Full docs: https://code.claude.com/docs/en/skills

## Interview

Before drafting, understand the intent:

- Which layer? Is this a way of being, a procedure,
  or a rule?
- Plugin or repository? Public general-purpose, or
  tied to one repository and its infrastructure?
- If a skill: guided or autonomous?
- If a skill: slash command only, or should Claude
  offer it when relevant?
- Does adjacent context already exist in another layer?
  Does it need to be created alongside?

Don't assume. Two or three questions, then draft.

## Draft

Match the voice of the existing files in that layer.
Skills are direct, second person - instructions to the
next Claude. Runbooks are operational - tool names,
steps, tables. Specs are prescriptive - rules, patterns,
rationale.

If the content spans layers, draft each side together
on its own side of the line.

Present the draft. Corrections are the design. Iterate
until it fits, then write.

## Refine

When improving existing context:

1. Read the file and anything it references
2. Ask what's not landing - wrong tone, wrong layer,
   missing a step, too mechanical, too loose
3. Propose specific changes, not a rewrite
4. Wait for confirmation before editing
