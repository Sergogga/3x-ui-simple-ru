# FRONTEND.md

## Purpose

This document is the source of truth for frontend-specific agent and developer guidance. Do not create a separate `frontend/AGENTS.md` unless explicitly requested.

Frontend-facing internal guidance should be written in English. UI copy, examples of Russian labels, end-user help text, and Russian-language user documentation should stay in Russian.

## Stack

The frontend is located in `frontend/` and uses:

- Vue 3
- Vite
- Ant Design Vue 4
- vue-i18n
- ESLint
- npm with `package-lock.json`

## Commands

Run frontend commands from `frontend/`:

```bash
npm ci
npm run lint
npm run build
```

Use `npm`, not yarn or pnpm.

## UI and i18n rules

- Keep UI changes small and consistent with existing Ant Design Vue patterns.
- Do not hardcode user-facing text if the surrounding feature uses translations.
- When adding or changing labels, buttons, messages, or validation text, update the relevant translation files.
- Preserve existing RTL/LTR and multilingual assumptions.
- Avoid replacing existing date/time libraries unless explicitly requested.
- Keep Russian user-facing copy natural and concise.

Examples of acceptable Russian UI copy:

- `Сохранить`
- `Отменить`
- `Добавить inbound`
- `Перезапустить Xray`
- `Проверить настройки`
- `Не удалось загрузить данные`

## API rules

- Reuse existing API/client patterns.
- Do not change backend endpoint paths unless the backend change is part of the same task.
- Handle loading, error, and empty states for new UI flows.
- Do not expose secrets, tokens, UUIDs, or private config values in the browser UI unless already intentionally displayed by the current product behavior.
- Be careful with subscription links, user credentials, Telegram settings, TLS paths, and Xray-generated values.

## Component and state rules

Document concrete project patterns here as the frontend evolves:

- component placement;
- composables/helpers;
- API client conventions;
- form validation patterns;
- table/filter patterns;
- modal/confirmation patterns;
- notification/error handling patterns.

## Documentation rules

Update this file when changing frontend architecture, build flow, API usage, i18n, routing, or component conventions.

Also update:

- `docs/DESIGN.md` for UI patterns, layout, navigation, form behavior, table behavior, or visual language changes;
- `docs/product-specs/` for user-facing workflow changes;
- `docs/QUALITY_SCORE.md` if lint/build/test expectations or review gates change;
- `ARCHITECTURE.md` if frontend build output, asset serving, or frontend/backend boundaries change.

If frontend changes do not require documentation updates, explain why in the final response.

## Verification

After frontend changes, run:

```bash
npm run lint
npm run build
```

If either command fails, report the exact error and do not claim the change is fully verified.
