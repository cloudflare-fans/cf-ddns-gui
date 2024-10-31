.PHONY: all
all: clean build-darwin-arm64 build-darwin-amd64

.PHONY: macos-pack-prepare
macos-pack-prepare:
	mkdir -p ./target/macOS_arm64/cfDDNS.app
	mkdir -p ./target/macOS_amd64/cfDDNS.app
	cp -r ./build/macOS_arm64/cfDDNS.app.template/* ./target/macOS_arm64/cfDDNS.app/
	cp -r ./build/macOS_arm64/cfDDNS.app.template/* ./target/macOS_amd64/cfDDNS.app/
	ln -s /Applications/ ./target/macOS_arm64/Applications
	ln -s /Applications/ ./target/macOS_amd64/Applications

.PHONY: build-darwin-arm64
build-darwin-arm64: macos-pack-prepare
	CGO_ENABLED=1 GOOS=darwin GOARCH=arm64 go build -o target/macOS_arm64/cfDDNS.app/Contents/MacOS/cf-ddns
	hdiutil create -srcFolder ./target/macOS_arm64/ -o ./target/cfDDNS.arm64.dmg

.PHONY: build-darwin-amd64
build-darwin-amd64: macos-pack-prepare
	CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build -o target/macOS_amd64/cfDDNS.app/Contents/MacOS/cf-ddns
	hdiutil create -srcFolder ./target/macOS_amd64/ -o ./target/cfDDNS.amd64.dmg

.PHONY: clean
clean:
	rm -rf ./target/*
