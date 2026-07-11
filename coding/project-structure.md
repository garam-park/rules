# 프로젝트 구조 규칙

> 상태: 확정 (2026-07-11)
> 디렉터리 이름 표기는 `naming.md`(kebab-case, Python은 snake_case)를 따른다.

## 레이아웃: 도메인형

- **기능(도메인)별로 묶고, 그 안에 계층을 둔다.** 기능 하나를 고칠 때
  한 폴더만 보면 되고, 기능 삭제가 폴더 삭제로 끝난다.

```
src/
├── orders/
│   ├── order-controller.ts
│   ├── order-service.ts
│   ├── order-repository.ts
│   └── order-model.ts
├── users/
│   └── ...
└── shared/
    └── ...
```

- 파일 이름에 도메인과 역할을 함께 쓴다 (`order-service.ts`).
  에디터 탭·검색에서 `service.ts` 여러 개가 구분되지 않는 문제를 피한다.

## 계층: controller → service → repository

| 계층 | 역할 |
|---|---|
| controller | 진입 경계. 요청 파싱·검증, 응답 변환. 비즈니스 로직 없음 |
| service | 비즈니스 로직. 트랜잭션 경계 |
| repository | 저장소 접근. 쿼리만, 판단 없음 |

- model(엔티티·스키마)은 도메인 폴더 안에 함께 둔다.
- 이보다 복잡한 구조(헥사고날 등)는 필요가 증명되기 전에는 쓰지 않는다.

## 의존 방향: 한 방향만, 건너뛰기 금지

- **controller → service → repository 방향만 허용한다.**
  - 역방향 금지: service가 controller를, repository가 service를 모른다.
  - 건너뛰기 금지: controller가 repository를 직접 호출하지 않는다.
    "단순 조회라서"라는 예외를 만들면 service 경유 여부가
    케이스바이케이스가 된다. 규칙은 기계적이어야 지켜진다.
- **다른 도메인은 service를 통해서만 접근한다.**
  `orders`가 사용자 정보가 필요하면 `user-service`를 호출한다.
  다른 도메인의 repository·model에 직접 접근하지 않는다 —
  도메인 내부 구조를 바꿔도 다른 도메인이 깨지지 않게 하기 위함.

## 공통 코드: shared/ + 입장 제한

- 여러 도메인이 쓰는 코드는 `shared/`에 둔다. 단 **입장 조건**이 있다:
  **도메인 지식이 없는 것만** — 로거, 날짜 유틸, 공통 에러 타입,
  envelope 직렬화 등.
- 도메인 로직이 shared로 새어 나가는 것을 금지한다. "orders와 users가
  둘 다 쓰니까 shared로"가 아니라, 두 도메인이 쓰는 도메인 로직이라면
  그것은 셋째 도메인이거나 한쪽 service의 공개 API다.
- `shared/`는 도메인을 import할 수 없다 (의존 방향: 도메인 → shared 단방향).
