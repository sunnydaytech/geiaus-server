
all:
	chmod +x ./build/*.sh
	./build/init.sh
	go get github.com/golang/protobuf/proto
	./build/genrpc.sh
	go install github.com/sunnydaytech/geiaus/service

clean:
	rm -rf .download
	rm -rf .bin

