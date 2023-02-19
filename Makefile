.PHONY: build

build:
	@echo "building..."
	@go build -o bin/pokedex.exe .\cmd\webserver\