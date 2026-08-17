# 브랜치 규칙

> 상태: 확정 (2026-07-26)

## 전략: Git Flow

| 브랜치 | 역할 | 어디서 → 어디로 |
| --- | --- | --- |
| `main` | 배포된 안정 버전만 | — |
| `develop` | 다음 릴리스 개발 통합 | — |
| `feat/*` | 기능 개발 | `develop` → `develop` |
| `fix/*` | develop의 버그 수정 | `develop` → `develop` |
| `refactor/*` | 동작 변화 없는 구조 개선 | `develop` → `develop` |
| `perf/*` | 성능 개선 | `develop` → `develop` |
| `test/*` | 테스트 추가·수정 | `develop` → `develop` |
| `docs/*` | 문서 변경 | `develop` → `develop` |
| `style/*` | 포맷팅 등 의미 없는 변경 | `develop` → `develop` |
| `chore/*` | 빌드·의존성·설정 등 잡무 | `develop` → `develop` |
| `release/*` | 릴리스 준비 (버전, QA) | `develop` → `main` + `develop` |
| `hotfix/*` | 배포본 긴급 수정 | `main` → `main` + `develop` |

## 네이밍: `타입/설명`

- 작업 브랜치 타입은 커밋 타입과 같은 어휘를 쓴다.
- `release/*`, `hotfix/*`는 배포 흐름을 나타내는 관리 브랜치다.
- 설명은 `kebab-case` — 경로는 kebab 원칙(파일·URL 규칙)과 동일 계열.

```text
feat/order-cancel
fix/token-expiry
docs/install-guide
release/1.2.0
hotfix/payment-timeout
```

## 작업 브랜치 동기화

- 개인 작업 브랜치는 최신 기준 브랜치 위로 rebase한다. 불필요한 중간
  merge commit 없이 변경 흐름을 선형으로 유지하기 위함이다.
- 다른 사람과 함께 사용하는 브랜치는 rebase하지 않는다. 커밋 ID가
  바뀌면 다른 사람의 작업 기준이 사라지기 때문이다. 기준 브랜치의 변경이
  필요하면 merge로 가져온다.
- `main`, `develop`, `release/*`, `hotfix/*`의 히스토리는 재작성하지 않는다.
- force push는 본인만 사용하는 작업 브랜치에서만 허용하며, 원격 변경을
  확인하는 `--force-with-lease`를 사용한다. `--force`는 사용하지 않는다.

## 보호 브랜치 머지: merge commit (`--no-ff`)

- `main` 또는 `develop`으로 PR을 병합할 때는 fast-forward 없이 merge
  commit을 남긴다.
  - 작업(브랜치) 단위의 경계가 히스토리에 보존되어 "이 커밋들이 하나의
    작업이었다"를 나중에 알 수 있다.
  - 문제가 생기면 merge 커밋 하나를 revert해서 작업 단위로 되돌릴 수 있다.
- squash merge와 rebase merge는 사용하지 않는다.
- 본인만 사용하는 작업 브랜치에서는 PR을 병합하기 전에 interactive
  rebase로 `WIP`, `fixup!` 커밋을 squash할 수 있다. 서로 다른 논리적
  변경까지 하나로 합치지는 않는다.
- squash를 쓰지 않으므로 **작업 중 커밋도 커밋 메시지 규칙을 지켜서**
  남긴다. 히스토리에 그대로 남기 때문.
