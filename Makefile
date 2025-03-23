watch:
	@air & \
	 templ generate -watch air
     
tw:
	npx @tailwindcss/cli -i ./static/css/input.css -o ./static/css/output.css --watch

run:
	@go run main.go

build:
	@go build -o tmp/main