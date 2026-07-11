# REST API 규칙

> 상태: 확정 (2026-07-11) — 응답 본문(성공·에러 구조)만 보류(하단 참고).
> 필드·enum 값 표기의 근거는 `naming.md`를 따른다.

## 요약표

| 항목 | 규칙 | 예시 |
|---|---|---|
| URL 경로 | `kebab-case` | `/order-items` |
| 리소스 이름 | 복수형 통일 | `/users`, `/users/{id}` |
| 쿼리 파라미터 | `snake_case` | `?created_after=...&user_id=1` |
| 쿼리의 enum 값 | 저장값 그대로 `UPPER` | `?status=IN_PROGRESS` |
| 수정 | PATCH(부분) + PUT(전체) 병용 | |
| 중첩 리소스 | 깊이 제한 없음 | `/users/{id}/orders/{id}/items` |
| 액션 | `POST /리소스/{id}/동사` | `POST /orders/{id}/cancel` |
| 페이지네이션 | 오프셋 방식 | `?page=2&per_page=20` |
| 버저닝 | 헤더 `X-Api-Version` | `X-Api-Version: 1` |

## URL

- **경로는 `kebab-case`.** 경로는 데이터 필드가 아니라 위치(path)이므로
  파일·디렉터리 규칙과 같은 계열로 묶는다: **경로는 kebab, 데이터는
  snake** (테이블·컬럼·JSON 필드·쿼리 파라미터).
  - 문서·브라우저에서 링크 밑줄에 언더스코어가 묻히는 문제
    (`/order_items` → `/order items`처럼 보임)도 피한다.
  - 호스트명은 언더스코어가 허용되지 않으므로 kebab이면 URL 전체가
    한 가지 구분자로 통일된다.
  - 같은 구조의 선례: Zalando REST 가이드라인 (경로 kebab, 쿼리·JSON snake)
- **리소스 이름은 복수형으로 통일한다.** 컬렉션이든 개별 조회든 항상 같다.
  DB 테이블 복수형 규칙과 일치.
  - `GET /users`, `GET /users/{id}`, `POST /users`
- **중첩 리소스는 깊이 제한 없이 허용한다.** 소유·소속 관계를 URL 계층으로
  그대로 드러낸다.
  - `/users/{id}/orders`, `/users/{id}/orders/{id}/items`
- **쿼리 파라미터는 `snake_case`.** 쿼리는 경로가 아니라 데이터 필드이므로
  JSON 필드와 같은 표기를 쓴다.
  - enum 값은 변환 없이 저장값 그대로 쓴다 — `?status=IN_PROGRESS`
  - `?created_after=2026-01-01&user_id=1&page=2`

## HTTP 메서드

`naming.md`의 표준 동사와 다음처럼 대응한다.

| 메서드 | 용도 | 표준 동사 |
|---|---|---|
| GET (컬렉션) | 목록 조회 | `list` |
| GET (단일) | 개별 조회 | `get` |
| POST | 생성 | `create` |
| PATCH | 부분 수정 — 보낸 필드만 반영 | `update` |
| PUT | 전체 교체 — 리소스 전부를 제출 | `update` (교체) |
| DELETE | 삭제 | `delete` |

- **PATCH와 PUT은 표준 의미대로 병용한다.** 일반적인 수정은 PATCH,
  "전체를 통째로 바꾼다"는 의도가 있는 곳은 PUT. 하나의 리소스가 둘 다
  제공해도 된다.

## 액션 (CRUD로 표현되지 않는 동작)

- **`POST /리소스/{id}/동사`** 형태로 표현한다. URL에서 동사를 허용하는
  유일한 예외. 상태 변경으로 우겨넣는 것(PATCH로 status만 변경)보다
  의도가 명확하고, 부수효과(알림 발송 등)가 있는 동작을 자연스럽게 담는다.
  - `POST /orders/{id}/cancel`, `POST /payments/{id}/approve`,
    `POST /emails/{id}/resend`

## 페이지네이션

- **오프셋 방식**: `?page=2&per_page=20`. 1부터 시작한다.
  - 구현이 단순하고 임의 페이지 이동이 가능하다.
  - 대용량·실시간 피드처럼 오프셋의 성능·정합성 한계가 실제로 문제 되는
    엔드포인트는 예외적으로 커서 방식을 허용하되, 문서에 명시한다.

## 버저닝

- **`X-Api-Version` 헤더로 정수 버전을 지정한다.** URI는 리소스의
  식별자이고 버전은 표현의 문제이므로, URI를 버전과 무관하게 유지한다.
  (선례: GitHub `X-GitHub-Api-Version`)
  - `X-Api-Version: 1`
- 헤더 방식의 약점은 규칙으로 보완한다:
  - **기본값**: 헤더가 없으면 최신이 아니라 **서버가 고정해둔 기본
    버전**으로 처리한다. 최신으로 따라가면 클라이언트가 소리 없이 깨진다.
  - **가시성**: 서버는 실제 적용한 버전을 응답 헤더 `X-Api-Version`으로
    되돌려주고, 액세스 로그에 남긴다.
- 호환 가능한 변경(필드 추가 등)은 버전을 올리지 않는다. 버전 상승은
  깨지는 변경(breaking change)에만 사용한다.

## 보류 항목

- 🔲 **응답 본문 구조 (성공·에러)** — envelope 여부, 에러 코드 체계 등
  응답 설계 전반을 별도로 깊게 논의 후 결정.
  참고: https://blog.storyg.co/rest-api-response-body-best-pratics
