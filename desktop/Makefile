react: app/src/*
	cd app && yarn build

copy: react
	cp -rf app/build/ resources/app

bundle: copy
	astilectron-bundler

run: bundle
	go run *.go
