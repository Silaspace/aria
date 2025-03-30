all: wasm cli

wasm:
	go mod tidy
	GOOS=js GOARCH=wasm go build -o ./target/main.wasm ./wasm/main.go

cli:
	go mod tidy
	go build -o ./target/aria ./cli/

run:
	go mod tidy
	go run ./cli/main.go

clean:
	rm ./target/aria
	rm ./target/main.wasm
	rm -r *.hex
	rm -r *.txt

test:
	clear
	go mod tidy
	go build -o ./target/aria ./cli/main.go
	./target/aria $(FILE)