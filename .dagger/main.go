// rules 저장소의 CI 파이프라인
//
// 검증 로직을 컨테이너 안에서 실행해 로컬과 GitHub Actions에서
// 동일한 결과를 보장한다. 로컬: `dagger call lint`
package main

import (
	"context"

	"dagger/rules/internal/dagger"
)

// markdownlint-cli2 공식 이미지. CI와 로컬이 같은 버전을 쓰도록 고정한다.
const markdownlintImage = "davidanson/markdownlint-cli2:v0.23.0"

type Rules struct{}

// Lint는 저장소의 모든 Markdown 문서를 markdownlint-cli2로 검사한다.
// 규칙은 저장소 루트의 .markdownlint.jsonc를 그대로 사용한다.
func (m *Rules) Lint(
	ctx context.Context,
	// 검사할 저장소 루트 (기본: 이 저장소)
	// +defaultPath="/"
	// +ignore=[".git"]
	source *dagger.Directory,
) (string, error) {
	return dag.Container().
		From(markdownlintImage).
		WithMountedDirectory("/workdir", source).
		WithWorkdir("/workdir").
		WithExec([]string{"markdownlint-cli2", "**/*.md"}).
		Stdout(ctx)
}
