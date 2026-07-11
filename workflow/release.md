# 버전·릴리스 규칙

> 상태: 확정 (2026-07-11)
> 브랜치 흐름은 `git/branch.md`(Git Flow)를 따른다.

## 버전: SemVer

`MAJOR.MINOR.PATCH` — 버전 번호 자체가 호환성 정보를 담는다.

| 자리 | 올리는 경우 | 커밋 타입과의 관계 |
|---|---|---|
| MAJOR | 깨지는 변경 (breaking change) | `!` 붙은 커밋 |
| MINOR | 하위 호환 기능 추가 | `feat` |
| PATCH | 하위 호환 버그 수정 | `fix` |

- Conventional Commits를 쓰므로 커밋 로그에서 다음 버전을 기계적으로
  결정할 수 있다.

## 릴리스 절차

1. `develop`에서 `release/x.y.z` 브랜치를 딴다.
2. 버전 번호 반영, 최종 점검. 수정은 release 브랜치에서.
3. `main`으로 머지(--no-ff)하고 **`vx.y.z` 태그**를 남긴다.
4. `develop`에도 머지해서 release 중 수정을 회수한다.
5. 배포는 main의 태그 기준으로 한다.

- hotfix도 동일: `main` → `hotfix/*` → `main`(+태그, PATCH 상승) +
  `develop` 머지.

## 체인지로그

- 태그 사이의 커밋 로그가 체인지로그다. Conventional Commits 덕에
  `feat`/`fix`만 뽑으면 릴리스 노트가 된다. 별도 CHANGELOG 파일은
  필요해질 때 도구(release-please 등)로 자동화한다.
