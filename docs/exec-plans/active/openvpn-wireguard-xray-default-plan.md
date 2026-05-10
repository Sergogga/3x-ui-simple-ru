# OpenVPN support, WireGuard multi-client inbound, and default Xray version bump

## Goal

Implement three coordinated product/runtime improvements:
1. Add OpenVPN connection support in 3x-ui-simple-ru.
2. Add support for creating multiple WireGuard clients for a single inbound (instead of only one generated client config).
3. Set default installed Xray version to `26.1.23` during installation.

## Scope

Expected affected areas (to be confirmed during implementation):

- Backend:
  - `web/` (handlers/APIs for connection creation and config generation)
  - `xray/` (inbound/client generation logic for WireGuard)
  - `config/`, `util/` (defaults/version plumbing if needed)
  - `database/` (if schema changes are required for storing multiple WireGuard clients)
- Frontend:
  - `frontend/` (UI for OpenVPN creation flows and multiple WireGuard clients)
  - `web/translation/` (new i18n keys)
- Runtime/install:
  - `install.sh` (default Xray version `26.1.23`)
  - potentially `update.sh` / `x-ui.sh` if version defaults are duplicated
- Documentation:
  - `ARCHITECTURE.md`
  - `docs/FRONTEND.md`
  - `docs/PRODUCT_SENSE.md`
  - `docs/product-specs/`
  - `docs/RELIABILITY.md`
  - `docs/SECURITY.md`

## User-visible behavior

- Admin can create and manage OpenVPN client connections from panel workflows consistent with existing connection types.
- For a WireGuard inbound, admin can add multiple client entries and generate/download per-client configurations.
- Fresh install path defaults Xray runtime version to `26.1.23` unless explicitly overridden by existing mechanisms.

## Technical approach

1. **OpenVPN support design (reference-based):**
   - Review `https://github.com/nyr/openvpn-install` as behavior reference for OpenVPN config generation patterns, certificate/profile packaging assumptions, and operational defaults.
   - Define adaptation boundaries for 3x-ui architecture (do not directly copy installer behavior that conflicts with current service/runtime model).
   - Introduce backend models/services for OpenVPN client records and config artifact generation.

2. **WireGuard multi-client support:**
   - Analyze current single-client creation flow for WireGuard inbound.
   - Extend storage + API contracts to support N clients per WireGuard inbound with stable identifiers.
   - Implement per-client config generation and lifecycle actions (create/revoke/delete/export) with compatibility for existing single-client entries.

3. **Default Xray version update:**
   - Locate all installation defaults/version constants and set default to `26.1.23`.
   - Ensure update/upgrade paths do not unintentionally downgrade/overwrite custom versions.

4. **Compatibility and migration:**
   - If DB schema changes are needed, add migration with backward-compatible behavior.
   - Ensure existing inbounds/users continue to work without requiring immediate manual conversion.

## Risks

- **Security risk:** OpenVPN key/cert material handling may leak secrets if storage/logging is not carefully controlled.
- **Reliability risk:** Multi-client WireGuard changes can break existing single-client assumptions in API/UI.
- **Operational risk:** Xray default version bump may conflict with some environments if binary/source changes in upstream packaging.
- **Rollback risk:** Partial rollout (backend without frontend or migration mismatch) can lead to broken admin workflows.

## Steps

- [x] Create detailed OpenVPN domain model + API contract draft and map against existing panel architecture.
- [x] Audit current WireGuard inbound data flow (DB model, API handlers, UI state, config generation).
- [ ] Implement backend support for OpenVPN entities and config generation.
- [ ] Implement backend support for multiple WireGuard clients per inbound.
- [ ] Add/adjust DB migrations for new entities/relations (if required).
- [ ] Extend frontend forms/tables/actions for OpenVPN and WireGuard multi-client management.
- [ ] Add/update translation keys in `web/translation`.
- [ ] Update install defaults to Xray `26.1.23` and verify non-destructive upgrade behavior.
- [ ] Add tests for new backend logic and critical API flows.
- [ ] Run lint/build/test verification gates and perform manual smoke checks.
- [ ] Update required architecture/product/security/reliability/frontend docs.

## Progress updates

### 2026-05-10

Completed discovery + design preparation for the first implementation wave.

1. **OpenVPN model and API draft prepared**
   - Proposed DB entity: `openvpn_clients` (id, inbound_id nullable for future server binding, remark, username/cn, status, cert_ref, key_ref, profile_path, created_at, updated_at, revoked_at).
   - Proposed secret handling: certificate/private key material is never stored in plaintext logs; only file refs/metadata in DB.
   - Proposed API surface under existing inbound/client management style:
     - `GET /panel/api/inbounds/openvpn/clients`
     - `POST /panel/api/inbounds/openvpn/clients`
     - `POST /panel/api/inbounds/openvpn/clients/{id}/revoke`
     - `DELETE /panel/api/inbounds/openvpn/clients/{id}`
     - `GET /panel/api/inbounds/openvpn/clients/{id}/export`
   - Compatibility note: keep route/auth/session conventions aligned with existing panel API patterns; no change to public service names, install paths, or DB root path.

2. **WireGuard flow audit completed**
   - Frontend currently already renders multiple generated WireGuard configs/links in:
     - `frontend/src/pages/inbounds/InboundInfoModal.vue`
     - `frontend/src/pages/inbounds/QrCodeModal.vue`
   - Frontend model generation methods are centralized in:
     - `frontend/src/models/inbound.js` (`genWireguardConfigs`, `genWireguardLinks` call sites confirmed).
   - Next backend tasks: confirm whether DB/API still assumes single-client peer storage and implement persistent per-peer lifecycle (create/revoke/delete/export), not only derived multi-output rendering.

3. **Default Xray version rollout started**
   - Backend version-list endpoint now prefers `v26.1.23` as the first item when this version exists in the fetched GitHub release set.
   - This preserves compatibility with current API contracts and lets UI flows that preselect the first version align with the required default baseline.
   - Added deterministic unit tests with a mocked HTTP server for the Xray release-list ordering behavior (preferred-present and preferred-missing scenarios).

## Verification

Planned verification commands (to execute during implementation phase):

```bash
go test ./...

cd frontend
npm ci
npm run lint
npm run build
```

Additional targeted checks to define during implementation:

- API integration checks for OpenVPN create/export/revoke flows.
- API integration checks for WireGuard multi-client CRUD and config rendering.
- Install-script smoke check validating default Xray version path (`26.1.23`).

## Documentation updates

Required by completion:

- `ARCHITECTURE.md` (new OpenVPN + WireGuard multi-client architecture/data flow)
- `docs/FRONTEND.md` (UI architecture changes for new workflows)
- `docs/PRODUCT_SENSE.md` and related `docs/product-specs/*` (admin workflows and assumptions)
- `docs/SECURITY.md` (key/cert handling, secret storage, access controls)
- `docs/RELIABILITY.md` (runtime behavior and failure handling)
- `docs/QUALITY_SCORE.md` (if quality gates are changed)
- `docs/design-docs/index.md` + new design doc(s) if implementation complexity warrants ADR-style tracking
- `docs/generated/db-schema.md` if schema is changed

## Rollback notes

- Revert feature flags/entry points for OpenVPN and WireGuard multi-client endpoints/UI if critical regressions appear.
- Roll back DB migration(s) only with explicit downgrade scripts and backup validation.
- Restore previous default Xray version in install scripts if compatibility incidents are detected.
- Keep pre-change backup of `/etc/x-ui/` DB and generated configs before deployment rollout.
