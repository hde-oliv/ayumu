MAIN_PACKAGE_PATH	:=	./cmd
BINARY_NAME			:=	ayumu

all: tidy run


tidy:
	@go fmt ./...
	@go mod tidy -v


audit:
	@go mod verify
	@go vet ./...
	@go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-ST1000,-U1000 ./...
	@go run golang.org/x/vuln/cmd/govulncheck@latest ./...
	@go test -race -buildvcs -vet=off ./...


test:
	@go test -v -race -buildvcs ./...


cover:
	@go test -v -race -buildvcs -coverprofile=/tmp/coverage.out ./...
	@go tool cover -html=/tmp/coverage.out


build:
	@go build -o=/tmp/bin/${BINARY_NAME} ${MAIN_PACKAGE_PATH}


run: build
	@/tmp/bin/${BINARY_NAME}


live:
	@go run github.com/cosmtrek/air@v1.43.0 \
		--build.cmd "make build" --build.bin "/tmp/bin/${BINARY_NAME}" --build.delay "100" \
		--build.exclude_dir "" \
		--build.include_ext "go, tpl, tmpl, html, css, scss, js, ts, sql, jpeg, jpg, gif, png, bmp, svg, webp, ico" \
		--misc.clean_on_exit "true"


prod: tidy audit no-dirty
	GOOS=linux GOARCH=amd64 go build -ldflags='-s' -o=/tmp/bin/linux_amd64/${BINARY_NAME} ${MAIN_PACKAGE_PATH}
	@upx -5 /tmp/bin/linux_amd64/${BINARY_NAME}
	@cp /tmp/bin/linux_amd64/${BINARY_NAME} .


.PHONY: prod live run build tidy audit test cover
