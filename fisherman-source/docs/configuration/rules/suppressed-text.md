---
id: rules
title: Rules
---

## suppressed-text

The rule that prohibits committing if forbidden lines were added.

``` yaml
- type: suppressed-text
  when: 1 == 1
  substrings: [ 'suppressed', 'text' ]
  exclude: [ 'some/excluded/file.go' ]
```

- **substrings** - List of lines that should not be included in the commit.
- **exclude** - List of globs in which you do not need to check this rule.
