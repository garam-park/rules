# 작명 규칙 (Naming)

> 상태: 확정 (2026-07-11) — Enum 값 표기만 보류(하단 참고).
> API URL 경로 규칙은 별도의 REST API 규칙 문서에서 다룬다.

## 요약표

| 항목 | 표기 | 예시 |
|---|---|---|
| DB 테이블 | `snake_case` 복수형 | `users`, `order_items` |
| DB 조인 테이블 (다대다) | 단수명 알파벳순 연결 | `role_user` |
| DB 컬럼 | `snake_case` 단수형 + 타입 접미사 | `email`, `created_at`, `user_id` |
| DB 인덱스·제약조건 | 접두사 + 테이블 + 컬럼 | `ix_users_email`, `fk_orders_users` |
| API JSON 필드 | `snake_case` (컬렉션은 복수형) | `user_id`, `order_items` |
| 변수·매개변수 | `snake_case` (컬렉션은 복수형) | `user_name`, `active_users` |
| 함수·메서드 | `camelCase` + 표준 동사 | `getUser()`, `sendEmail()` |
| 클래스·인터페이스·타입 | `PascalCase` | `User`, `PaymentService` |
| 상수 | `UPPER_SNAKE_CASE` | `MAX_RETRY_COUNT` |
| Boolean | `is_` / `has_` / `can_` 접두사 | `is_active`, `hasPermission()` |
| private/내부용 멤버 | `_` 접두사 (모든 언어) | `_cache`, `_buildQuery()` |
| 약어 | 일반 단어처럼 취급 | `HttpServer`, `user_id`, `parseUrl()` |
| 단위가 있는 값 | 단위 접미사 필수 | `timeout_ms`, `max_size_mb` |
| 파일·디렉터리 | `kebab-case` (Python 모듈은 `snake_case`) | `user-service.ts`, `user_service.py` |
| 테스트 파일 | `*.test.ts` / `tests/test_*.py` | `user-service.test.ts` |
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
- **약어는 일반 단어처럼 취급한다.** `Http`, `Api`, `Id`, `Url`로 쓴다.
  연속 약어에서도 단어 경계가 명확해지기 때문 (`XmlHttpRequest`).
  - `HttpServer`, `ApiClient`, `user_id`, `parseUrl()` (O)
  - `HTTPServer`, `APIClient`, `user_ID`, `parseURL()` (X)

## DB

- **테이블**: `snake_case` 복수형. 테이블은 행(row)의 집합이기 때문.
  - `users`, `orders`, `order_items`
- **조인 테이블 (다대다)**: 두 테이블의 **단수명을 알파벳순으로 연결**한다.
  순서 규칙이 고정되어 이름을 고민할 필요가 없다. (복수형 규칙의 예외)
  - `users` + `roles` → `role_user`
  - 관계 자체가 도메인 개념이 되면(가입, 수강 등) 일반 테이블로 승격하고
    의미 있는 이름을 붙인다 — `memberships`, `enrollments`
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

- **변수·매개변수**: `snake_case`, 모든 언어 공통. 매개변수도 변수이므로
  같은 규칙을 따른다. 컬렉션(배열·리스트 등)은 복수형.
  - `user_name`, `active_users`, `sendEmail(recipient_address, retry_count)`
- **함수·메서드**: `camelCase`. 변수(snake_case)와 한눈에 구분하기 위함.
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
- **private/내부용 멤버**: 모든 언어에서 `_` 접두사로 통일한다.
  이름만 보고 내부용임을 알 수 있게 하기 위함. TS의 `private` 키워드 등
  언어 기능은 접두사와 병행해서 사용한다.
  - `_cache`, `_retry_count`, `_buildQuery()`
- **단위가 있는 값**: 단위 접미사를 반드시 붙인다. 호출부에서
  단위 착각으로 인한 버그를 막기 위함. 타입 자체가 단위를 표현하는 경우
  (`Duration` 등)는 예외.
  - `timeout_ms`, `delay_seconds`, `max_size_mb` (O) / `timeout` (X)

### 표준 동사

함수 이름의 동사는 아래 의미로 고정한다. 이름만 보고 동작을 예측할 수
있게 하기 위함.

| 동사 | 의미 |
|---|---|
| `get` | 반드시 존재하는 것을 반환. 없으면 에러 |
| `find` | 없을 수 있는 것을 탐색. 없으면 null 반환 |
| `fetch` | 원격·IO를 통해 가져옴 (네트워크, 파일 등) |
| `list` | 목록(컬렉션)을 반환 |
| `create` | 새로 만들어 저장 |
| `update` | 기존 것을 수정 |
| `delete` | 삭제 |

- `getUser(id)` — 없으면 예외 / `findUser(email)` — 없으면 null
- `fetchExchangeRate()` — 외부 API 호출 / `listOrders(user_id)`

컬렉션 조회에 `get` + 복수형(`getUsers`)을 쓰지 않는다. 컬렉션은 0건이
정상 상황이라 `get`의 계약(없으면 에러)과 충돌하기 때문. `list`는
"나열하다"라는 동사이며, 0건이면 에러 없이 빈 컬렉션을 반환한다.
`getUser`/`getUsers`처럼 한 글자 차이로 갈리는 것도 피할 수 있다.

## 파일·환경

- **파일·디렉터리**: `kebab-case`. macOS 등 대소문자 무구분 파일시스템에서
  안전하기 때문. 단 Python 모듈은 import 가능해야 하므로 `snake_case`.
  - `user-service.ts`, `api-client/`, `user_service.py`
- **테스트 파일**: 생태계 기본 패턴을 따른다 (도구 설정 불필요).
  - TypeScript: 대상 파일 옆에 `*.test.ts` — `user-service.test.ts`
  - Python: `tests/` 디렉터리에 `test_*.py` — `tests/test_user_service.py`
- **환경변수**: `UPPER_SNAKE_CASE`.
  - `DATABASE_URL`, `API_KEY`

## 보류 항목

- 🔲 **Enum 값(상태 문자열) 표기** — 추가 논의 필요.
  현재 의견: 코드와 저장값 모두 UPPER (`OrderStatus.PENDING = "PENDING"`).
  이름(naming)이 아니라 값(value)의 문제라는 관점. 단, DB·JSON 필드의
  snake_case 규칙과 값 표기가 달라지는 트레이드오프가 있어 결정 전 재논의.
