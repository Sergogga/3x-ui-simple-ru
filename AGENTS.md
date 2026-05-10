# AGENTS.md

## Project context

This repository is a customized Russian-oriented fork of `MHSanaei/3x-ui`.

It is a web control panel for Xray-core with:

- Go backend and CLI entrypoint in `main.go`
- backend packages in `config/`, `database/`, `logger/`, `sub/`, `util/`, `web/`, `xray/`
- Vue 3 frontend in `frontend/`
- built frontend assets served from `web/dist`
- translations in `web/translation`
- install/update/runtime scripts: `install.sh`, `update.sh`, `x-ui.sh`, `DockerInit.sh`, `DockerEntrypoint.sh`
- deployment files: `Dockerfile`, `docker-compose.yml`, `x-ui.service.*`, `x-ui.rc`

Treat this as a security-sensitive networking/admin panel. Small changes can affect authentication, TLS, Xray config generation, user limits, traffic accounting, subscriptions, Telegram bot behavior, fail2ban, Docker, and systemd installs.

## Language policy

- Write agent-facing instructions, internal implementation notes, architecture notes, execution plans, quality gates, and technical design notes in English.
- Write user interaction, final summaries, deployment notes for the repository owner, and end-user-facing documentation in Russian unless the user explicitly asks for another language.
- UI copy examples, help text, user-facing product documentation, release notes, admin instructions, and troubleshooting content intended for Russian-speaking users must stay in Russian.
- Do not translate existing Russian UI strings or user-facing documentation to English unless the task explicitly requires it.
- When a file mixes internal instructions with user-facing content, keep the internal instructions in English and keep the user-facing text/examples in Russian.

## Default working rules

- Answer and summarize changes in Russian unless the user asks otherwise.
- Before editing, identify whether the change touches backend, frontend, installer/runtime scripts, Docker, translations, documentation, or Xray config generation.
- Prefer small, reviewable diffs. Do not rewrite unrelated code.
- Keep upstream compatibility with `MHSanaei/3x-ui` where possible.
- Do not change public CLI flags, service names, Docker volume paths, database paths, config file paths, or public API routes unless the task explicitly requires it.
- Do not remove existing translations or language keys. When adding UI text, update relevant translation files.
- Do not introduce new production dependencies without explaining why they are necessary.
- Never hardcode credentials, tokens, UUIDs, private keys, domains, IPs, Telegram tokens, or certificate paths.
- Never print secrets in logs.
- Be careful with commands that may modify the host system. Do not run install scripts as root unless explicitly requested.
- Do not add nested `AGENTS.md` files unless explicitly requested. Keep frontend-specific guidance in `docs/FRONTEND.md`.

## Required documentation structure

Maintain this documentation layout as part of development work:

```text
AGENTS.md
ARCHITECTURE.md
docs/
├── design-docs/
│   ├── index.md
│   ├── core-beliefs.md
│   └── ...
├── exec-plans/
│   ├── active/
│   ├── completed/
│   └── tech-debt-tracker.md
├── generated/
│   └── db-schema.md
├── product-specs/
│   ├── index.md
│   ├── new-user-onboarding.md
│   └── ...
├── references/
│   ├── design-system-reference-llms.txt
│   ├── nixpacks-llms.txt
│   ├── uv-llms.txt
│   └── ...
├── DESIGN.md
├── FRONTEND.md
├── PLANS.md
├── PRODUCT_SENSE.md
├── QUALITY_SCORE.md
├── RELIABILITY.md
└── SECURITY.md
```

Documentation is part of the definition of done:

- Update `ARCHITECTURE.md` when changing module boundaries, runtime topology, storage, Xray integration, deployment model, or high-level data flow.
- Update `docs/DESIGN.md` when changing UI patterns, navigation, forms, tables, layout, visual language, or design decisions.
- Update `docs/FRONTEND.md` when changing frontend architecture, build flow, state/API patterns, i18n, routing, component conventions, or frontend quality gates.
- Update `docs/SECURITY.md` when changing auth, 2FA, TLS, secrets, permissions, Telegram bot token handling, subscription URLs, install scripts, Docker privileges, or network exposure.
- Update `docs/RELIABILITY.md` when changing startup/shutdown behavior, migrations, service restart logic, Xray restart logic, cron/background jobs, fail2ban, observability, backup/restore, or error handling.
- Update `docs/QUALITY_SCORE.md` when adding or changing quality gates, test expectations, review criteria, lint/build requirements, or release readiness scoring.
- Update `docs/PRODUCT_SENSE.md` and `docs/product-specs/` when changing user-facing workflows, onboarding, settings, subscription behavior, admin UX, or product assumptions.
- Update `docs/design-docs/` for architectural decisions and larger technical designs. Keep `docs/design-docs/index.md` as the table of contents.
- Keep `docs/design-docs/core-beliefs.md` aligned with the project's non-negotiable principles.
- Use `docs/exec-plans/active/` for active execution plans and move completed plans to `docs/exec-plans/completed/` when done.
- Keep `docs/exec-plans/tech-debt-tracker.md` updated when deferring cleanup, accepting tradeoffs, or discovering risks.
- Treat `docs/generated/` as generated or mechanically maintained documentation. Do not hand-edit generated content if a generator exists; update the generator instead.
- Keep `docs/generated/db-schema.md` in sync with database models and migrations.
- Store external LLM-oriented references in `docs/references/` and avoid mixing them into design docs.
- Do not create new top-level documentation files unless the task explicitly requires it. Prefer extending the structure above.
- If a code change does not require documentation updates, explicitly say why in the final response.

