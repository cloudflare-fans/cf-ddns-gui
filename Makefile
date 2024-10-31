.PHONY: all
all: clean build-darwin-arm64 build-darwin-amd64

.PHONY: macos-pack-prepare
macos-pack-prepare:
	mkdir -p ./target/cfDDNS.arm64.app
	mkdir -p ./target/cfDDNS.amd64.app
	cp -r ./build/macOS/cfDDNS.arm64.app.template/* ./target/cfDDNS.arm64.app/
	cp -r ./build/macOS/cfDDNS.amd64.app.template/* ./target/cfDDNS.amd64.app/

.PHONY: build-darwin-arm64
build-darwin-arm64: macos-pack-prepare
	CGO_ENABLED=1 GOOS=darwin GOARCH=arm64 go build -o target/cfDDNS.arm64.app/Contents/MacOS/cf-ddns

.PHONY: build-darwin-amd64
build-darwin-amd64: macos-pack-prepare
	CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build -o target/cfDDNS.amd64.app/Contents/MacOS/cf-ddns

.PHONY: clean
clean:
	rm -rf ./target/*
