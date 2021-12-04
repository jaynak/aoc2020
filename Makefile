build: ## Build your project and put the output binary in build
	mkdir -p bin
	GO111MODULE=on go build -mod vendor -o ./bin/aoc2020 ./cmd/aoc2020

run: build
	./bin/aoc2020