
all: 
	chmod +x ./build/*.sh
	./build/init.sh
	./build/genrpc.sh
clean:
	rm -rf .download
	rm -rf .bin
install: all
	go install .
test: all 
	go test server/*_test.go

