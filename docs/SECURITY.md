# SECURITY.md

## Purpose

This document captures security guidance for the project.

Internal security analysis should be written in English. Admin-facing setup and hardening documentation for Russian-speaking users should be written in Russian.

## Security-sensitive areas

Update this document when changing:

- authentication;
- password handling;
- 2FA / TOTP;
- session handling;
- TLS certificate paths;
- subscription URLs;
- generated client links;
- Telegram bot token and chat ID handling;
- install/update scripts;
- Docker privileges and mounts;
- systemd service permissions;
- file permissions;
- Xray config generation;
- network exposure and listen addresses.

## Rules

- Never hardcode secrets, credentials, private keys, tokens, UUIDs, IP addresses, or domains.
- Never log secrets or full subscription URLs if they contain sensitive tokens.
- Prefer least privilege in runtime scripts and Docker behavior.
- Treat public network exposure as a security decision, not a default implementation detail.
- When changing authentication or secret handling, explain the threat model and rollback plan.

## Russian admin documentation

Russian-speaking admin instructions should be concrete and operational.

Example:

```md
После установки измените стандартный логин и пароль администратора.
Не публикуйте ссылку подписки в открытых чатах: она может содержать чувствительный токен доступа.
```
