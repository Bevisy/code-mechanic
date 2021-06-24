.PHONY: gen clean build gen2

gen: clean
	protoc -I=$$GOPATH/src -I=. \
	--go_out=. \
	--go_opt=paths=source_relative \
	--go-grpc_out=. \
	--go-grpc_opt=paths=source_relative \
	hello_service.proto

gen2: clean
	protoc -I=$$GOPATH/src -I=. \
	--go_out=. \
	--go-grpc_out=. \
	hello_service.proto

clean:
	rm -f hello_service.pb.go hello_service_grpc.pb.go