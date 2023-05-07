test:
	go test -count=1 ./...

debug:
	go test -count=1 ./... -v

coverage:
	go test ./... -cover

run_dev:
	PORT=8443 go run main.go

build:
	go build -ldflags="-w -s" -o build

run_build:
	go run ./build

package:
	docker build -t gin-server -f Dockerfile .

run_package:
	docker run -p8443:8443 -e PORT=8443 gin-server

dive:
	CI=true dive gin-server --ci-config docker/.dive.yaml

clean:
	rm -rf build
