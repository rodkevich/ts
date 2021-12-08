# makefile

#ticket:
#	 curl -d "@examples/createTicket.json" POST 0.0.0.0:12300/api/v1/tickets -v -H "Content-Type: application/json"

.Phony gen:
gen: gen-customer gen-photo gen-profile gen-sub gen-tag gen-ticket
	echo '... finished'

gen-customer:
	protoc \
 	--proto_path=. \
	--go_out=. \
	--go-grpc_out=. \
	--grpc-gateway_out=. \
	customer/proto/customer/v1/customer.proto


gen-photo:
	protoc \
	--proto_path=. \
	--go_out=. \
	--go-grpc_out=. \
	--grpc-gateway_out=. \
	photo/proto/photo/v1/photo.proto

gen-profile:
	protoc -I=. \
	--go_out=. \
	--go-grpc_out=. \
	--grpc-gateway_out=. \
	profile/proto/profile/v1/profile.proto

gen-sub:
	protoc -I=. \
	--go_out=. \
	--go-grpc_out=. \
	--grpc-gateway_out=. \
	sub/proto/sub/v1/sub.proto

gen-tag:
	protoc -I=. \
	--go_out=. \
	--go-grpc_out=. \
	--grpc-gateway_out=. \
	tag/proto/tag/v1/tag.proto

gen-ticket:
	protoc -I=. \
	--go_out=. \
	--go-grpc_out=. \
	--grpc-gateway_out=. \
	ticket/proto/ticket/v1/ticket.proto


#gen-gateway_out:
#	protoc -I=. --grpc-gateway_out . \
#    --grpc-gateway_opt logtostderr=true \
#    --grpc-gateway_opt paths=import \
#    --grpc-gateway_opt generate_unbound_methods=true \
#    proto/photo/v1/photo.proto


linter:
	echo "Starting linters"
	cd customer && golangci-lint run ./...
	cd ..
	cd photo && golangci-lint run ./...
	cd ..
	cd ticket && golangci-lint run ./...
	cd ..
	cd profile && golangci-lint run ./...

upgrade:
	# go get $(go list -f '{{if not (or .Main .Indirect)}}{{.Path}}{{end}}' -m all)
	go get -u -t -d -v ./...
	go mod tidy
	go mod vendor

FILES := $(shell docker ps -aq)

down-local:
	docker stop $(FILES)
	docker rm $(FILES)

clean:
	docker system prune -f

logs-local:
	docker logs -f $(FILES)

develop:
	echo "Starting develop environment"
	docker-compose -f docker-compose.yml up --build

local:
	echo "Starting local environment"
	docker-compose -f docker-compose.local.yml up --build
