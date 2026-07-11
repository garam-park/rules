# 작명 규칙 (Naming)

> 상태: 초안 — `🔲 결정 필요` 항목은 확정되지 않은 규칙입니다.

## 기본 원칙

규칙은 두 층위로 나눈다. 층위를 섞으면 (예: 변수 snake_case vs 인스턴스 camelCase)
규칙끼리 충돌하므로, 각 항목이 어느 층위인지 먼저 구분한다.

1. **언어 무관 층** — 프로젝트 전체에서 하나로 고정한다.
   (DB, API JSON, 파일/디렉터리, 환경변수)
2. **언어별 층** — 해당 언어 생태계의 관례를 그대로 따른다.
   (변수, 함수, 클래스, 상수 등. 린터/생태계와 싸우지 않기 위함)

## 언어 무관 규칙

| 항목 | 표기 | 이유 |
|---|---|---|
| DB 테이블 | `snake_case` 복수형 (`users`) | 테이블은 행(row)의 집합이므로 복수형 |
| DB 컬럼 | `snake_case` 단수형 (`created_at`) | |
| API JSON 필드 | `snake_case` | DB 컬럼과 표기가 일치해 변환 계층이 필요 없고, 백엔드↔프론트 간 필드명 추적이 쉽다 |
| 파일/디렉터리 (웹 프로젝트) | `kebab-case` | 대소문자 구분 없는 파일시스템(macOS 등)에서 안전 |
| 파일 (Python 모듈) | `snake_case` | import 가능해야 함 |
| 환경변수 | `UPPER_SNAKE_CASE` | 관례 |

🔲 결정 필요: DB 인덱스·제약조건 네이밍 (예: `ix_테이블_컬럼`, `fk_자식_부모`)

## 언어별 규칙

인스턴스는 별도 규칙을 두지 않는다 — 인스턴스도 변수이므로 해당 언어의 변수 규칙을 따른다.

| 항목 | Python | TypeScript/JS |
|---|---|---|
| 변수 | `snake_case` | `camelCase` |
| 함수/메서드/매개변수 | `snake_case` | `camelCase` |
| 클래스/인터페이스/타입 | `PascalCase` | `PascalCase` |
| 상수 | `UPPER_SNAKE_CASE` | `UPPER_SNAKE_CASE` |

- 컬렉션(배열, 리스트 등)을 담는 변수는 복수형으로 쓴다. (`users`, `items`)
- 표기는 CamelCase가 아니라 **PascalCase**라고 부른다. (CamelCase는 camelCase와 혼동됨)

## Boolean 접두사

`is_` / `has_` / `can_` 접두사를 사용한다. 표기는 해당 언어 규칙을 따른다.

- Python: `is_active`, `has_permission`, `can_edit`
- TypeScript: `isActive`, `hasPermission`, `canEdit`
- DB 컬럼: `is_active`

## 미결정 목록

- 🔲 DB 인덱스·제약조건 네이밍 규칙
- 🔲 테스트 파일 네이밍 (`*.test.ts` vs `*.spec.ts`, `test_*.py` 위치)
- 🔲 Git 브랜치 네이밍 (→ 확정되면 `git/`으로 이동)
