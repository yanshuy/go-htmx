watch:
	npx @tailwindcss/cli -i ./static/css/input.css -o ./static/css/output.css --watch  & \
	templ generate --watch

run:
	@go run main.go

build:
	@go build -o tmp/main