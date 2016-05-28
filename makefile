
all: 
	chmod +x ./build/*.sh
	./build/init.sh
	go get github.com/golang/protobuf/proto
	./build/genrpc.sh

clean:
	rm -rf .download
	rm -rf .bin
test: all 
	go test server/*_test.go

