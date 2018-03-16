build:
	go build .

run: build
	./athens
	
cli: 
	go build -o athens ./cmd/cli

docs:
	cd docs && hugo

test:
	go test ./...
