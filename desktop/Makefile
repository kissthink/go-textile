build:
	cd app && yarn build && \
	cp -rf build/ ../resources/app && \
	cd .. && \
	astilectron-bundler

run: build
	go run *.go
