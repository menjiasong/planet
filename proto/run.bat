protoc -I. -ID:\gopath\bin\protoc-3.15.8-win64\include\google\protobuf  -ID:\gopath\pkg\mod\github.com\grpc-ecosystem\grpc-gateway@v1.16.0\third_party\googleapis  -ID:\gopath\pkg\mod\github.com\envoyproxy\protoc-gen-validate@v0.5.1 --go_out=plugins=grpc:./../pb --validate_out=lang=go:./../pb *.proto
protoc -I. -ID:\gopath\bin\protoc-3.15.8-win64\include\google\protobuf  -ID:\gopath\pkg\mod\github.com\grpc-ecosystem\grpc-gateway@v1.16.0\third_party\googleapis  --grpc-gateway_out=logtostderr=true:./../pb *.proto
protoc -I. -ID:\gopath\bin\protoc-3.15.8-win64\include\google\protobuf  -ID:\gopath\pkg\mod\github.com\grpc-ecosystem\grpc-gateway@v1.16.0\third_party\googleapis  --swagger_out=logtostderr=true:./../doc *.proto