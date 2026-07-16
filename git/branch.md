# 브랜치 규칙

> 상태: 확정 (2026-07-11)

## 전략: Git Flow

| 브랜치 | 역할 | 어디서 → 어디로 |
| --- | --- | --- |
| `main` | 배포된 안정 버전만 | — |
| `develop` | 다음 릴리스 개발 통합 | — |
| `feature/*` | 기능 개발 | `develop` → `develop` |
| `fix/*` | develop의 버그 수정 | `develop` → `develop` |
| `release/*` | 릴리스 준비 (버전, QA) | `develop` → `main` + `develop` |
| `hotfix/*` | 배포본 긴급 수정 | `main` → `main` + `develop` |

## 네이밍: `타입/설명`

- 설명은 `kebab-case` — 경로는 kebab 원칙(파일·URL 규칙)과 동일 계열.
- 커밋 타입과 같은 어휘를 쓴다. 필요하면 `docs/`, `chore/`도 허용.

```text
feature/order-cancel
fix/token-expiry
release/1.2.0
hotfix/payment-timeout
```

## 머지: 항상 merge 커밋 (`--no-ff`)

- **모든 머지는 fast-forward 없이 merge 커밋을 남긴다.**
  - 작업(브랜치) 단위의 경계가 히스토리에 보존되어 "이 커밋들이 하나의
    작업이었다"를 나중에 알 수 있다.
  - 문제가 생기면 merge 커밋 하나를 revert해서 작업 단위로 되돌릴 수 있다.
- squash를 쓰지 않으므로 **작업 중 커밋도 커밋 메시지 규칙을 지켜서**
  남긴다. 히스토리에 그대로 남기 때문.
