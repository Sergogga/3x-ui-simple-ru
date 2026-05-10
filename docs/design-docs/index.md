# Design Docs Index

## Purpose

This file is the table of contents for design documents.

Write internal design docs in English. If a design doc contains user-facing instructions or examples, keep those parts in Russian.

## Documents

- `core-beliefs.md` — non-negotiable project principles.

## Add a design doc when

- a change affects multiple modules;
- there are meaningful alternatives;
- security, reliability, migration, or rollback implications need to be recorded;
- the change will be hard to reconstruct from code alone.

## Naming convention

Use lowercase kebab-case:

```text
YYYY-MM-DD-short-topic.md
```

Example:

```text
2026-05-10-xray-restart-lifecycle.md
```
