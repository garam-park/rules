# rules

개인 개발 규칙(convention & rules)을 정리하는 저장소입니다.

프로젝트를 진행하며 반복적으로 지키고 싶은 원칙, 코딩 컨벤션, 커밋 규칙 등을
한곳에 모아두고 계속 다듬어 나가는 것이 목적입니다.

## 구조

```text
rules/
├── .dagger/                   # Dagger 파이프라인 (lint 함수)
├── .github/
│   └── workflows/
│       └── ci.yml             # CI — Dagger 함수를 호출하는 얇은 래퍼
├── .markdownlint.jsonc        # Markdown lint 규칙
├── dagger.json                # Dagger 모듈 설정
├── README.md                  # 저장소 소개
├── git/
│   ├── branch.md              # 브랜치 전략, 네이밍, 동기화, 머지
│   ├── commit.md              # 커밋 단위와 되돌리기
│   └── commit-message.md      # 커밋 메시지 (Conventional Commits, 한국어)
├── coding/
│   ├── naming.md              # 작명 규칙 (변수·함수·DB·파일 등)
│   ├── project-structure.md   # 프로젝트 구조 (도메인형, 3계층, 의존 방향)
│   ├── rest-api.md            # REST API (URL, 메서드, 응답 구조)
│   ├── test.md                # 테스트 (구조, 이름, 비중, mock)
│   └── error-logging.md       # 에러 처리·로깅
└── workflow/
    ├── pull-request.md        # PR·셀프 리뷰
    └── release.md             # 버전(SemVer)·릴리스 절차
```

## 작성 원칙

- 규칙은 하나의 마크다운 문서로 작성한다.
- 규칙마다 **왜 이 규칙이 필요한지(이유)** 를 함께 기록한다.
- 지키기 어렵거나 불필요해진 규칙은 과감히 수정하거나 삭제한다.

## 자동 검증

- 검증 로직은 Dagger 파이프라인(`.dagger/`)에 두고, 로컬과 CI가
  **같은 컨테이너에서 같은 함수를 실행**한다. 로컬에서 통과하면
  CI에서도 통과한다.
- 로컬 실행: `dagger call lint` (Docker 필요)
- CI: `main`, `develop` push와 PR에서 같은 함수를 호출한다.
- Markdown lint 규칙은 `.markdownlint.jsonc`에 둔다.
