git_user_id := mostafa
git_repo_id := go-api-client

# If env-vars are not set, use default values
GIT_USER_ID ?= $(git_user_id)
GIT_REPO_ID ?= $(git_repo_id)

# Check if go is installed
.PHONY: is-gofmt-installed
is-go-installed:
	@command -v go >/dev/null 2>&1 || { echo >&2 "Go 1.22+ is required but it's not installed. Aborting."; exit 1; }

# Generate OpenAPI client
.PHONY: generate-openapi-client
generate-openapi-client: is-go-installed
	@GO_POST_PROCESS_FILE="gofmt -w" docker run --rm \
		-v "./:/local" \
		-v "./openapi/templates:/openapi/templates" \
		openapitools/openapi-generator-cli generate \
		-i https://openapi.logto.io/source.json \
		-g go \
		--package-name logto \
		--template-dir=/openapi/templates \
		--enable-post-process-file \
		-p withGoMod=true \
		-p goImportAlias=logto \
		--git-user-id=$(GIT_USER_ID) \
		--git-repo-id=$(GIT_REPO_ID) \
		-o /local/
	@go mod tidy
