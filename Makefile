.PHONY: proto
proto:
	protoc \
		-I . \
		-I ${GOPATH}/src \
		-I ${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate \
		--go_out=:. \
		--validate_out="lang=go:." \
		example/example.proto