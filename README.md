# rules

개인 개발 규칙(convention & rules)을 정리하는 저장소입니다.

프로젝트를 진행하며 반복적으로 지키고 싶은 원칙, 코딩 컨벤션, 커밋 규칙 등을
한곳에 모아두고 계속 다듬어 나가는 것이 목적입니다.

## 구조

```text
rules/
├── .github/
│   └── workflows/
│       └── ci.yml             # Markdown lint CI
├── .markdownlint.jsonc        # Markdown lint 규칙
├── README.md                  # 저장소 소개
├── git/
│   ├── commit-message.md      # 커밋 메시지 (Conventional Commits, 한국어)
│   └── branch.md              # 브랜치 전략 (Git Flow), 네이밍, 머지
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

- `main`, `develop` 브랜치 push와 PR에서 Markdown lint CI를 실행한다.
- Markdown lint 규칙은 `.markdownlint.jsonc`에 둔다.
