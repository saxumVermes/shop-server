build:
	go build -ldflags "-X main.GITCOMMIT=$$(git rev-parse --short=8 HEAD)" -o shoe-server .
