build: main.go
	mkdir bin
	go build -o bin/html main.go

run: main.go
	go run main.go

clean:
	rm -rf index.html bin
