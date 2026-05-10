# RELIABILITY.md

## Purpose

This document captures operational reliability guidance for the panel.

Internal reliability notes should be written in English. Operator-facing troubleshooting steps for Russian-speaking users should be written in Russian.

## Reliability-sensitive areas

Update this document when changing:

- startup and shutdown behavior;
- process signal handling;
- panel restart logic;
- Xray restart/reload behavior;
- database migrations;
- cron/background jobs;
- Telegram bot runtime behavior;
- fail2ban behavior;
- Docker entrypoint behavior;
- systemd unit behavior;
- backup/restore assumptions;
- logging and observability.

## Operational rules

- Prefer explicit errors over silent fallback.
- Avoid restart loops without clear logs.
- Do not make migrations destructive unless the task explicitly requires it and rollback is documented.
- Preserve existing service names and paths unless the task explicitly requires a change.
- Treat failed Xray reload/restart as a high-impact event.

## Russian troubleshooting content

When writing operator-facing troubleshooting, use Russian and include exact commands where possible.

Example:

````md
Проверить статус сервиса:

```bash
systemctl status x-ui
```

Посмотреть последние логи:

```bash
journalctl -u x-ui -n 100 --no-pager
```
````
