# QUALITY_SCORE.md

## Purpose

This document defines quality gates and review expectations.

Internal quality criteria should be written in English. Final reports to the repository owner should be written in Russian.

## Baseline checks

For Go changes:

```bash
gofmt -w <changed-go-files>
go test ./...
```

For frontend changes:

```bash
cd frontend
npm ci
npm run lint
npm run build
```

For Docker or runtime changes:

```bash
docker compose config
docker build .
```

Run only the checks that are relevant and feasible in the current environment. If a check cannot be run, report why.

## Review dimensions

Score release readiness using these dimensions when useful:

- correctness;
- security;
- reliability;
- maintainability;
- user impact;
- documentation completeness;
- rollback safety.

## Documentation as quality gate

A change is incomplete if it changes architecture, frontend patterns, security behavior, reliability behavior, product workflows, or quality expectations without updating the corresponding documentation file or explaining why no documentation update is required.
