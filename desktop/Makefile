react: app/src/*
	cd app && yarn build

copy: react
	rm -rf resources/app
	cp -rf app/build/ resources/app

bundle: copy
	astilectron-bundler

run: bundle
	go run *.go
