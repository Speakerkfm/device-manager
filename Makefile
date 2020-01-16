run:
	export `cat .env` && go run cmd/device-manager/*.go --port=8000

flatten:
	mkdir -p tmp
	swagger -o=tmp/flatten_log.txt flatten swagger/swagger.yaml > tmp/swagger.yaml && swagger validate tmp/swagger.yaml

server: flatten
	rm -r pkg/restapi/
	swagger generate server -f tmp/swagger.yaml -P models.JWTUserKey -P models.JWTDeviceKey -t pkg --tags=users --tags=devices --exclude-main

doc: flatten
	swagger serve -p 8095 -F swagger tmp/swagger.yaml

iface:
	ifacemaker -f pkg/store -s Store -i StoreInterface -p store -o pkg/store/store_interface.go
	ifacemaker -f pkg/service/jwt.go -s JWT -i JWT -p serviceiface -o pkg/service/serviceiface/jwt.go
	ifacemaker -f pkg/service/user.go -s UserService -i UserService -p serviceiface -o pkg/service/serviceiface/user.go
	ifacemaker -f pkg/service/device.go -s DeviceService -i DeviceService -p serviceiface -o pkg/service/serviceiface/device.go