all: build

build:
	cd front && npm run build
	go build -o evexp cmd/main.go

build-backend:
	go build -o evexp cmd/main.go

run:
	source worker/env/bin/activate && ./evexp
