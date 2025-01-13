binary-name=nes-cards

build: templ-build
	@GOOS=windows GOARCH=amd64 go build -o ./bin/${binary-name}-win.exe ./cmd/main.go
	@GOOS=linux GOARCH=amd64 go build -o ./bin/${binary-name}-linux ./cmd/main.go
	@GOOS=darwin GOARCH=amd64 go build -o ./bin/${binary-name}-darwin ./cmd/main.go

run: build
	@./bin/${binary-name}-linux

kill-services:
	@lsof -t -i:3000 | xargs -r kill

test:
	@go test ./cmd/main.go

clean:
	@rm -rf ./bin/*
	@go clean

db-clean:
	@rm -rf ~/.local/share/nes-cards/data.db

css-build:
	@tailwindcss -i ./public/static/css/input.css -o ./public/static/css/style.css

css-watch:
	@tailwindcss -i ./public/static/css/input.css -o ./public/static/css/style.css --watch

templ-build:
	@templ generate

templ-watch:
	@templ generate --watch
