init:
	go mod init skeleton

tidy:
	go mod tidy
	
gen:
	protoc --proto_path=proto --go_out=paths=source_relative:./pb --go-grpc_out=paths=source_relative:./pb proto/*/*.proto
	
.PHONY: init gen