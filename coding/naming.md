# 작명 규칙 (Naming)

> 상태: 확정 (2026-07-11) — 11개 항목을 하나씩 결정하여 정리함.

## 요약표

| 항목 | 표기 | 예시 |
|---|---|---|
| DB 테이블 | `snake_case` 복수형 | `users`, `order_items` |
| DB 컬럼 | `snake_case` 단수형 + 타입 접미사 | `email`, `created_at`, `user_id` |
| DB 인덱스·제약조건 | 접두사 + 테이블 + 컬럼 | `ix_users_email`, `fk_orders_users` |
| API JSON 필드 | `snake_case` (컬렉션은 복수형) | `user_id`, `order_items` |
| 변수 | `snake_case` (컬렉션은 복수형) | `user_name`, `active_users` |
| 함수·메서드 | `camelCase` | `getUser()`, `sendEmail()` |
| 클래스·인터페이스·타입 | `PascalCase` | `User`, `PaymentService` |
| 상수 | `UPPER_SNAKE_CASE` | `MAX_RETRY_COUNT` |
| Boolean | `is_` / `has_` / `can_` 접두사 | `is_active`, `hasPermission()` |
| 파일·디렉터리 | `kebab-case` (Python 모듈은 `snake_case`) | `user-service.ts`, `user_service.py` |
| 환경변수 | `UPPER_SNAKE_CASE` | `DATABASE_URL` |

## 기본 원칙

- **표기(case)로 종류를 구분한다.** 이름만 보고 변수(snake_case)인지,
  함수(camelCase)인지, 클래스(PascalCase)인지, 상수(UPPER_SNAKE_CASE)인지
  알 수 있어야 한다.
- 이 규칙은 **모든 언어에 동일하게** 적용한다. 언어 생태계 관례
  (TS의 camelCase 변수, Python의 snake_case 함수)와 다른 부분이 있지만,
  종류 구분의 일관성을 우선한다. 필요하면 린터 설정을 규칙에 맞춘다.
- 외부 라이브러리·프레임워크가 강제하는 이름(오버라이드 메서드,
  프레임워크 훅 등)은 예외로 두고 해당 API의 표기를 따른다.

## DB

- **테이블**: `snake_case` 복수형. 테이블은 행(row)의 집합이기 때문.
  - `users`, `orders`, `order_items`
- **컬럼**: `snake_case` 단수형. 컬럼은 하나의 값을 담기 때문.
  - 날짜·시각: `_at`(시각), `_on`(날짜) 접미사 — `created_at`, `expired_on`
  - 외래키: `_id` 접미사 — `user_id`, `order_id`
  - Boolean: `is_` 등 접두사 — `is_active`
- **인덱스·제약조건**: `접두사_테이블_컬럼`. 에러 메시지에 이름만 보여도
  무엇이 깨졌는지 알 수 있게 하기 위함.
  - `pk_` 기본키, `fk_자식테이블_부모테이블` 외래키, `uq_` 유니크,
    `ix_` 인덱스, `ck_` 체크
  - `ix_users_email`, `uq_users_email`, `fk_orders_users`, `ck_orders_status`

## API

- **JSON 필드**: `snake_case`. DB 컬럼과 표기가 일치해 변환 계층이 필요
  없고, 백엔드↔프론트 간 필드명 추적이 쉽다.
- 변수와 같은 규칙으로 단일 값은 단수형, 컬렉션(배열)은 복수형.
  - `{ "user_id": 1, "order_items": [...] }`

## 코드

- **변수**: `snake_case`, 모든 언어 공통. 컬렉션(배열·리스트 등)은 복수형.
  - `user_name`, `active_users`
- **함수·메서드·매개변수 아닌 호출 가능한 것**: `camelCase`.
  변수(snake_case)와 한눈에 구분하기 위함.
  - `getUser()`, `sendEmail()`
- **클래스·인터페이스·타입**: `PascalCase`. `I` 같은 접두사는 붙이지 않는다.
  (표기 명칭에 주의: CamelCase가 아니라 PascalCase다. CamelCase는
  camelCase와 혼동됨)
  - `User`, `OrderItem`, `PaymentService`
- **상수**: `UPPER_SNAKE_CASE`. 모듈 최상위의 진짜 고정값에만 적용한다.
  - `MAX_RETRY_COUNT`, `API_TIMEOUT`
- **Boolean**: 상태는 `is_`, 소유는 `has_`, 가능 여부는 `can_` 접두사.
  표기는 해당 항목 규칙을 따른다.
  - 변수: `is_active`, `has_permission` / 함수: `isValid()`, `canEdit()`
  - DB 컬럼·JSON 필드: `is_active`

## 파일·환경

- **파일·디렉터리**: `kebab-case`. macOS 등 대소문자 무구분 파일시스템에서
  안전하기 때문. 단 Python 모듈은 import 가능해야 하므로 `snake_case`.
  - `user-service.ts`, `api-client/`, `user_service.py`
- **환경변수**: `UPPER_SNAKE_CASE`.
  - `DATABASE_URL`, `API_KEY`
