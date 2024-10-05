
dev:
	@$(MAKE) css
	@$(MAKE) templ
	@go run .

css: 
	@tailwindcss -i site/assets/styles.css -o static/css/styles.css

templ:
	@TEMPL_EXPERIMENT=rawgo templ generate


build: 
	@$(MAKE) css
	@$(MAKE) templ
	@go build -o ./bin/portfolio

run:
	@$(MAKE) build
	@./bin/portfolio