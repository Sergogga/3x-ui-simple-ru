# PLANS.md

## Purpose

This document defines the execution-plan format for large or risky changes.

Plans are internal agent/developer artifacts and should be written in English. Final user summaries and deployment notes for the repository owner should be written in Russian.

## When to create an ExecPlan

Create an ExecPlan for changes touching:

- Xray config generation;
- install/update scripts;
- Docker or systemd;
- database migrations;
- authentication, 2FA, TLS, or secrets;
- multi-module refactoring;
- user-facing workflow redesign;
- reliability or security-sensitive behavior.

Save active plans under `docs/exec-plans/active/`.
Move completed plans to `docs/exec-plans/completed/`.
Track deferred cleanup in `docs/exec-plans/tech-debt-tracker.md`.

## ExecPlan template

```md
# <short task name>

## Goal

Describe the outcome.

## Scope

List affected files/directories.

## User-visible behavior

Describe what the user will see. If user-facing text is included, write that text in Russian.

## Technical approach

Describe the implementation approach.

## Risks

List security, reliability, migration, and rollback risks.

## Steps

- [ ] Step 1
- [ ] Step 2
- [ ] Step 3

## Verification

List exact commands to run.

## Documentation updates

List documentation files that must be updated.

## Rollback notes

Describe how to undo the change.
```
