build: main.go
	mkdir bin
	go build -o bin/html main.go

run: bin/html
	./bin/html

clean:
	rm -rf index.html bin
