test:
	docker run -v ${PWD}/:/app -w /app golang:1.14-stretch sh -c \
	'go mod vendor && go test -mod vendor -bench=. -benchmem -coverprofile=coverage.out -v ./...'