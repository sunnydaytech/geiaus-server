# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/sunnydaytech/geiaus-server

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)

RUN echo $OSTYPE
RUN apt-get update && apt-get install -y unzip
RUN go get -u github.com/golang/protobuf/proto
RUN go get -u github.com/golang/protobuf/protoc-gen-go
RUN cd /go/src/github.com/sunnydaytech/geiaus-server && make clean && make install

# Run the outyet command by default when the container starts.
ENTRYPOINT ["/go/bin/geiaus-server", "-gcloud_project_id", "nich01as-com"]

# Document that the service listens on port 5001.
EXPOSE 5001 
