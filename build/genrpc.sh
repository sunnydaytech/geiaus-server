protoc=".bin/protoc"

# download protoc binary.
if [ ! -f "$protoc" ]
then
  if [ "$(uname)" == "Darwin" ]; then
    echo "osx"
    wget -O .download/protoc.zip https://github.com/google/protobuf/releases/download/v3.0.0-beta-2/protoc-3.0.0-beta-2-osx-x86_64.zip
  else
    echo "linux"
    wget -O .download/protoc.zip https://github.com/google/protobuf/releases/download/v3.0.0-beta-3/protoc-3.0.0-beta-3-linux-x86_64.zip
  fi
  unzip .download/protoc.zip -d .bin
fi

$protoc -I proto/ proto/*.proto --go_out=plugins=grpc:proto
