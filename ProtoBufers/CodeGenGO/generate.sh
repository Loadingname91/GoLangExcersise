go get -u -v github.com/gogo/protobuf/proto
go get -u -v github.com/gogo/protobuf/protoc-gen-gogo
go get -u -v github.com/gogo/protobuf/gogoproto



protoc -I=src/ --gogo_out=src/ src/simple.proto