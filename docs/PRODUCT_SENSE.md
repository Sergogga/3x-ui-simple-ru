# PRODUCT_SENSE.md

## Purpose

This document captures product assumptions, user workflows, and admin experience principles.

Internal product reasoning should be written in English. User-facing product documentation, onboarding text, release notes, help text, and UI copy should be written in Russian unless explicitly requested otherwise.

## Product principles

- The panel is an administrative tool for networking and Xray-core management.
- Safety, predictability, and operational clarity are more important than visual novelty.
- User flows must make destructive or security-sensitive actions explicit.
- Defaults should avoid accidental public exposure, credential leaks, or broken Xray configuration.

## Workflows to document

Update this document or `docs/product-specs/` when changing:

- first login and onboarding;
- panel settings;
- user management;
- inbound/outbound management;
- subscription behavior;
- Telegram bot setup;
- TLS/certificate setup;
- backup/restore or migration behavior;
- installation/update experience.

## Russian user-facing documentation rule

End-user documentation should be written in Russian with direct operational steps.

Good style:

```md
1. Откройте раздел «Настройки».
2. Укажите путь к сертификату.
3. Нажмите «Сохранить».
4. Перезапустите панель, если появится соответствующее предупреждение.
```

Avoid overly abstract wording and avoid unexplained English terms where a clear Russian equivalent exists.
