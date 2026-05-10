# DESIGN.md

## Purpose

This document captures UI and interaction design guidance for the project.

Internal design guidance should be written in English. User-facing UI copy, help text, labels, release notes, screenshots captions, and end-user documentation examples should be written in Russian unless explicitly requested otherwise.

## Design principles

- Keep admin workflows predictable and conservative.
- Prefer clarity over visual complexity.
- Avoid hiding security-sensitive state.
- Do not introduce UI changes that make it easier to misconfigure Xray, TLS, users, limits, or subscriptions.
- Preserve existing multilingual and RTL/LTR assumptions.

## UI patterns to document

Update this file when changing:

- navigation;
- forms;
- validation behavior;
- tables and filters;
- modals and confirmations;
- empty, loading, and error states;
- destructive action flows;
- design tokens, colors, icons, spacing, or layout conventions.

## Russian user-facing copy

When adding or changing Russian UI text, keep it short and operational.

Preferred tone examples:

- `Сохранить изменения`
- `Перезапустить Xray`
- `Проверить подключение`
- `Ошибка сохранения настроек`
- `Изменения применены`

Avoid vague text such as:

- `ОК`
- `Готово`, when the action result is ambiguous
- `Что-то пошло не так`, unless a specific error cannot be shown safely

## Documentation updates

For major UI or workflow changes, also update:

- `docs/FRONTEND.md` for implementation patterns;
- `docs/PRODUCT_SENSE.md` for user workflow assumptions;
- relevant files under `docs/product-specs/`.
