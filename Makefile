build-linux-amd64:
	mkdir -p bin
	rm -rf bin/* | true
	GOOS=linux GOARCH=amd64 go build -o bin/module module/main.go

build-linux-arm64:
	mkdir -p bin
	rm -rf bin/* | true
	GOOS=linux GOARCH=arm64 go build -o bin/module module/main.go
	
build-darwin-amd64:
	mkdir -p bin
	rm -rf bin/* | true
	go build -o bin/module module/main.go

get-data:
	cp -R .artifact bin/.artifact
	artifact pull

package-darwin-amd64: build-darwin-amd64 get-data
	tar -czf module.tar.gz bin .artifact

package-linux-amd64: build-linux-amd64 get-data
	tar -czf module.tar.gz bin .artifact

package-linux-arm64: build-linux-arm64 get-data
	tar -czf module.tar.gz bin .artifact