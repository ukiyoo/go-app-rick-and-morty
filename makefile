build-work:
	go env -w GOARCH=wasm GOOS=js
	go build -o web/app.wasm ./app

	go env -w GOARCH=amd64 GOOS=windows
	go build -o start ./server


build-home:
	go env -w GOARCH=wasm GOOS=js
	go build -o web/app.wasm ./app

	go env -w GOARCH=amd64 GOOS=windows
	go build -o start ./server

	sass web/sass/custom.scss:web/custom.css

run-work: build-work
	./start

run: build-home
	./start