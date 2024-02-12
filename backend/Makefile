rundb:
	docker run -d --name mymongodb \
  -p 5436:27017 \
  -e MONGO_INITDB_ROOT_USERNAME=admin \
  -e MONGO_INITDB_ROOT_PASSWORD=123321 \
  -d --rm\
  -v db-team-work-space\
  mongo

docsinit:
	swag init -g ./cmd/main.go -o ./docs

runtests:
	go test -v ./...