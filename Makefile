run:
	templ generate
	go build -o ./tmp ./cmd/main.go && air
client:
	npm i
css:
	npx tailwindcss -i static/styles/main.css -o static/styles/tw.css --watch
