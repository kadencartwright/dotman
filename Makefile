build:
	@go build -o bin/dotman
run: build
	./bin/dotman
install: build
	sudo cp ./bin/dotman /usr/bin/
