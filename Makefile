build: main.go
	mkdir bin
	go build -o bin/html main.go

run: main.go
	go run main.go

install: bin/html
	mv bin/html /usr/local/bin

uninstall: /usr/local/bin/html
	rm -f /usr/local/bin/html

clean:
	rm -rf index.html bin
