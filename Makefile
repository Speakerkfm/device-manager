start:
	export `cat .env` && go run cmd/device-manager/*.go --port=8000

flatten:
	mkdir -p tmp
	swagger -o=tmp/flatten_log.txt flatten swagger/swagger.yaml > tmp/swagger.yaml && swagger validate tmp/swagger.yaml

server: flatten
	rm -r pkg/restapi/
	swagger generate server -f tmp/swagger.yaml -P models.AuthKey -t pkg --tags=users --tags=devices --exclude-main

doc: flatten
	swagger serve -p 8095 -F swagger tmp/swagger.yaml

iface:
	ifacemaker -f pkg/store -s Store -i StoreInterface -p interfaces -o pkg/interfaces/store_interface.go
	ifacemaker -f pkg/service/jwt.go -s JWT -i JWT -p interfaces -o pkg/interfaces/jwt.go
	ifacemaker -f pkg/service/user.go -s UserService -i UserService -p interfaces -o pkg/interfaces/user.go
	ifacemaker -f pkg/service/device.go -s DeviceService -i DeviceService -p interfaces -o pkg/interfaces/device.go

mock: iface
	mockery -dir pkg/store -all -output pkg/mocks -case underscore
	mockery -dir pkg/service/serviceiface -all -output pkg/mocks -case underscore

docker:
	GOOS=linux GOARCH=amd64 go build -v cmd/device-manager/*.go
	docker build -t device-manager .

unittest:
	go test -failfast `go list ./... | grep -e /pkg/service -e /pkg/store` -v -coverprofile .testCoverage.txt
	go tool cover -func=.testCoverage.txt

test:
	export `cat .env` && make unittest

coverage:
	go tool cover -html=.testCoverage.txt
