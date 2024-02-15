build:
	mkdir -p bin
	rm -rf bin/* | true
	go build -o bin/module module/main.go

get-data:
	artifact pull

package: build get-data
	tar -czf module.tar.gz bin .artifact