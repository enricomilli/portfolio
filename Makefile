install:
	@go install github.com/a-h/templ/cmd/templ@latest
	@go mod tidy
	@go mod download
	@npm install -g tailwindcss



dev:
	@$(MAKE) css
	@$(MAKE) templ
	@go run .

css: 
	@tailwindcss -i site/assets/tailwind.css -o static/css/tailwind.css

templ:
	@TEMPL_EXPERIMENT=rawgo templ generate


build: 
	@$(MAKE) css
	@$(MAKE) templ
	@go build -o ./bin/portfolio

run:
	@$(MAKE) build
	@./bin/portfolio