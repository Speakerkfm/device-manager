flatten:
	mkdir -p tmp
	swagger -o=tmp/flatten_log.txt flatten swagger/swagger.yaml > tmp/swagger.yaml && swagger validate tmp/swagger.yaml

server: flatten
	rm -r pkg/restapi/
	swagger generate server -f tmp/swagger.yaml -t pkg --tags=users --tags=devices --exclude-main

doc: flatten
	swagger serve -p 8095 -F swagger tmp/swagger.yaml

run:
	export `cat .env` && go run cmd/device-manager/*.go --port=8000

iface:
	ifacemaker -f pkg/store -s Store -i StoreInterface -p store -o pkg/store/store_interface.go