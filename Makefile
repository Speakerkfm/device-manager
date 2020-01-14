flatten:
	mkdir -p tmp
	swagger -o=tmp/flatten_log.txt flatten swagger/swagger.yaml > tmp/swagger.yaml && swagger validate tmp/swagger.yaml

server: flatten
	rm -f pkg/restapi/configure_device_manager.go
	swagger generate server -f tmp/swagger.yaml -t pkg --tags=users --tags=devices --exclude-main

doc: flatten
	swagger serve -p 8095 -F swagger tmp/swagger.yaml