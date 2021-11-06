compile:
		protoc proto/*.proto --go_out=./proto --proto_path=.
test:
		go test -race ./...