## Execution plans

For large changes, especially changes touching Xray config generation, install/update scripts, database migrations, auth, 2FA, Docker, systemd, or multi-module refactoring, create an ExecPlan before implementation.

Use `docs/PLANS.md` as the plan format. Save active plans under `docs/exec-plans/active/`.

An ExecPlan must include:

- goal
- affected files/directories
- user-visible behavior
- technical approach
- risks and rollback notes
- implementation steps
- verification commands
- documentation updates required

Keep the plan updated while working. When the task is complete, move the plan to `docs/exec-plans/completed/` and add any remaining follow-up work to `docs/exec-plans/tech-debt-tracker.md`.

## Backend / Go rules

- Use idiomatic Go.
- Keep package boundaries stable.
- Prefer explicit error handling over silent fallback.
- Do not ignore errors from database, filesystem, network, Xray, TLS, or process-management operations.
- When touching Go files, run:

```bash
gofmt -w <changed-go-files>
go test ./...
```

- If `go test ./...` is not feasible because of environment limitations, explain exactly what failed and which narrower check was run instead.
- For changes touching startup, CLI, settings, DB migration, or signal handling, inspect `main.go` and the relevant service code before editing.
- For changes touching Xray inbound/outbound generation, validate assumptions against the current Xray-related code before changing behavior.

## Frontend rules

- The frontend lives in `frontend/`.
- Use npm, not yarn or pnpm, because `frontend/package-lock.json` is present and Docker uses `npm ci`.
- Keep detailed frontend guidance in `docs/FRONTEND.md` instead of a nested `frontend/AGENTS.md`.
- When touching frontend files, run from `frontend/`:

```bash
npm ci
npm run lint
npm run build
```

- Keep Vue components compatible with Vue 3, Ant Design Vue 4, and Vite.
- Do not bypass existing i18n patterns. Add or update translation keys instead of hardcoding Russian text in components unless the surrounding code already does that intentionally.
- Avoid large UI rewrites unless requested.
- Update `docs/FRONTEND.md` and, when relevant, `docs/DESIGN.md` for frontend architecture or UI behavior changes.

## Docker / install / runtime rules

- Docker build flow is multi-stage: frontend build first, then Go builder, then Alpine runtime image.
- `docker-compose.yml` uses host networking and mounts:
  - `$PWD/db/` to `/etc/x-ui/`
  - `$PWD/cert/` to `/root/cert/`
- Be very conservative with:
  - `install.sh`
  - `update.sh`
  - `x-ui.sh`
  - `DockerInit.sh`
  - `DockerEntrypoint.sh`
  - `x-ui.service.*`
- If editing shell scripts:
  - preserve POSIX/bash compatibility expected by the existing script
  - quote variables
  - avoid destructive commands without checks
  - do not change service names or install paths unless explicitly requested
- Update `ARCHITECTURE.md`, `docs/RELIABILITY.md`, or `docs/SECURITY.md` when runtime behavior changes.

## Security-sensitive areas

Before changing any of these, pause and reason about regressions:

- authentication and password handling
- 2FA / TOTP
- Telegram bot token and chat ID handling
- TLS certificate paths
- database migrations
- subscription URLs and generated links
- Xray config generation
- traffic/IP/user limits
- file permissions
- systemd service behavior
- Docker runtime privileges and mounted paths

For security fixes, prefer minimal targeted changes and include a short explanation of the threat model.

## Verification checklist

Before final response, report:

- files changed
- documentation files updated, or why no documentation update was required
- commands run
- whether `gofmt`, `go test ./...`, `npm run lint`, `npm run build`, or Docker checks passed
- any commands that could not be run and why
- remaining risks or follow-up checks

## PR / final answer style

Final responses to the repository owner must be in Russian. Use this structure when applicable:

1. Что изменено
2. Какие документы обновлены
3. Как проверено
4. Что не удалось проверить
5. На что обратить внимание при деплое

Keep the summary concise but include exact command results when relevant.
