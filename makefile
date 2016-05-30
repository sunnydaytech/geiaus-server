
all: 
	chmod +x ./build/*.sh
	./build/init.sh
	./build/genrpc.sh
clean:
	rm -rf .download
	rm -rf .bin
install: all
	go get github.com/golang/protobuf/proto
	go get golang.org/x/crypto/bcrypt
	go get golang.org/x/net/context
	go get google.golang.org/grpc
	go install .
test: all 
	go test server/*_test.go

