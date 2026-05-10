# ARCHITECTURE.md

## Purpose

This document describes the high-level architecture of `3x-ui-simple-ru`.

Keep internal architecture notes in English. If a section is intended as deployment or administration documentation for Russian-speaking users, write that section in Russian.

Update this document when a change affects:

- backend/frontend/runtime boundaries;
- panel startup and service lifecycle;
- Xray-core integration;
- data storage and migrations;
- Docker, systemd, install, or update flow;
- network model, TLS, subscriptions, Telegram bot, or fail2ban behavior.

## Component map

Maintain this section as the architecture evolves.

### Backend

Describe the Go packages, service boundaries, CLI entrypoint, and responsibilities.

### Frontend

Describe the Vue/Vite application structure, build output, and how frontend assets are served.

### Database

Describe the database location, models, migrations, backup implications, and schema ownership.

### Xray integration

Describe how the panel generates, validates, applies, reloads, and restarts Xray configuration.

### Runtime scripts

Describe installer, updater, shell helper, Docker entrypoint, and platform-specific scripts.

### Docker and systemd

Describe supported deployment topologies, mounted paths, service names, restart behavior, and operational constraints.

## Data flows

Document the main flows:

1. User login and authenticated panel access.
2. Panel and user settings management.
3. Inbound/outbound Xray configuration creation and updates.
4. Subscription link generation and consumption.
5. Panel and Xray restart/reload lifecycle.
6. Telegram bot, fail2ban, cron jobs, and other background tasks.

## Runtime topology

Document the supported ways to run the project:

- bare metal / systemd;
- Docker / Docker Compose;
- migration and update flow;
- backup and restore considerations;
- network exposure and TLS termination.

## Architecture decisions

Large decisions must be captured in `docs/design-docs/` and linked from `docs/design-docs/index.md`.

Each design document should include:

- context;
- decision;
- alternatives considered;
- security implications;
- reliability implications;
- migration and rollback notes.
