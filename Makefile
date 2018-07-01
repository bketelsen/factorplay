wasm:
	GOROOT=~/gowasm GOARCH=wasm GOOS=js ~/gowasm/bin/go build -o example.wasm .

packr:
	cd div && packr && cd ..

build-server:
	go build -o server-app server/server.go

run: build-server packr wasm
	./server-app