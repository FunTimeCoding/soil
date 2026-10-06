package constant

const (
	FixtureTableRow      = "| 25–44 | 23.2±1.5 | 3.0±0.5 | 1.1±0.4 |"
	FixtureTableDocument = `# Results

### Table 1

Table 1. Alfa bravo charlie.

| a | b |
|---|---|
| 1 a | 1 b |
| 2 a | 2 b |
| 3 a | 3 b |
| 4 a | 4 b |
| 5 a | 5 b |
| 6 a | 6 b |

Closing paragraph.
`
	FixtureSectionDocument = `---
title: Alfa
---

# Alfa
Intro line one.
Intro line two.

## Bravo

- first item
- second item

~~~
# not a heading
~~~

## Charlie

| a | b |
|---|---|
| 1 | 2 |
`
	FixtureCompareBefore       = "# A\n\nalfa bravo charlie.\n\ndelta echo.\n"
	FixtureCompareRestructured = "# A\n\nalfa bravo charlie.\n\n## B\n\ndelta echo.\n"
	FixtureCompareShaved       = "# A\n\nalfa charlie.\n\ndelta echo.\n"
	FixtureWindowLine          = "alfa bravo charlie delta echo\n"
	FixtureWindowHeading       = "### Late\n"
)
